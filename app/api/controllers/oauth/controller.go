package oauth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/api/middlewares"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/internal/auth/jwt"
	"gitlab.com/shaninalex/lumna/app/models"
)

type OAuthController struct {
	clients             ClientStore
	authCodes           AuthorizationCodeStore
	refreshTokenService RefreshTokenStore
}

func NewOAuthContoller() *OAuthController {
	s := &OAuthController{
		clients:             NewInMemoryClientStore(),
		authCodes:           NewInMemoryAuthorizationCodeStore(),
		refreshTokenService: NewPersistentRefreshTokenStore(),
	}
	return s
}
func RegisterOAuthController(router *gin.RouterGroup) {
	controller := NewOAuthContoller()

	// router already prefixed with "/oauth"
	router.GET("/authorize", LoginTokenMiddleware, controller.handleAuthorize)
	router.POST("/token", controller.handleToken)
	router.POST("/revoke", middlewares.AccessTokenMiddleware, controller.handleRevoke)
	router.POST("/revoke/all", middlewares.AccessTokenMiddleware, controller.handleRevokeAll)
}

var (
	OAuthErrorUnsuportedGratType      error = errors.New("unsupported grant type")
	OAuthErrorInvalidRequest          error = errors.New("invalid request")
	OAuthErrorInvalidClient           error = errors.New("invalid client")
	OAuthErrorInvalidGrant            error = errors.New("invalid grant")
	OAuthErrorServerError             error = errors.New("server error")
	OAuthErrorUnsupportedResponseType error = errors.New("unsupported response type")
	OAuthErrorInvalidPKCE             error = errors.New("invalid pkce")
	OAuthErrorInvalidRedirectURI      error = errors.New("invalid redirect uri")
	OAuthErrorUserNotFound            error = errors.New("user not found")
	OAuthErrorRefreshTokenRevoked     error = errors.New("refresh token revoked")
	OAuthErrorTokenExpired            error = errors.New("token expired")
)

func (s *OAuthController) handleToken(c *gin.Context) {
	grantType := c.PostForm("grant_type")

	switch grantType {
	case "authorization_code":
		s.processAuthorizationCode(c)
		return
	case "refresh_token":
		s.processRefreshToken(c)
		return
	default:
		utils.Error(c, http.StatusBadRequest, OAuthErrorInvalidGrant)
	}
}

func (s *OAuthController) processRefreshToken(c *gin.Context) {
	refreshToken := c.PostForm("refresh_token")
	clientID := c.PostForm("client_id")

	if refreshToken == "" || clientID == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_request"})
		return
	}

	hash := hashToken(refreshToken)
	rt, err := s.refreshTokenService.FindByHash(c.Request.Context(), hash)
	if err != nil {
		log.Printf("Refresh token not found by hash: %s", hash)
		utils.Error(c, http.StatusBadRequest, OAuthErrorInvalidGrant)
		return
	}

	// reuse detection
	if rt.Revoked {
		log.Printf("Security event! Attempt to refresh already revoked token ( %s )", rt.ID.String())
		utils.Error(c, http.StatusBadRequest, OAuthErrorRefreshTokenRevoked)
		return
	}

	if rt.ClientID != clientID {
		log.Println("invalid client id: ", clientID)
		utils.Error(c, http.StatusBadRequest, OAuthErrorInvalidClient)
		return
	}

	if time.Now().After(rt.ExpiresAt) {
		log.Println("unable to generate refresh token. Now: ", rt.ExpiresAt, " | Token: ", rt.ExpiresAt)
		utils.Error(c, http.StatusBadRequest, OAuthErrorTokenExpired)
		return
	}

	_ = s.refreshTokenService.Revoke(c.Request.Context(), rt.ID)

	newRefreshPlain, newRefreshHash, err := generateRefreshToken()
	if err != nil {
		log.Println("unable to generate refresh token", err)
		utils.Error(c, http.StatusInternalServerError, OAuthErrorServerError)
		return
	}

	_ = s.refreshTokenService.Save(c.Request.Context(), &models.RefreshToken{
		Hash:       newRefreshHash,
		IdentityID: rt.IdentityID,
		ClientID:   rt.ClientID,
		Scopes:     rt.Scopes,
		ExpiresAt:  time.Now().Add(30 * 24 * time.Hour),
	})

	accessToken, err := jwt.GenerateAccessJWTToken(rt.IdentityID.String(), rt.Scopes, time.Minute*15)
	if err != nil {
		log.Println("unable to generate access token", err)
		utils.Error(c, http.StatusInternalServerError, OAuthErrorServerError)
		return
	}

	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": newRefreshPlain,
		"token_type":    "Bearer",
		"expires_in":    900,
	})
}

func (s *OAuthController) processAuthorizationCode(c *gin.Context) {
	grantType := c.PostForm("grant_type")
	if grantType != "authorization_code" {
		utils.Error(c, http.StatusBadRequest, OAuthErrorUnsuportedGratType)
		return
	}

	code := c.PostForm("code")
	clientID := c.PostForm("client_id")
	redirectURI := c.PostForm("redirect_uri")
	codeVerifier := c.PostForm("code_verifier")

	if code == "" || clientID == "" || redirectURI == "" || codeVerifier == "" {
		utils.Error(c, http.StatusBadRequest, OAuthErrorInvalidRequest)
		return
	}

	// Load authorization code
	authCode, err := s.authCodes.Find(c, code)
	if err != nil {
		log.Println("auth code not found", err)
		utils.Error(c, http.StatusBadRequest, OAuthErrorInvalidGrant)
		return
	}

	// Validate client
	if authCode.ClientID != clientID {
		utils.Error(c, http.StatusBadRequest, OAuthErrorInvalidClient)
		return
	}

	// Validate redirect_uri
	if authCode.RedirectURI != redirectURI {
		log.Println("invalid redirect uri")
		utils.Error(c, http.StatusBadRequest, OAuthErrorInvalidGrant)
		return
	}

	// PKCE verification
	if !verifyPKCE(codeVerifier, authCode.CodeChallenge) {
		log.Println("invalid PRCE")
		utils.Error(c, http.StatusBadRequest, OAuthErrorInvalidGrant)
		return
	}

	// Mark code as used
	if err := s.authCodes.MarkUsed(c, code); err != nil {
		log.Println("unable to mark used")
		utils.Error(c, http.StatusBadRequest, OAuthErrorInvalidGrant)
		return
	}

	accessToken, err := jwt.GenerateAccessJWTToken(authCode.UserID, strings.Join(authCode.Scopes, " "), time.Minute*15)
	if err != nil {
		log.Println("unable to generate access token", err)
		utils.Error(c, http.StatusInternalServerError, OAuthErrorServerError)
		return
	}

	refreshToken, refreshTokenHash, err := generateRefreshToken()
	if err != nil {
		log.Println("unable to generate refresh token", err)
		utils.Error(c, http.StatusInternalServerError, OAuthErrorServerError)
		return
	}

	dbRefreshToken := models.RefreshToken{
		IdentityID: uuid.MustParse(authCode.UserID),
		Hash:       refreshTokenHash,
		Scopes:     strings.Join(authCode.Scopes, " "),
		ClientID:   authCode.ClientID,
		ExpiresAt:  time.Now().Add(30 * 24 * time.Hour),
	}

	if err := s.refreshTokenService.Save(c.Request.Context(), &dbRefreshToken); err != nil {
		panic(err)
	}

	utils.Success(c, map[string]any{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"token_type":    "Bearer",
		"expires_in":    900,
	})
}

func (s *OAuthController) handleAuthorize(c *gin.Context) {
	// Parse query params
	responseType := c.Query("response_type")
	clientID := c.Query("client_id")
	redirectURI := c.Query("redirect_uri")
	state := c.Query("state")
	scope := c.Query("scope")
	codeChallenge := c.Query("code_challenge")
	codeChallengeMethod := c.Query("code_challenge_method")

	// Basic validation
	if responseType != "code" {
		utils.Error(c, http.StatusBadRequest, OAuthErrorUnsupportedResponseType)
		return
	}

	if clientID == "" || redirectURI == "" {
		utils.Error(c, http.StatusBadRequest, OAuthErrorInvalidRequest)
		return
	}

	if !validatePKCE(codeChallenge, codeChallengeMethod) {
		utils.Error(c, http.StatusBadRequest, OAuthErrorInvalidPKCE)
		return
	}

	// Load client
	client, err := s.clients.FindByID(c, clientID)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, OAuthErrorInvalidClient)
		return
	}

	// Validate redirect_uri
	if !isRedirectAllowed(client.RedirectURIs, redirectURI) {
		utils.Error(c, http.StatusBadRequest, OAuthErrorInvalidRedirectURI)
		return
	}

	// Parse scopes
	var scopes []string
	if scope != "" {
		scopes = strings.Split(scope, " ")
	}

	userID, err := utils.GetUserID(c)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	// Generate authorization code
	code, err := generateSecureCode()
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, OAuthErrorServerError)
		return
	}

	authCode := &AuthorizationCode{
		Code:                code,
		ClientID:            clientID,
		UserID:              userID.String(),
		RedirectURI:         redirectURI,
		CodeChallenge:       codeChallenge,
		CodeChallengeMethod: codeChallengeMethod,
		Scopes:              scopes,
		ExpiresAt:           time.Now().Add(2 * time.Minute),
		Used:                false,
	}

	if err := s.authCodes.Save(c, authCode); err != nil {
		utils.Error(c, http.StatusInternalServerError, OAuthErrorServerError)
		return
	}

	// Redirect back
	redirect := redirectURI + "?code=" + url.QueryEscape(code)
	if state != "" {
		redirect += "&state=" + url.QueryEscape(state)
	}

	fmt.Println(url.QueryEscape(code))
	c.Redirect(http.StatusFound, redirect)
}

func (s *OAuthController) handleRevoke(c *gin.Context) {
	token := c.PostForm("token")
	if token == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_request"})
		return
	}

	_userID := c.GetString("userID")
	clientID := c.GetString("clientID")

	userID, err := uuid.Parse(_userID)
	if err != nil {
		// RFC: always return 200
		log.Println("invalid user id: ", err)
		c.Status(200)
		return
	}

	hash := hashToken(token)

	rt, err := s.refreshTokenService.FindByHash(c, hash)
	if err != nil {
		// RFC: always return 200
		c.Status(200)
		return
	}

	if rt.IdentityID != userID || rt.ClientID != clientID {
		// RFC: always return 200
		c.Status(200)
		return
	}

	if err := s.refreshTokenService.Revoke(c, rt.ID); err != nil {
		// RFC: always return 200
		log.Println("Unable to revoke token: ", err)
	}

	c.Status(200)
}

func (s *OAuthController) handleRevokeAll(c *gin.Context) {
	userID := c.GetString("userID")
	clientID := c.GetString("clientID")

	if err := s.refreshTokenService.RevokeAll(c, uuid.MustParse(userID), clientID); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "server_error"})
		return
	}

	c.Status(200)
}

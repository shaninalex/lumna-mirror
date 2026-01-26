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
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/internal/jwt"
)

type OAuthController struct {
	clients   ClientStore
	authCodes AuthorizationCodeStore
}

func NewOAuthContoller() *OAuthController {
	s := &OAuthController{
		clients:   NewInMemoryClientStore(),
		authCodes: NewInMemoryAuthorizationCodeStore(),
	}
	return s
}
func RegisterOAuthController(router *gin.RouterGroup) {
	controller := NewOAuthContoller()

	// router already prefixed with "/oauth"
	router.GET("/authorize", OAuthMiddleware, controller.handleAuthorize)
	router.POST("/token", controller.handleToken)
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
)

func (s *OAuthController) handleToken(c *gin.Context) {
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

	accessToken, err := jwt.GenerateAccessJWTToken(authCode.UserID, authCode.Scopes, time.Minute*15)
	if err != nil {
		log.Println("unable to generate secure code", err)
		utils.Error(c, http.StatusInternalServerError, OAuthErrorServerError)
		return
	}

	utils.Success(c, map[string]any{
		"access_token": accessToken,
		"token_type":   "Bearer",
		"expires_in":   900,
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

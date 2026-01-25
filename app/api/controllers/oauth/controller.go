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

func (s *OAuthController) handleToken(c *gin.Context) {
	grantType := c.PostForm("grant_type")
	if grantType != "authorization_code" {
		c.AbortWithStatusJSON(400, gin.H{"error": "unsupported_grant_type"})
		return
	}

	code := c.PostForm("code")
	clientID := c.PostForm("client_id")
	redirectURI := c.PostForm("redirect_uri")
	codeVerifier := c.PostForm("code_verifier")

	if code == "" || clientID == "" || redirectURI == "" || codeVerifier == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_request"})
		return
	}

	// Load authorization code
	authCode, err := s.authCodes.Find(c, code)
	if err != nil {
		log.Println("auth code not found")
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_grant"})
		return
	}

	// Validate client
	if authCode.ClientID != clientID {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_client"})
		return
	}

	// Validate redirect_uri
	if authCode.RedirectURI != redirectURI {
		log.Println("invalid redirect uri")
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_grant"})
		return
	}

	// PKCE verification
	if !verifyPKCE(codeVerifier, authCode.CodeChallenge) {
		log.Println("invalid PRCE")
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_grant"})
		return
	}

	// Mark code as used
	if err := s.authCodes.MarkUsed(c, code); err != nil {
		log.Println("unable to mark used")
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_grant"})
		return
	}

	accessToken, err := jwt.GenerateJWT(authCode.UserID, authCode.Scopes, time.Minute*15)
	if err != nil {
		log.Println("unable to generate secure code")
		c.AbortWithStatusJSON(500, gin.H{"error": "server_error"})
		return
	}

	c.JSON(200, gin.H{
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
		c.AbortWithStatusJSON(400, gin.H{"error": "unsupported_response_type"})
		return
	}

	if clientID == "" || redirectURI == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_request"})
		return
	}

	if !validatePKCE(codeChallenge, codeChallengeMethod) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_pkce"})
		return
	}

	// Load client
	client, err := s.clients.FindByID(c, clientID)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, errors.New("invalid_client"))
		return
	}

	// Validate redirect_uri
	if !isRedirectAllowed(client.RedirectURIs, redirectURI) {
		utils.Error(c, http.StatusBadRequest, errors.New("invalid_redirect_uri"))
		return
	}

	// Parse scopes
	var scopes []string
	if scope != "" {
		scopes = strings.Split(scope, " ")
	}

	userIDAny, ok := c.Get("userID")
	if !ok {
		utils.Error(c, http.StatusBadRequest, errors.New("user not found")) // make errors vars
		return
	}

	userID := uuid.MustParse(userIDAny.(string))

	// Generate authorization code
	code, err := generateSecureCode()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "server_error"})
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
		c.AbortWithStatusJSON(500, gin.H{"error": "server_error"})
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

package oauth

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
	router.GET("/authorize", controller.handleAuthorize)
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

	log.Println("code:", code)
	log.Println("clientID:", clientID)
	log.Println("redirectURI:", redirectURI)
	log.Println("codeVerifier:", codeVerifier)

	if code == "" || clientID == "" || redirectURI == "" || codeVerifier == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_request"})
		return
	}

	// Load authorization code
	authCode, err := s.authCodes.Find(c, code)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_grant 1"})
		return
	}

	// Validate client
	if authCode.ClientID != clientID {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_client"})
		return
	}

	// Validate redirect_uri
	if authCode.RedirectURI != redirectURI {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_grant 2"})
		return
	}

	// PKCE verification
	if !verifyPKCE(codeVerifier, authCode.CodeChallenge) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_grant 3"})
		return
	}

	// Mark code as used
	if err := s.authCodes.MarkUsed(c, code); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_grant 4"})
		return
	}

	accessToken, err := jwt.GenerateJWT()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "server_error"})
		return
	}

	c.JSON(200, gin.H{
		"access_token": accessToken,
		"token_type":   "Bearer",
		"expires_in":   900,
	})
}

// handleAuthorize
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
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_client"})
		return
	}

	// Validate redirect_uri
	if !isRedirectAllowed(client.RedirectURIs, redirectURI) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid_redirect_uri"})
		return
	}

	// Parse scopes
	var scopes []string
	if scope != "" {
		scopes = strings.Split(scope, " ")
	}

	// TODO: real authentication
	userID := "dev-user"

	// Generate authorization code
	code, err := generateSecureCode()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "server_error"})
		return
	}

	authCode := &AuthorizationCode{
		Code:                code,
		ClientID:            clientID,
		UserID:              userID,
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

	log.Println(redirect)
	c.Redirect(http.StatusFound, redirect)
}

func isRedirectAllowed(allowed []string, uri string) bool {
	return slices.Contains(allowed, uri)
}

func validatePKCE(challenge, method string) bool {
	if challenge == "" {
		return false
	}
	if method == "" {
		method = "S256"
	}
	return method == "S256"
}

func verifyPKCE(verifier, expectedChallenge string) bool {
	sum := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(sum[:])
	return challenge == expectedChallenge
}

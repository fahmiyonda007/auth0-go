// controllers/books.go

package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

//	@BasePath	/api/v1

// Login godoc
//
//	@Summary	auth0
//	@Schemes
//	@Description	login with auth0 user
//	@Tags			Authentication
//	@Param			login	body	LoginInput	true	"Login"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	LoginOutput
//	@Router			/login [post]
func Login(c *gin.Context) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	domain := os.Getenv("AUTH0_DOMAIN")
	audience := os.Getenv("AUTH0_AUDIENCE")
	clientId := os.Getenv("AUTH0_CLIENTID")
	clientSecret := os.Getenv("AUTH0_CLIENTSECRET")

	url := "https://" + domain + "/oauth/token"

	// Validate input
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	payload := strings.NewReader("grant_type=password&username=" + input.Username + "&password=" + input.Password + "&client_id=" + clientId + "&client_secret=" + clientSecret + "&audience=" + audience)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

	var m map[string]interface{}
	if err := json.Unmarshal(body, &m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error})
		return
	}

	c.JSON(http.StatusOK, m)
}

func Authenticate(c *gin.Context) (*jwt.Token, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return nil, fmt.Errorf("Authorization header missing")
	}

	tokenString = tokenString[len("Bearer "):]

	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	// Replace with your Auth0 secret
	// 	return []byte(os.Getenv("AUTH0_SECRET")), nil
	// })
	// if err != nil {
	// 	return nil, fmt.Errorf(err.Error())
	// }

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return token, nil
}

func HasPermission(token *jwt.Token, permission string) bool {
	// claims, ok := token.Claims.(jwt.MapClaims)
	claims := token.Claims.(jwt.MapClaims)
	// if !ok || !token.Valid {
	// 	return false
	// }
	permissions, exists := claims["permissions"].([]interface{})
	if !exists {
		return false
	}

	for _, p := range permissions {
		if p == permission {
			return true
		}
	}

	return false
}

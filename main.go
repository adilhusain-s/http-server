package main

import (
	"github.com/gin-gonic/gin"
)

func generateJWT() string {
	//generate JWT token
	return "jwt token"
}

func sensitiveAPI(jwt string) (string, error) {
	//validates JWT
	// return sensitive data
	return "sensitive data", nil
}
func handleJwtReq(c *gin.Context) {
	//validate the request
	//call generateJwt function
	jwt := generateJWT()
	c.String(200, jwt)
}
func HandleSensitiveAPI(c *gin.Context) {
	// call sensitiveAPI function with JWT token
	// return the sensitive data in http response  body to the client
	jwt := "jwt"
	sensitiveData, err := sensitiveAPI(jwt)
	if err != nil {
		return
	}
	c.String(200, sensitiveData)
}
func main() {
	router := gin.Default()

	router.POST("/jwt", handleJwtReq)
	router.POST("/sensitive", HandleSensitiveAPI)
	router.Run(":8080")
}

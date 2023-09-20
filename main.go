package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"

	"github.com/gin-gonic/gin"
)

func generateJWT(clientInfo string) string {
	//generate JWT token for this specific client
	return "jwt token"
}

func sensitiveAPI(jwt string) (string, error) {
	//validates JWT
	// return sensitive data
	return "sensitive data", nil
}
func handleJwtReq(c *gin.Context) {
	//validate the request
	//decrypt  the request body using api_key (secret)
	// parse the client info from decrypted request body

	//call generateJwt function and pass the client info as a param
	clientInfo := "client_info"
	jwt := generateJWT(clientInfo)
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

func DecryptRequest(key []byte, encryptedText string) {
	ciphertext, _ := hex.DecodeString(encryptedText)

	c, err := aes.NewCipher(key)
	CheckError(err)

	ptplaintext := make([]byte, len(ciphertext))
	c.Decrypt(ptplaintext, ciphertext)

	s := string(ptplaintext[:])
	fmt.Println("DECRYPTED:", s)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

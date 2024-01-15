package middlewares

// "github.com/dgrijalva/jwt-go"
import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// func AuthorizeJWT2(jwtService service.JWTService) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         authHeader := c.GetHeader("Authorization")
//         if authHeader == "" {
//             response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
//             c.AbortWithStatusJSON(http.StatusBadRequest, response)
//             return
//         }
//         token, err := jwtService.ValidateToken(authHeader)
//         if token.Valid {
//             claims := token.Claims.(jwt.MapClaims)
//             log.Println("Claim[user_id]: ", claims["user_id"])
//             log.Println("Claim[issuer] :", claims["issuer"])
//         } else {
//             log.Println(err)
//             response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
//             c.AbortWithStatusJSON(http.StatusUnauthorized, response)
//         }
//     }
// }

// func generateJWT() (string, error) {
// 	token := jwt.New(jwt.SigningMethodEdDSA)

// 	tokenString,err := token.SignedString(sampleSecretKey)

// 	if err != nil {
// 		return "", err
// 	}

//		return tokenString, nil
//	}
func generateJWT(privateKeyPath string) (string, error) {
	// Load private key from file
	keyBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return "", err
	}

	h := hmac.New(sha256.New, keyBytes)
	// Write the data to the hasher
	// _, err := h.Write([]byte(data))
	// if err != nil {
	// 	return "", err
	// }

	// Get the hashed result
	hashedResult := h.Sum(nil)

	// Encode the hashed result in base64 to create the token
	token := base64.StdEncoding.EncodeToString(hashedResult)

	return token, nil
	// block, _ := pem.Decode(keyBytes)
	// if block == nil {
	// 	return "", fmt.Errorf("failed to decode PEM block containing private key")
	// }

	// priv, err := x509.ParseECPrivateKey(block.Bytes)
	// if err != nil {
	// 	return "", err
	// }
	// time := 1 * time.Hour
	// fmt.Println(reflect.TypeOf(priv), priv)
	// // Create a new JWT token
	// token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
	// 	"sub": "5151",
	// 	"iss": "is",
	// 	"exp": time, // Token expiration time
	// 	// "aud": "Audience",
	// })
	// // token.Header["kid"] = "KEY-IDENTIFIER"

	// fmt.Println(token)

	// // Sign the token with the private key
	// tokenString, err := token.SignedString(priv)
	// if err != nil {
	// 	return "", err
	// }

	// return tokenString, nil
}
func AuthorizeJWT(bearerToken string) bool {
	privateKeyPath := "certs/private-key.pem"
	token, err := generateJWT(privateKeyPath)
	if err != nil {
		log.Fatal("Error generating JWT:", err)
	}
	fmt.Println("Generated JWT:", token)

	reqToken := strings.Split(bearerToken, " ")[2]
	fmt.Println("Generated JWT:", reqToken)
	if reqToken == token {
		return true
	} else {
		return false
	}
	// return func(c *gin.Context) {
	// 	bearerToken := c.Request.Header.Get("Authorization")
	// 	reqToken := strings.Split(bearerToken, " ")[1]
	// 	claims := &Claims{}
	// 	tkn, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
	// 		return jwtKey, nil
	// 	})
	// 	if err != nil {
	// 		if err == jwt.ErrSignatureInvalid {
	// 			c.JSON(http.StatusUnauthorized, gin.H{
	// 				"message": "unauthorized",
	// 			})
	// 			return
	// 		}
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": "bad request",
	// 		})
	// 		return
	// 	}
	// 	if !tkn.Valid {
	// 		c.JSON(http.StatusUnauthorized, gin.H{
	// 			"message": "unauthorized",
	// 		})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data": "resource data",
	// 	})
	// }

}

package middlewares

// "github.com/dgrijalva/jwt-go"
import (
	"fmt"
	"io/ioutil"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

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
func generateJWT() {
	signBytes, err := ioutil.ReadFile("certs/healthcheck")
	if err != nil {
		log.Fatalln(err)
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(signKey))

	verifyBytes, err := ioutil.ReadFile("certs/healthcheck.pub")
	if err != nil {
		log.Fatalln(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalln(err)
	}

}
func AuthorizeJWT() {

}

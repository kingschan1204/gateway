package plugin

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type TokenClaims struct {
	// user name eq account
	Username string `json:"username"`
	// tenant id
	Tenant string `json:"tenant"`
	jwt.RegisteredClaims
}

//func Generate() {
//	MySecret := []byte("mySecret")
//	fmt.Println(time.Now())
//	token, err := GenToken("guest", "tenant001", MySecret,"1h10m")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println("token : ", token)
//
//	mc, err := ParseToken(token, MySecret)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println("TokenClaims -> ", mc)
//	return
//}

// GenToken Generate jwt token
// username : user name
// tenant : saas tenant id
// secret : jwt secret key
// expireTime : jwt expire time
func GenToken(username string, tenant string, secret []byte, expireTime string) (string, error) {
	expire, _ := time.ParseDuration(expireTime)
	// Create a statement of our own
	c := TokenClaims{
		username, // custom field
		tenant,
		jwt.RegisteredClaims{
			//token expiration time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
			Issuer:    "api-gateway",
		},
	}

	// Create a signature object using the specified signature method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Sign with the specified secret and get the complete encoded string token
	return token.SignedString(secret)
}

// ParseToken parsing token
// tokenString : token
// secret :
func ParseToken(tokenString string, secret []byte) (*TokenClaims, error) {
	if tokenString == "" {
		return nil, errors.New("invalid token")
	}
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid { // validate token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

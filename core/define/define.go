package define

import (
	"github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var (
	JwtKey    = "cloud-disk-key"
	AWSBucket = "gobuckettest"
	Region    = "us-west-2"
)

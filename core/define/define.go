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
	PageSize  = 20
	Datetime  = "2006-01-02 15:04:05"
)

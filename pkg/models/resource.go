package models

import (
	"crypto"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"github.com/golang-jwt/jwt/v4"
	"github.com/minio/minio-go/v7"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// build
type Resources struct {
	LogConfig      logger.Config
	FastHTTPClient *fasthttp.Client
	MainDbConn     *gorm.DB
	LogDbConn      *mongo.Database
	RedisStorage   *redis.Storage
	JwtResources   *JwtResources
	SessConfig     session.Config
	Minio          *minio.Client
	Mqtt           *mqtt.Client
}

type JwtResources struct {
	JwtVerifyKey     crypto.PublicKey
	JwtSignKey       crypto.PrivateKey
	JwtSigningMethod jwt.SigningMethod
	JwtKeyfunc       jwt.Keyfunc
	JwtParser        *jwt.Parser
}

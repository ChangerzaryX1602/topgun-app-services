package infrastructure

import (
	"fmt"

	"top-gun-app-services/internal/datasources"
	"top-gun-app-services/pkg/models"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/storage/redis"
	"github.com/golang-jwt/jwt/v4"
	"github.com/minio/minio-go/v7"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func NewResources(fasthttpClient *fasthttp.Client, mainDbConn *gorm.DB, logDbConn *mongo.Database, redisStorage *redis.Storage, jwtResources *models.JwtResources, minio *minio.Client, mqtt mqtt.Client, mqttOption *mqtt.ClientOptions) models.Resources {
	return models.Resources{
		FastHTTPClient: fasthttpClient,
		MainDbConn:     mainDbConn,
		LogDbConn:      logDbConn,
		RedisStorage:   redisStorage,
		JwtResources:   jwtResources,
		Minio:          minio,
		Mqtt:           mqtt,
	}
}

func NewJwt(privateKeyPath string) (jwtResources *models.JwtResources, err error) {
	jwtResources = new(models.JwtResources)
	jwtResources.JwtSignKey, jwtResources.JwtVerifyKey, jwtResources.JwtSigningMethod, jwtResources.JwtKeyfunc, err = datasources.NewJwtLocalKey(privateKeyPath)
	jwtResources.JwtKeyfunc = func(token *jwt.Token) (publicKey interface{}, err error) {
		if jwtResources.JwtVerifyKey == nil {
			err = fmt.Errorf("JWTVerifyKey not init yet")
		}
		return jwtResources.JwtVerifyKey, err
	}
	jwtResources.JwtParser = jwt.NewParser()

	return
}

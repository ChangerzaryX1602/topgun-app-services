package infrastructure

import (
	"fmt"
	"log"

	"top-gun-app-services/internal/handlers"
	"top-gun-app-services/pkg/auth"
	"top-gun-app-services/pkg/mqtt"
	"top-gun-app-services/pkg/user"
	"top-gun-app-services/pkg/workshop"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes is the Router for GoFiber App
func (s *Server) SetupRoutes(app *fiber.App) {

	// Prepare a static middleware to serve the built React files.
	app.Static("/", "./web/build")
	// API routes group
	groupApiV1 := app.Group("/api/v:version?", handlers.ApiLimiter)
	{
		groupApiV1.Get("/", handlers.Index())
	}
	router := handlers.NewRouterResources(s.JwtResources.JwtKeyfunc)
	// App
	userRepository := user.NewUserRepository(s.MainDbConn)
	authRepository := auth.NewAuthRepository(s.MainDbConn)
	mqttRepository := mqtt.NewMQTTRepository(s.MainDbConn)
	workshopRepository := workshop.NewWorkshopRepository(s.MainDbConn)
	checkAndAutoMigrate(s.MainDbConn, &user.User{}, &workshop.RawData{})
	userUsecase := user.NewUserService(userRepository)
	authUsecase := auth.NewAuthService(authRepository)
	mqttUsecase := mqtt.NewMQttService(mqttRepository)
	workshopUsecase := workshop.NewWorkshopService(workshopRepository)
	user.NewUserHandler(app.Group("/api/v1/users"), userUsecase, router)
	auth.NewAuthHandler(app.Group("/api/v1/auth"), authUsecase, *s.JwtResources)
	mqtt.NewMQttHandler(app.Group("/api/v1/mqtt"), mqttUsecase, s.Mqtt, s.MqttOption)
	workshop.NewWorkshopHandler(app.Group("/api/v1/machine"), workshopUsecase)
	// wsURL := viper.GetString("workshop.ws")
	// apiKey := viper.GetString("workshop.key")
	// workshopUsecase.ConnectWebSocket(wsURL, apiKey)
	// Prepare a fallback route to always serve the 'index.html', had there not be any matching routes.
	app.Static("*", "./web/build/index.html")
}
func checkAndAutoMigrate(db *gorm.DB, model ...interface{}) {
	for _, m := range model {
		// Check if the table does not exist
		if !db.Migrator().HasTable(m) {
			// Auto migrate the table if it does not exist
			if err := db.AutoMigrate(m); err != nil {
				log.Fatalf("Failed to auto migrate table: %v", err)
			}
			fmt.Printf("Table %T created successfully.\n", m)
		} else {
			fmt.Printf("Table %T already exists.\n", m)
		}
	}
}

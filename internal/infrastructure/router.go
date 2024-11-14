package infrastructure

import (
	"fmt"
	"log"

	"top-gun-app-services/internal/handlers"
	"top-gun-app-services/pkg/attachment"
	"top-gun-app-services/pkg/auth"
	"top-gun-app-services/pkg/mqtt"
	"top-gun-app-services/pkg/user"
	"top-gun-app-services/pkg/workshop"

	//swagger
	_ "top-gun-app-services/docs"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
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
	//swagger path
	app.Get("/api/v1/swagger/*", fiberSwagger.WrapHandler)
	router := handlers.NewRouterResources(s.JwtResources.JwtKeyfunc)
	// App
	userRepository := user.NewUserRepository(s.MainDbConn)
	authRepository := auth.NewAuthRepository(s.MainDbConn)
	mqttRepository := mqtt.NewMQTTRepository(s.MainDbConn)
	workshopRepository := workshop.NewWorkshopRepository(s.MainDbConn)
	attachmentRepository := attachment.NewAttachmentRepository(s.MainDbConn)
	checkAndAutoMigrate(s.MainDbConn, &user.User{}, &workshop.RawData{}, &attachment.AttachFile{})
	userUsecase := user.NewUserService(userRepository)
	authUsecase := auth.NewAuthService(authRepository)
	mqttUsecase := mqtt.NewMQttService(mqttRepository, s.Mqtt, s.MqttOption)
	workshopUsecase := workshop.NewWorkshopService(workshopRepository)
	attachmentUsecase := attachment.NewAttachmentService(attachmentRepository)
	user.NewUserHandler(app.Group("/api/v1/users"), userUsecase, router)
	auth.NewAuthHandler(app.Group("/api/v1/auth"), authUsecase, *s.JwtResources, router)
	mqtt.NewMQttHandler(app.Group("/api/v1/mqtt"), mqttUsecase, s.Mqtt, s.MqttOption)
	workshop.NewWorkshopHandler(app.Group("/api/v1/machine"), workshopUsecase, router)
	attachment.NewWorkshopHandler(app.Group("/api/v1/attachment"), attachmentUsecase, mqttUsecase, router)
	
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

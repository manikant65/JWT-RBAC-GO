package main

// load required packages
import (
	"fmt"
	"log"
	"manikant/jwt-go-rbac/controller"
	"manikant/jwt-go-rbac/database"
	"manikant/jwt-go-rbac/model"
	"manikant/jwt-go-rbac/util"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// main function
func main() {
	// load environment file
	loadEnv()
	// load database configuration and connection
	loadDatabase()
	// start the server
	serveApplication()
}

// load environment variables file
func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

// run database migrations and add seed data
func loadDatabase() {
	database.InitDb()
	database.Db.AutoMigrate(&model.Role{})
	database.Db.AutoMigrate(&model.User{})
	database.Db.AutoMigrate(&model.Room{})
	database.Db.AutoMigrate(&model.Booking{})
	seedData()
}

// load seed data into the database
func seedData() {
	var roles = []model.Role{{Name: "admin", Description: "Administrator role"}, {Name: "customer", Description: "Authenticated customer role"}, {Name: "visitor", Description: "Unauthenticated customer role"}}
	var user = []model.User{{Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: os.Getenv("ADMIN_PASSWORD"), RoleID: 1}}
	database.Db.Save(&roles)
	database.Db.Save(&user)
}

// start the server on port 8000
func serveApplication() {
	router := gin.Default()

	authRoutes := router.Group("/auth/user")
	authRoutes.POST("/register", controller.Register)
	authRoutes.POST("/login", controller.Login)

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(util.JWTAuth())
	adminRoutes.GET("/users", controller.GetUsers)
	adminRoutes.GET("/user/:id", controller.GetUser)
	adminRoutes.PUT("/user/:id", controller.UpdateUser)
	adminRoutes.POST("/user/role", controller.CreateRole)
	adminRoutes.GET("/user/roles", controller.GetRoles)
	adminRoutes.PUT("/user/role/:id", controller.UpdateRole)
	adminRoutes.POST("/room/add", controller.CreateRoom)
	adminRoutes.PUT("/room/:id", controller.UpdateRoom)
	adminRoutes.GET("/room/bookings", controller.GetBookings)

	publicRoutes := router.Group("/api/view")
	publicRoutes.GET("/rooms", controller.GetRooms)
	publicRoutes.GET("/room/:id", controller.GetRoom)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(util.JWTAuthCustomer())
	protectedRoutes.GET("/rooms/booked", controller.GetUserBookings)
	protectedRoutes.POST("/room/book", controller.CreateBooking)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}

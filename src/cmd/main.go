package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/trongtb88/wagerservice/src/cmd/db"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	resthandler "github.com/trongtb88/wagerservice/src/handler/rest"
	// Business Layer Dep
	domain "github.com/trongtb88/wagerservice/src/business/domain"
	usecase "github.com/trongtb88/wagerservice/src/business/usecase"
)

var (
	sqlClient0     *gorm.DB

	// Server Infrastructure
	logger     log.Logger
	//staticConf cfg.Conf
	//remoteConf cfg.Conf
	//secretConf cfg.Conf
	//tele       telemetry.Telemetry
	//schedule   scheduler.Scheduler
	//assign     assignor.Assignor
	//healt      health.Health
	//parse      parser.Parser
	//httpMware  httpmiddleware.HttpMiddleware
	//httpMux    httpmux.HttpMux
	//aut        auth.Auth
	//httpServer httpserver.HTTPServer
	//app        grace.App

	// Business Layer
	dom *domain.Domain
	uc  *usecase.Usecase
)


//func ConnectDB(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *gorm.DB {
//	var (
//		db *gorm.DB
//		err error
//	)
//	if Dbdriver == "mysql" {
//		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
//		db, err = gorm.Open(mysql.Open(DBURL))
//		if err != nil {
//			fmt.Printf("Cannot connect to %s database", Dbdriver)
//			log.Fatal("This is the error:", err)
//		} else {
//			fmt.Printf("We are connected to the %s database", Dbdriver)
//		}
//	}
//
//	err = db.Debug().AutoMigrate(&entity.Wager{}) //database migration
//	if err != nil {
//		log.Fatal("Error when migration table:", err)
//	}
//	return db
//}

func main() {


	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	db := db.ConnectDB (
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	// Business layer Initialization
	dom = domain.Init(
		db,
	)
	uc = usecase.Init(dom)

	serverPort := os.Getenv("SERVER_PORT")

	router := mux.NewRouter().
		PathPrefix("/api/v1"). // add prefix for v1 api `/api/v1/`
		Subrouter()

	// REST Handler Initialization
	_ = resthandler.Init(logger, router,  uc)

	log.Println("Starting server at port: ", serverPort)

	err = http.ListenAndServe(":"+serverPort, router)
	if err != nil {
		log.Println(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	iconfig "github.com/inhun/GoropBox/config"
	"github.com/inhun/GoropBox/endpoints"
	iauth "github.com/inhun/GoropBox/internal/auth"
	iaws "github.com/inhun/GoropBox/internal/aws"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	cfg, _ := iconfig.LoadConfig("config.json")

	S3Client, err := iaws.LoadS3Client(cfg.AWS)
	if err != nil {
		log.Fatal(err)
	}

	/*
		putoutput, err := S3Client.GetObject(context.TODO(), &s3.GetObjectInput{Bucket: aws.String("choicafe"), Key: aws.String("pepsi.jpg")})
		fmt.Println(&putoutput.Body)

			fmt.Println(client)
			output, _ := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})

			for i, a := range output.Buckets {
				fmt.Println(i, *a.Name)
			}
	*/

	OauthConf := iauth.LoadAuthConfig(cfg.Google)

	DBUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=%s", cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Port, cfg.DB.TimeZone)

	db, err := gorm.Open(postgres.Open(DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	ep := endpoints.Endpoints{
		DB:       db,
		Oauth:    OauthConf,
		S3Client: S3Client,
	}

	// db.AutoMigrate(&models.User{})
	// db.Create(&models.User{Email: "inhun321@khu.ac.kr", Password: "test"})

	router := httprouter.New()

	router.GET("/api/user", ep.GetUserList)
	router.GET("/api/oauth2/google", ep.Oauth2Google)
	// router.POST("/api/signin", ep.Signin)
	router.GET("/login/oauth2/code/google", ep.CallbackGoogle)
	router.POST("/api/file/new", ep.Uploads)

	handler := cors.AllowAll().Handler(router)
	port := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(port, handler))

}

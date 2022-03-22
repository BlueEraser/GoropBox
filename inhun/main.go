package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	iconfig "github.com/inhun/GoropBox/config"
	"github.com/inhun/GoropBox/endpoints"
	iaws "github.com/inhun/GoropBox/internal/aws"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	"golang.org/x/oauth2"
)

func main() {
	// practice.A()

	cfg, _ := iconfig.LoadConfig("config.json")

	client, err := iaws.LoadS3Client(cfg.AWS)
	if err != nil {
		log.Fatal(err)
	}
	/*
		fmt.Println(client)
		output, _ := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})

		for i, a := range output.Buckets {
			fmt.Println(i, *a.Name)
		}
	*/

	output, err := client.PutObject(context.TODO(), &s3.PutObjectInput{Bucket: aws.String("goropbox"), Key: aws.String("config.txt"), Body: bytes.NewBuffer([]byte("example object!")), ContentLength: 15})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)

	var OauthConf = &oauth2.Config{
		ClientID:     cfg.Google.ClientID,
		ClientSecret: cfg.Google.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		RedirectURL:  cfg.Google.RedirectUrl,
		Endpoint: oauth2.Endpoint{
			TokenURL: cfg.Google.TokenUrl,
			AuthURL:  cfg.Google.AuthUrl,
		},
	}

	DBUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=%s", cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Port, cfg.DB.TimeZone)

	db, err := gorm.Open(postgres.Open(DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	ep := endpoints.Endpoints{
		DB:    db,
		Oauth: OauthConf,
	}

	// db.AutoMigrate(&models.User{})
	// db.Create(&models.User{Email: "inhun321@khu.ac.kr", Password: "test"})

	router := httprouter.New()

	router.GET("/api/user", ep.GetUserList)
	router.GET("/api/oauth2/google", ep.Oauth2Google)
	// router.POST("/api/signin", ep.Signin)
	router.GET("/login/oauth2/code/google", ep.CallbackGoogle)

	handler := cors.AllowAll().Handler(router)
	port := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(port, handler))

}

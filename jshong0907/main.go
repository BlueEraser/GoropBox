package main

import (
	"fmt"
	"gorop-box/routers"
	"gorop-box/services"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routers.Init(e)

	// S3 파일 업로드 함수 테스트
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println("file open error!")
		os.Exit(1)
	}
	defer file.Close()
	err = services.UploadFile(file)
	if err != nil {
		fmt.Println("file upload error!")
	}

	e.Logger.Fatal(e.Start(":1323"))
}

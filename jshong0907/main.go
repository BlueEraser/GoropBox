package main

import (
	"fmt"

	"gorop-box/services"
)

func main() {
	user := services.CreateUser("jshong0907@google.com", "1234qwer", "준식홍")
	fmt.Println(user.CheckPassword("1234qwer"))
}

package main

import (
	"fmt"
	"learn-sql/user"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	user.ConnectDB()
	// u := &user.User{
	// 	FirstName: "Kanokon",
	// 	LastName:  "Chongnguluam",
	// 	Email:     "singpor@gmail",
	// }
	// user.Insert(u)
	us, err := user.All()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", us)

	u, err := user.FindByID(2)
	if err != nil {
		log.Fatal(err)
	}
	err = user.Delete(u)
	if err != nil {
		log.Fatal(err)
	}
	us, err = user.All()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", us)
}

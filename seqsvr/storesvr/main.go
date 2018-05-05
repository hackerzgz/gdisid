package main

import (
	"fmt"

	"github.com/hackez/gdisid/seqsvr/storesvr/router"
)

func main() {
	router.Register()
	err := router.Run()
	if err != nil {
		fmt.Println("storesvr panic: ", err)
	}

	fmt.Println("storesvr leave success")
}

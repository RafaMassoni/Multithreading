package main

import (
	"fmt"

	"github.com/RafaMassoni/Multithreading.git/client"
)

func main() {

	adress, err := client.GetFirstAdress("01153000")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nRESPONSE -> ", adress)

}

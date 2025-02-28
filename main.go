package main

import (
	"fmt"
	"log"

	"redjax/my-gh/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}
}

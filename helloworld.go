package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	executable, _ := os.Executable()
	curwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello World!")
	fmt.Println(executable)

	fmt.Println("current path: ", curwd)

	species := flag.String("species", "gopher", "the species we are studying")
	flag.Parse()
	fmt.Println(*species)

}

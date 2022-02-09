package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := ioutil.TempFile("", "webd")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file.Name())
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Fatal(err)
		}
	}(file.Name())

}
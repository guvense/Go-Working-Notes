package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	body, lenght := response("http://golang.org.tr/")

	fmt.Println(lenght)
	fmt.Println(body)

}

func response(url string) (string, int) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body), len(body)

}

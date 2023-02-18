package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Co webRequests")
	response, err := http.Get(url)
	checkErrorNill(err)
	fmt.Printf("Type of response %T", response)
	defer response.Body.Close()
	databyte, err := io.ReadAll(response.Body)
	checkErrorNill(err)
	content := string(databyte)

	fmt.Println(" Body content of website is: ", content)

}

func checkErrorNill(err error) {
	if err != nil {
		panic(err)
	}
}

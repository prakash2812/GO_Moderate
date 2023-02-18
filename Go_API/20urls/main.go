package main

import (
	"fmt"
	"net/url"
)

const URL = "http://lco.dev:3000/learn?coursename=react&payment=ifjiojreif"

func main() {
	fmt.Println("URL concepts")
	result, _ := url.Parse(URL)
	fmt.Printf("type of result %T\n", result)
	fmt.Println("SCHEME", result.Scheme)
	fmt.Println("host", result.Host)
	fmt.Println("port", result.Port())
	fmt.Println("params", result.RawQuery)

	qparams := result.Query()
	fmt.Println("qparams", qparams)

	for _, val := range qparams {
		fmt.Println("params are", val)
	}

	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "loc.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	fmt.Println("url formation", partsOfURL.String())
}

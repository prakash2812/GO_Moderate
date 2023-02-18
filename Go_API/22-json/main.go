package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Platform string
	Password string `json:"-"`
	Price    int
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSOn concepts")
	EncodeJSON()
	DecodeJSON()
}
func EncodeJSON() {
	lcoCourses := []course{
		{"react", "dev", "123", 212, []string{"dev"}},
		{"mern", "dev", "fds3", 212, []string{"dev"}},
		{"GO", "dev", "ar3", 212, nil},
	}
	data, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	// var responseData strings.Builder
	// result,_:=responseData.Write(data)
	// fmt.Println("data", responseData.String(result))
	fmt.Printf("%s\n", data)

}

func DecodeJSON() {
	jsonFormatfromWeb := []byte(`
			{
				"coursename": "react",
        "Platform": "dev",
        "Price": 212,
        "tags": ["dev"]
			}
		`)
	var lcoCourse course
	checkValid := json.Valid(jsonFormatfromWeb)
	if checkValid {
		fmt.Println("valid json")
		json.Unmarshal(jsonFormatfromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json not valid")
	}
	// sometimes where we add data to key values
	var onlineData map[string]interface{}
	json.Unmarshal(jsonFormatfromWeb, &onlineData)
	fmt.Println("My onlive data", onlineData)
	fmt.Printf("My onlive data %#v", onlineData)
	onlineData["Price"] = "2812"
	for key, value := range onlineData {
		fmt.Printf("key is %v and value is %v and type is %T \n", key, value, value)
	}
}

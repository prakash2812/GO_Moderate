package service

import (
	"encoding/json"
	"io/ioutil"
)

type Data struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

type Payload struct {
	Data []Data
}

const path = "data.json"

func Raw() ([]Data, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	isJsonvalid := json.Valid(file)
	var payload Payload
	if isJsonvalid {
		json.Unmarshal(file, &payload.Data)

	}
	return payload.Data, err
}
func GetAll() ([]Data, error) {
	data, err := Raw()
	if err != nil {
		return nil, err
	}
	return data, err
}

func GetById(idx int) (any, error) {
	data, err := Raw()
	if err != nil {
		return nil, err
	}
	if idx > len(data) {
		res := make([]string, 0)
		return res, nil
	}
	for _, item := range data {
		if item.Id == idx {
			return item, nil
		}
	}
	return data[idx], nil
}

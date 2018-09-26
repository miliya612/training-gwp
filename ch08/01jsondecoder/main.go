package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	ID int `json:"id"`
	Content string `json:"content"`
	Author Author `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		fmt.Println("error decoding JSON:", err)
		return
	}
	return
}

func unmarshall(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("error reading JSON data:", err)
		return
	}
	json.Unmarshal(jsonData, &post)
	return
}

func main() {
	_, err := decode("post.json")
	if err != nil {
		fmt.Println("error:", err)
	}
}
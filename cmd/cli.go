package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	/*	err := cli.NewApp().Run(os.Args)
		if err != nil {
			log.Fatal(err)
		}*/
	fmt.Println("Client Started")
	//appPort := "8000"
	content, err := ioutil.ReadFile("abc.txt")
	//var person Person
	//err = json.Unmarshal(content, &person)
	// json.Unmarshal Error
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Successfully Opened file passing it to server", content)

	req, err := http.NewRequest("POST", "http://localhost:9000/add", bytes.NewBuffer(content))

	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("Content-Type", "text/plain")
	//req.Header.Set("fileType", "XML")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
}

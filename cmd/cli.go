package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("Client Started")
	/*/var person Person
	err = json.Unmarshal(content, &person)
	/* json.Unmarshal Error
	if err != nil {
		log.Panic(err)
	}*/

	/*/Sending files to store
		content, err := ioutil.ReadFile("abc.txt")

		fmt.Println("Successfully Opened file passing it to server")

		req, err := http.NewRequest("POST", "http://localhost:9000/add", bytes.NewBuffer(content))

		if err != nil {
			log.Panic(err)
		}
		req.Header.Set("Content-Type", "text/plain")
		req.Header.Set("fileName", "abc.txt")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
	/*/
	//To get list of all files
	response, err := http.Get("http://localhost:9000/list")
	if err != nil {
		fmt.Println("HTTP req failed with error", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	} //*/

}

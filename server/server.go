package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//To Store directory path

var dir string

func init() {
	dir, _ = os.Getwd()
	dir = dir + "\\store"
}

//For sending rsponse to client
type response struct {
	Message string
}

//To add file to the store

func addFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add file to store endpoint Hit")

	fileName := r.Header.Get("fileName")
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	check := false
	w.Header().Set("Content-Type", "application/json")
	for _, file := range files {
		if file.Name() == fileName {
			check = true
			resp := &response{Message: "Failed to add file-File Already present"}

			json.NewEncoder(w).Encode(resp)
		}
	}
	if !check {
		fileName = dir + "\\" + fileName
		fmt.Println(dir)

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("---", string(fileBytes))
		// write this byte array to  file
		ioutil.WriteFile(fileName, fileBytes, 0644)

		resp := &response{Message: "File uploaded successfully"}
		json.NewEncoder(w).Encode(resp)
	}
}

//To list file present on the store
func listFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}

//To delete file
func removeFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fileName := params["fileName"]
	fileP := dir + "\\" + fileName
	//fmt.Println("-------", fileP)
	err := os.Remove(fileP)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("File Successfully deleted")

}

//Update file
func updateFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update hit")
	params := mux.Vars(r)
	fileName := params["fileName"]

	fileName = dir + "\\" + fileName
	fmt.Println(fileName)

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	// write this byte array to  file

	ioutil.WriteFile(fileName, fileBytes, 0644)
	fmt.Fprintf(w, "Successfully Updated File\n")
}

func main() {
	fmt.Println("Server Started")

	router := mux.NewRouter()

	router.HandleFunc("/add", addFile).Methods("POST")
	router.HandleFunc("/list", listFile).Methods("GET")
	router.HandleFunc("/remove/{fileName}", removeFile).Methods("DELETE")
	router.HandleFunc("/update/{fileName}", updateFile).Methods("PUT")
	//http.HandleFunc("/add", addFile)

	http.ListenAndServe(":9000", router)

}

package main

import (
	"bufio"
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

//For sending list ofcommand
type lsFile struct {
	File []string
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
		//fmt.Println("---", string(fileBytes))
		// write this byte array to  file
		ioutil.WriteFile(fileName, fileBytes, 0644)

		resp := &response{Message: "File uploaded successfully"}
		json.NewEncoder(w).Encode(resp)
	}
}

//To list file present on the store
func listFile(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var listF lsFile
	for _, file := range files {
		fmt.Println(file.Name())
		listF.File = append(listF.File, file.Name())
	}
	json.NewEncoder(w).Encode(listF)
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
	mesg := fileName + " File deleted successfully "
	resp := &response{Message: mesg}
	json.NewEncoder(w).Encode(resp)
	fmt.Println(mesg)

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
	mesg := params["fileName"] + " File updated successfully "
	resp := &response{Message: mesg}
	json.NewEncoder(w).Encode(resp)
	fmt.Println(mesg)

}

//Count Words in all the file
func countWords(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	count := 0

	for _, file := range files {

		fileHandle, err := os.Open(dir + "\\" + file.Name())

		// check if file-handle was initiated correctly
		if err != nil {
			panic(err)
		}

		// to close file-handle upon return
		defer fileHandle.Close()

		// initiate scanner from file handle
		fileScanner := bufio.NewScanner(fileHandle)

		// tell the scanner to split by words
		fileScanner.Split(bufio.ScanWords)

		// for looping through results
		for fileScanner.Scan() {
			//fmt.Printf("word: '%s' - position: '%d'\n", fileScanner.Text(), count)

			count++
		}

		// check if there was an error while reading words from file
		if err := fileScanner.Err(); err != nil {
			panic(err)
		}
	}
	json.NewEncoder(w).Encode(count)
	fmt.Println(count)

}

//Count Frequency of all Words in all the file
func freqWords(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	wordMap := make(map[string]int)
	for _, file := range files {

		fileHandle, err := os.Open(dir + "\\" + file.Name())

		// check if file-handle was initiated correctly
		if err != nil {
			panic(err)
		}

		// to close file-handle upon return
		defer fileHandle.Close()

		// initiate scanner from file handle
		fileScanner := bufio.NewScanner(fileHandle)

		// tell the scanner to split by words
		fileScanner.Split(bufio.ScanWords)

		// for looping through results
		for fileScanner.Scan() {
			//fmt.Printf("word: '%s' - position: '%d'\n", fileScanner.Text(), count)
			wordMap[fileScanner.Text()] += 1
			count++
		}

		// check if there was an error while reading words from file
		if err := fileScanner.Err(); err != nil {
			panic(err)
		}
	}
	json.NewEncoder(w).Encode(wordMap)
	fmt.Println(count)
	fmt.Println("Word Map---", wordMap)
}
func main() {
	fmt.Println("Server Started")

	router := mux.NewRouter()

	router.HandleFunc("/add", addFile).Methods("POST")
	router.HandleFunc("/list", listFile).Methods("GET")
	router.HandleFunc("/remove/{fileName}", removeFile).Methods("DELETE")
	router.HandleFunc("/update/{fileName}", updateFile).Methods("PUT")
	router.HandleFunc("/count", countWords).Methods("GET")
	router.HandleFunc("/freq", freqWords).Methods("GET")
	http.ListenAndServe(":9000", router)

}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//var serverdir

//To add file to the store
func addFile(w http.ResponseWriter, r *http.Request) {

	fmt.Println("ADD FILE HIT")
	// params := mux.Vars(r)
	//fileName:=params["fileName"]

	// maximum file upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	//r.Fi
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	dir, _ := os.Getwd()
	fmt.Println(dir)
	tempFile, err := ioutil.TempFile(dir, "upload-*.pdf")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

//To list file present on the store
func listFile(w http.ResponseWriter, r *http.Request) {

}
func main() {
	fmt.Println("Main Started")
	////router := mux.NewRouter()
	//router.HandleFunc("/add", addFile).Methods("POST")
	//router.HandleFunc("/list", listFile).Methods("GET")
	//router.HandleFunc("/delete/{name}", deletePerson).Methods("DELETE")
	//router.HandleFunc("/update/{name}", updatePerson).Methods("PUT")
	http.HandleFunc("/add", addFile)
	http.ListenAndServe(":7000", nil)

}

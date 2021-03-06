/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// storeUpdateCmd represents the storeUpdate command
var storeUpdateCmd = &cobra.Command{
	Use:   "storeUpdate",
	Short: "This command will update file present on directory",
	Long:  `It is part of fileStoreService command line Application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("storeUpdate called")
		content, err := ioutil.ReadFile(args[0])

		fmt.Println("Successfully Opened file passing it to server to update")
		updateReq := "http://localhost:9000/update/" + args[0]
		req, err := http.NewRequest("PUT", updateReq, bytes.NewBuffer(content))

		if err != nil {
			log.Panic(err)
		}
		req.Header.Set("Content-Type", "text/plain")
		//req.Header.Set("fileName", "abc.txt")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		var a Response
		json.Unmarshal(data, &a)
		fmt.Println(a)

	},
}

func init() {
	rootCmd.AddCommand(storeUpdateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storeUpdateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storeUpdateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

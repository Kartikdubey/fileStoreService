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
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// storeWcCmd represents the storeWc command
var storeWcCmd = &cobra.Command{
	Use:   "storeWc",
	Short: "Counts the number of words in all the file stored on server",
	Long:  `It is part of fileStoreService command line Application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("storeWc called")
		response, err := http.Get("http://localhost:9000/count")
		if err != nil {
			fmt.Println("HTTP req failed with error", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println("Total no of words in all files", string(data))
		}
	},
}

func init() {
	rootCmd.AddCommand(storeWcCmd)

}

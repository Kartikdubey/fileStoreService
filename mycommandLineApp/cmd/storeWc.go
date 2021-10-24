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

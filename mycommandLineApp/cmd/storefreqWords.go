package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// storefreqWordsCmd represents the storefreqWords command
var storefreqWordsCmd = &cobra.Command{
	Use:   "storefreqWords",
	Short: "Command to print most or less frequentWords",
	Long:  `Part of filestoreService application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("storefreqWords called")
		response, err := http.Get("http://localhost:9000/freq")
		if err != nil {
			fmt.Println("HTTP req failed with error", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println("Freq of words in all files", string(data))
		}
	},
}

func init() {
	rootCmd.AddCommand(storefreqWordsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storefreqWordsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storefreqWordsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

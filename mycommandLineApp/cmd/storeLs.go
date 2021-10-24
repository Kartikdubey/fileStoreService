package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

//To store list of files
type lsFile struct {
	File []string
}

// storeLsCmd represents the storeLs command

var storeLsCmd = &cobra.Command{
	Use:   "storeLs",
	Short: "This command will give list of files",
	Long:  `Part of filestoreService application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("storeLs called")
		response, err := http.Get("http://localhost:9000/list")
		if err != nil {
			fmt.Println("HTTP req failed with error", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			var a lsFile
			json.Unmarshal(data, &a)
			fmt.Println("List of files present in directory", a)

		}
	},
}

func init() {
	rootCmd.AddCommand(storeLsCmd)

}

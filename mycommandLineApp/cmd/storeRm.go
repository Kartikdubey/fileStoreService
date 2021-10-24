package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// storeRmCmd represents the storeRm command
var storeRmCmd = &cobra.Command{
	Use:   "storeRm",
	Short: "This command will remove given file from filestore",
	Long:  `Part of filestoreService application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("storeRm called")
		client := &http.Client{}
		for _, args := range args {
			// Delete request
			delReq := "http://localhost:9000/remove/" + args
			req, err := http.NewRequest("DELETE", delReq, nil)
			if err != nil {
				fmt.Println(err)
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)

			}
			data, _ := ioutil.ReadAll(resp.Body)
			var a Response
			json.Unmarshal(data, &a)
			fmt.Println(a)
		}
	},
}

func init() {
	rootCmd.AddCommand(storeRmCmd)

}

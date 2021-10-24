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

//For storing rsponse to client

type response struct {
	Message string
}

// storeAddCmd represents the storeAdd command

var storeAddCmd = &cobra.Command{
	Use:   "storeAdd",
	Short: "This command will add files to store",
	Long:  `It is part of fileStoreService command line Application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("storeAdd called")
		for _, args := range args {
			fmt.Println("-AA", args)
			content, err := ioutil.ReadFile(args)

			fmt.Println("Successfully Opened file passing it to server", string(content), content)

			req, err := http.NewRequest("POST", "http://localhost:9000/add", bytes.NewBuffer(content))

			if err != nil {
				log.Panic(err)
			}
			req.Header.Set("Content-Type", "text/plain")
			req.Header.Set("fileName", args)
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			data, _ := ioutil.ReadAll(resp.Body)
			var a response
			json.Unmarshal(data, &a)
			fmt.Println(a)
		}
	},
}

func init() {
	rootCmd.AddCommand(storeAddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storeAddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storeAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

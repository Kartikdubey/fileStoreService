package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var wordMap map[string]int

//To get more frequent words
func moreFrequent(word map[string]int) {
	fmt.Println("more freq")
	var a []string
	max := 0
	for _, v := range word {
		if max < v {
			max = v
		}
	}
	fmt.Println(max)
	for k, v := range word {
		if v == max {
			a = append(a, k)
		}
	}
	fmt.Println("more frequent words", a)

}

//To get less frequent words
func lessFrequent(word map[string]int) {
	fmt.Println("less freq")
	var a []string
	min := 0
	for _, v := range word {
		min = v
		break
	}
	for _, v := range word {
		if min > v {
			min = v
		}
	}
	fmt.Println(min)
	for k, v := range word {
		if v == min {
			a = append(a, k)
		}
	}

	fmt.Println("more frequent words", a)

}

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

			json.Unmarshal(data, &wordMap)
			fmt.Println("Freq of all words in all files", wordMap)
		}
		fstatus, _ := cmd.Flags().GetBool("least")
		if fstatus { // if status is true, call least
			lessFrequent(wordMap)
		} else {
			moreFrequent(wordMap)
		}
	},
}

func init() {
	rootCmd.AddCommand(storefreqWordsCmd)
	storefreqWordsCmd.Flags().BoolP("least", "l", false, "Least Frequency Words")

}

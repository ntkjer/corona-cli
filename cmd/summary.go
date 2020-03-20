/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"net/http"
	"os/exec"

	"github.com/spf13/cobra"
)

const BASE_API = "https://coronavirus-19-api.herokuapp.com"
const API_TOTAL_ENDPOINT = BASE_API + "/all"
const API_ALL_COUNTRIES_ENDPOINT = BASE_API + "/countries"

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Summarizes COVID-19 stats.",
	Long: `Usage:

summary --country/-c US 
summary --country/-c China

Leaving summary without a country code flag entered will return a summary of all countries. 

To Plot:

`,
	Run: func(cmd *cobra.Command, args []string) {
		country, _ := cmd.Flags().GetString("country")
		curlCmd := "curl -s "
		//casesFilter = ` jq ".cases"`
		//deathsFilter = ` jq ".deaths"`
		//recoveredFilter = ` jq ".recovered"`
		if country == "" {
			//resp := initRequest(API_TOTAL_ENDPOINT)
			//body, err := ioutil.ReadAll(resp.Body)
			//filter := `jq -c '.[]' `
			cmd := curlCmd + API_ALL_COUNTRIES_ENDPOINT + ` | jq -s .`
			out, err := exec.Command("bash", "-c", cmd).Output()
			errorHandler(err)
			fmt.Println(string(out))
		} else if country != "" {
			cmd := curlCmd + API_ALL_COUNTRIES_ENDPOINT + "/" + country + " | jq -s ."
			out, err := exec.Command("bash", "-c", cmd).Output()
			errorHandler(err)
			res := string(out)
			fmt.Println(res)
		}

	},
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func initRequest(url string) *http.Response {
	resp, err := http.Get(url)
	errorHandler(err)
	defer resp.Body.Close()
	fmt.Println(resp)
	return resp
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Flags().StringP("country", "c", "", "Choose country, default=all")
	summaryCmd.Flags().StringP("flag", "f", "", "Flag ")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// summaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

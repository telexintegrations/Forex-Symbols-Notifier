package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//integration service

type TwelveDataResponse struct {
	Values []struct {
		Close string `json:"close"`
	} `json:"values"`
}

func TwelveDemo(from string, to string) (string){
	demoKey := os.Getenv("API_DEMO_KEY")
	url := fmt.Sprintf("https://api.twelvedata.com/time_series?symbol=%s/%s&interval=1day&apikey=%s", from, to, demoKey)
	fmt.Println("In twelvedemo start")
	var ex string
	if from == to {
		ex = "1"
	} else {
		// Make HTTP GET request
		res, err := http.Get(url)
		if err != nil {
			return err.Error()
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return err.Error()
		}

		var response TwelveDataResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println(err)
			return err.Error()
		}

		if len(response.Values) > 0 {
			_, err := fmt.Sscanf(response.Values[0].Close, "%s", &ex)
			if err != nil {
				fmt.Println(err)
				return err.Error()
			}
		} else {
			return "error"
		}
	}
	fmt.Println("In twelvedemo end")
	fmt.Println(ex)

	return ex
}

func PostToReturnURL(url string, data interface{}) (string, error) {
	fmt.Println("Posting to telex")
	// Marshal the data into JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %v", err)
	}
	fmt.Println("post data is..")
	fmt.Println(string(jsonData))
	// Make the POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	return string(body), nil
}
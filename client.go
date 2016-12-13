package gotsrpc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// CallClient calls a method on the remove service
func CallClient(url string, endpoint string, method string, args []interface{}, reply []interface{}) error {
	// Marshall args
	jsonArgs := []string{}
	for _, value := range args {
		jsonArg, err := json.Marshal(value)
		if err != nil {
			return err
		}
		jsonArgs = append(jsonArgs, string(jsonArg))
	}
	// Create request
	request := "[" + strings.Join(jsonArgs, ",") + "]"
	// Create post url
	postURL := fmt.Sprintf("%s%s/%s", url, endpoint, method)
	// Post
	resp, err := http.Post(postURL, "application/json", strings.NewReader(request))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Read in body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// Check status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: %s", resp.Status, string(body))
	}
	// Unmarshal reply
	if err := json.Unmarshal(body, &reply); err != nil {
		return err
	}
	return nil
}

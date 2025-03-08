package ssex

import (
	"bufio"
	"fmt"
	"net/http"
)

var client http.Client

func Connect(req *http.Request) (chan string, error) {
	message := make(chan string)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}
	go func() {
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			message <- line
		}
		close(message)
	}()

	return message, nil
}

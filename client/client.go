package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mdstaff/alpha-vantage-go/quote"
)

const (
	baseURL = "https://www.alphavantage.co/"
)

type Client struct {
	apiKey string
}

// Return a new alpha vantage client
func NewClient() (c *Client) {
	var env, err = readEnvFile()
	if err != nil {
		log.Fatal("Error in readEnvFile request %w", err)
	}
	fmt.Println("Env", env)
	c = &Client{
		apiKey: env["KEY"],
	}
	fmt.Println("Client", c)
	return
}

func (c *Client) Request(function string, params map[string]string) ([]byte, error) {
	var url = fmt.Sprintf(
		"%squery?function=%s%s&apikey=%s", 
		baseURL, 
		function, makeGetParams(params), c.apiKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error in Get request %v", err)
	}

	// Check response codes
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error with request, recieved status code: %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body %w", err)
	}
	response.Body.Close()
	return body, nil
}

func (c *Client) GetQuote(symbol string) (*quote.Quote, error) {
	params := map[string]string{
		"symbol": symbol,
	}
	body, err := c.Request("GLOBAL_QUOTE", params)
	if err != nil {
		return nil, err
	}
	res, err := quote.ParseQuote(body)
	fmt.Printf("Res %v", res)
	return res, err
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

// Declare the global variable declaration
var (
	ClienID           string
	ClientSecretID    string
	AccessToken       string
	AccessSecretToken string
	httpClient        *http.Client
)

// this function called first so initialize all variable with add values
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// load credential values
	ClienID = os.Getenv("TWITTER_CONSUMER_KEY")
	ClientSecretID = os.Getenv("TWITTER_CONSUMER_SECRET")
	AccessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	AccessSecretToken = os.Getenv("TWITTER_ACCESS_SECRET")

	// fmt.Println("Consumer Key:", ClienID)
	// fmt.Println("Consumer Secret:", ClientSecretID)
	// fmt.Println("Access Token:", AccessToken)
	// fmt.Println("Access Secret Token:", AccessSecretToken)

	if ClienID == "" || ClientSecretID == "" || AccessToken == "" || AccessSecretToken == "" {
		fmt.Println("Twitter API credentials are not set.")
		os.Exit(1)
	}

	// coonect the loaded credential with oauth1
	config := oauth1.NewConfig(ClienID, ClientSecretID)
	token := oauth1.NewToken(AccessToken, AccessSecretToken)
	httpClient = config.Client(oauth1.NoContext, token)
}

func main() {

	tweet := "Hello, Twitter World With GO!"
	err := postTweet(tweet)
	if err != nil {
		fmt.Println("Error posting tweet:", err)
	}

	fmt.Println("Tweet posted successfully! After 10 second it deleted ...")
	time.Sleep(10 * time.Second)

	err = deleteTweet(tweetDeleteID)
	if err != nil {
		fmt.Println("Error deleting tweet:", err)
		return
	}

	fmt.Println("Tweet deleted successfully!")
}

var tweetDeleteID string

func postTweet(tweetMsg string) error {
	apiURL := "https://api.twitter.com/2/tweets"

	// this for the convert tweet data for api format
	tweetData := map[string]interface{}{
		"text": tweetMsg,
	}
	jsonData, err := json.Marshal(tweetData)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)

	}

	// handle a POST request
	postRequest, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	postRequest.Header.Set("Content-Type", "application/json")
	postRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AccessToken))

	// Send the request
	handleResponse, err := httpClient.Do(postRequest)
	if err != nil {
		return fmt.Errorf("error posting tweet: %w", err)
	}

	var responseBody map[string]interface{}
	err = json.NewDecoder(handleResponse.Body).Decode(&responseBody)
	if err != nil {
		return fmt.Errorf("error parsing response body: %w", err)
	}

	// Fetch the tweetID from the response
	tweetID, ok := responseBody["data"].(map[string]interface{})["id"].(string)
	if !ok {
		return fmt.Errorf("error extracting tweet ID: %w", err)
	}
	tweetDeleteID = tweetID

	defer handleResponse.Body.Close()

	// check the respoce status code
	if handleResponse.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to post tweet, status code: %d", handleResponse.StatusCode)
	}

	fmt.Println("Tweet posted successfully!", tweetID)
	return nil

}

func deleteTweet(tweetID string) error {
	apiURL := fmt.Sprintf("https://api.twitter.com/2/tweets/%s", tweetID)

	// Handle a delete request
	deleteRequest, err := http.NewRequest("DELETE", apiURL, nil)
	if err != nil {
		return fmt.Errorf("error creating delete request: %w", err)
	}
	deleteRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AccessToken))

	// Send the request
	handleResponse, err := httpClient.Do(deleteRequest)
	if err != nil {
		return fmt.Errorf("error deleting tweet: %w", err)
	}
	defer handleResponse.Body.Close()

	// Check the response status code
	if handleResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete tweet, status code: %d", handleResponse.StatusCode)
	}

	fmt.Println("Tweet deleted successfully!")
	return nil
}

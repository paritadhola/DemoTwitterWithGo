## Introduction

This assignment focuses on interacting with the Twitter API using Go. It demonstrates how to post and delete tweets programmatically while introducing important concepts like API requests, authentication with OAuth1, and handling API responses. It teaches API integration, environment variable usage for storing sensitive credentials, and how to handle errors gracefully in Go.

## Setup Instructions

- Create a Twitter Developer Account: First, sign up for a Twitter Developer account. After approval, log in to the Twitter Developer portal.

  **Generate API Keys:**

- Apply for Twitter API access by providing basic information about the application and its intended use.
- Create a new project named "WorkOnTwitterAPIWithGoLang" in the developer portal.
- Twitter will provide an API Key, API Secret Key, Access Token, and Access Token Secret under the "Projects & Apps" section.

**Store Credentials:**

- Store these keys in a .env file. This file will hold the API credentials and ensure they are not hardcoded in the program.

## Program Details

**Posting a Tweet:**

- The program loads all Twitter API credentials from the .env file. These include the API key, API secret, Access token, and Access token secret.
- The program constructs a JSON object for the tweet body using json.Marshal to format the data.
- It then sends a POST request to Twitter’s API using the provided URL, including the necessary headers for authorization.
- Handle the response: The response from the API is handled, and if successful, the program extracts the tweet ID from the response body.

**Deleting a Tweet:**

- The tweet ID is extracted from the response of the postTweet function and stored in a global variable.
- To delete the tweet, the program constructs a DELETE request using the tweet ID and sends it to Twitter's API with the appropriate authorization headers.

## Examples of API requests and responses.

**For make request for post tweet:**

- API URL: https://api.twitter.com/2/tweets
- Method: POST
- Body: {“text": "Hello, Twitter World!”}
- For get response of the post request
- {“data": {“id": " 1844461446240825387", "text": "Hello, Twitter World!" } }

**For make request for delete tweet:**

- API URL: https://api.twitter.com/2/tweets/{tweetID}
- Method: DELETE
- Getting response of the post delete request : 200 OK.

## Error Handling:

- In the init() function, it checks for missing credentials and stops execution if any credentials are not found.
- The program also checks for failed requests and unsuccessful status codes. If any issues arise, such as bad responses, the program provides a clear error message and exits the function early to avoid further errors. This approach ensures smooth error handling and user feedback during execution.

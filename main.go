package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tankhiwale/ecFetch/logging"
	"github.com/tankhiwale/ecFetch/service"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func getClient(config *oauth2.Config) *http.Client {
  //specify the token file

  tokenFile := "token.json"
 
  token, err := tokenFromFile(tokenFile)

  if err != nil {
    token := tokenFromWeb(config)
    //saveToken(tokenFile, token)
    fmt.Println(token)
  }
  return config.Client(context.Background(), token)
}

func tokenFromFile(fileName string) (*oauth2.Token, error) {
  //returns an open file
  f, err := os.Open(fileName)

  if err != nil {
    log.Fatalf("error opening token file %v", fileName)
    return nil, err
  }
  
  defer f.Close()
  token := &oauth2.Token{}
  
  err = json.NewDecoder(f).Decode(token)

  return token, err

}

func tokenFromWeb(config *oauth2.Config) *oauth2.Token {
  //AccessTypeOffline is for when user is not accessing our app actively and we need to get a token behind the scene.
  //it gets us a refresh token, which helps us get an auth token in Background
  // refresh token is sent only on first request for getting access token.
  // in subsequent requests only access token and type are transmitted back.
  // also in app in in testing status in console.cloud.google then refresh token will expire more quickly.
  authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

  fmt.Printf("auth url : %v", authURL)
  
  var authCode string

  if _, err := fmt.Scan(&authCode); err != nil {
    log.Fatalf("Error scanning auth code : %v", err)
  }
  token, err := config.Exchange(context.TODO(), authCode)
  if err != nil {
    log.Fatalf("Error getting token from internet. config exchange : %v", err)
  } 

  return token
}

func saveToken(tokenFileName string, token string) error {
  return nil
}

func main() {

  loggingService := logging.NewLoggingService(&service.EmailFetcherImpl{})
  data, err := loggingService.FetchEmails(context.Background())

  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(data)
  
  //////////////////// gmail api sample code below ////////////////////  
  
  //context := context.Background()
  b, err := os.ReadFile("credentials.json")
  
  if err != nil {
    log.Fatal("Error reading credentials file", err)
  }

  config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)

  if err != nil {
    log.Fatalf("Error unable to parse configuration file %v", err)
  }
  
  fmt.Println(config.RedirectURL)

  client := getClient(config)
  fmt.Println(client)
}

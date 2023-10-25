package service

import (
	"context"
	"time"
)


// exported
type EmailFetcher interface {
	FetchEmails(context context.Context) (Message, error)
}


type Message struct {
	from string `json: from`
	to string `json: to`
	data string `json: data`
	timestamp time.Time `json: timestamp`
}

// not exported 
type EmailFetcherImpl struct {}

//This method is the concrete implementation of the fetching logic
func (e *EmailFetcherImpl) FetchEmails(c context.Context) (Message, error) {

	message := Message{
		from : "test",
		to : "test",
		data: "test",
		timestamp: time.Now(),
	}
  time.Sleep(time.Second * 3)
	return message , nil
}

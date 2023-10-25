package main

import (
	"context"
	"fmt"

	"github.com/tankhiwale/ecFetch/logging"
	"github.com/tankhiwale/ecFetch/service"
)

func main() {

  loggingService := logging.NewLoggingService(&service.EmailFetcherImpl{})
  data, err := loggingService.FetchEmails(context.Background())

  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(data)
}

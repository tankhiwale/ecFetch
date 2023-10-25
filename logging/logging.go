package logging

import (
	"context"
	"time"

	service "github.com/tankhiwale/ecFetch/service"
)

// this is tightly coupled with EmailFether. see if this can be decoupled.
// either accept a func and variadic args or look into generics
type loggingservice struct {
   next service.EmailFetcher 
}

func NewLoggingService(next service.EmailFetcher) service.EmailFetcher{
  return &loggingservice{
    next : next,
  }
}
func (s *loggingservice) FetchEmails(c context.Context) (message service.Message, err error) {

  defer func(startTime time.Time) {
    //  use start time here to log something. maybe use logrus or slog
  }(time.Now())

  return s.next.FetchEmails(c) // too tightly coupled
}

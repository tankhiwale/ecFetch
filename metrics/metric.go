package metrics

import (
	"context"

	"github.com/tankhiwale/ecFetch/service"
)



type metricservice struct {
  next service.EmailFetcher
}

// the mertic service has decorated the EmailFetcher service
func NewMetricService(next service.EmailFetcher) service.EmailFetcher {
  return &metricservice{
    next : next,
  }
}


func (s *metricservice) FetchEmails(c context.Context) (message service.Message, err error) {
  
  // implement logic for metric collection and export
  return s.next.FetchEmails(c)
}

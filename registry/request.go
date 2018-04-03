package registry

import (
	"errors"
	"time"

	"github.com/valyala/fasthttp"
)

const url = "https://rec-registry.gov.au/rec-registry/app/api/public-register/certificate-actions"

var (
	// ErrNoLatency is returned when the given date is later than equal to today.
	ErrNoLatency = errors.New("registry: no 1 day latency in the given date.")
)

// NewRequest returns a new fasthttp.Request given date to get certificate actions.
func NewRequest(date time.Time) (*fasthttp.Request, error) {
	// API has 1 day latency.
	year, month, day := time.Now().Date()
	if date.After(time.Date(year, month, day, 0, 0, 0, 0, date.Location())) {
		return nil, ErrNoLatency
	}
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url + "?date=" + date.Format("2006-01-02"))
	return req, nil
}

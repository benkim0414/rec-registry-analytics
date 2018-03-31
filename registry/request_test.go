package registry

import (
	"bytes"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/valyala/fasthttp"
)

func TestNewRequest(t *testing.T) {
	tests := []struct {
		desc string
		date time.Time
		req  *fasthttp.Request
		err  error
	}{
		{
			desc: "Request for today",
			date: time.Now(),
			req:  nil,
			err:  ErrNoLatency,
		},
		{
			desc: "Request for yesterday",
			date: time.Now().AddDate(0, 0, -1),
			req:  newTestRequest(time.Now().AddDate(0, 0, -1)),
			err:  nil,
		},
	}
	for _, test := range tests {
		req, err := NewRequest(test.date)
		if diff := cmp.Diff(req, test.req, cmp.Comparer(equalRequest)); diff != "" {
			t.Errorf("%s: after NewRequest, req differs: (-got +want)\n%s", test.desc, diff)
		}
		if diff := cmp.Diff(err, test.err, cmp.Comparer(equalError)); diff != "" {
			t.Errorf("%s: after NewRequest, err differs: (-got +want)\n%s", test.desc, diff)
		}
	}
}

func newTestRequest(date time.Time) *fasthttp.Request {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url + "?date=" + date.Format("2006-01-02"))
	return req
}

// equalRequest reports whether Requests x and y are considered equal.
func equalRequest(x, y *fasthttp.Request) bool {
	return x == nil && y == nil || x != nil && y != nil && bytes.Equal(x.RequestURI(), y.RequestURI())
}

// equalError reports whether errors x and y are considered equal.
func equalError(x, y error) bool {
	return x == nil && y == nil || x != nil && y != nil && x.Error() == y.Error()
}

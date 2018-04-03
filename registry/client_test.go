package registry

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/valyala/fasthttp"
)

func TestNewClient(t *testing.T) {
	want := &Client{
		hc: &fasthttp.Client{},
	}
	got := NewClient()
	if diff := cmp.Diff(got, want, cmp.Comparer(equalClient)); diff != "" {
		t.Errorf("after NewClient, Client differs: (-got +want)\n%s", diff)
	}
}

func TestDo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		v := q.Get("date")
		date, err := time.Parse("2006-01-02", v)
		if err != nil {
			t.Errorf(`date %s should be formatted "yyyy-MM-dd"`, v)
		}
		year, month, day := time.Now().Date()
		if date.After(time.Date(year, month, day, 0, 0, 0, 0, date.Location())) {
			t.Errorf("date %s should be earlier than today", v)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "Success", "result": []}`))
	}))
	defer ts.Close()

	client := NewClient()
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(ts.URL + "?date=" + time.Now().AddDate(0, 0, -1).Format("2006-01-02"))
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	got, want := resp.Body(), []byte(`{"status": "Success", "result": []}`)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("after Do, Body differs: (-got +want)\n%s", diff)
	}
}

func equalClient(x, y *Client) bool {
	return x == nil && y == nil || x != nil && y != nil && reflect.DeepEqual(x, y)
}

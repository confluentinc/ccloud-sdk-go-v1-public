package ccloud

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

var sdkUserAgent = fmt.Sprintf("ccloud-sdk-go-v1-public/v([0-9]+\\.?){3} \\(%s/%s; %s\\)", runtime.GOOS, runtime.GOARCH, runtime.Version())

func TestUserAgent(t *testing.T) {
	wantMatch := func(userAgent string) func(tt *testing.T, client *Client) {
		return func(tt *testing.T, client *Client) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Regexp(tt, "^"+userAgent+"$", r.Header.Get("User-Agent"))
			}))
			defer ts.Close()
			req, err := client.sling.Get(ts.URL).Request()
			require.NoError(tt, err)
			_, err = client.sling.Do(req, nil, nil)
			require.NoError(tt, err)
		}
	}
	tests := []struct {
		name   string
		args   *Params
		wantFn func(*testing.T, *Client)
	}{
		{
			name:   "adds default user-agent",
			args:   &Params{},
			wantFn: wantMatch(sdkUserAgent),
		},
		{
			name:   "can set custom user-agent",
			args:   &Params{UserAgent: "my-app/1.2.3"},
			wantFn: wantMatch("my-app/1.2.3 " + sdkUserAgent),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClient(tt.args)
			tt.wantFn(t, got)
		})
	}
}

package request

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/evcc-io/evcc/util"
)

type handler struct{}

func (*handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte(`3.14`))
}

func TestDoNil(t *testing.T) {
	srv := httptest.NewServer(new(handler))
	h := NewHelper(util.NewLogger("foo"))

	if err := h.GetJSON(srv.URL, nil); err != nil {
		t.Error(err)
	}

	if err := h.PostJSON(srv.URL, nil, nil); err != nil {
		t.Error(err)
	}
}

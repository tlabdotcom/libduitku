package libduitku

import (
	"fmt"
	"io"
	"strings"
)

// CoreDuitku is ...
type CoreDuitku struct {
	Client Client
}

// Call is used to call ...
func (gw *CoreDuitku) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("/%s", path)
	}
	path = gw.Client.Host + path
	return gw.Client.Call(method, path, header, body, v)
}

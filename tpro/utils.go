package tpro

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// ListenAndServe serves the application.
func ListenAndServe(port string, handler http.Handler) error {
	fmt.Println("Listening on:", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}

// MustGetEnv gets an environment variable or panics.
func MustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("%s missing", key))
	}
	return v
}

// DecodeJSON decodes *http.Request body into obj which needs to be a pointer
func DecodeJSON(r *http.Request, obj interface{}) ([]byte, error) {

	if obj == nil {
		return nil, nil
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return nil, err
	}

	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	// Unmarshal body into obj interface
	if err := json.Unmarshal(body, obj); err != nil {
		return nil, err
	}

	return body, nil
}

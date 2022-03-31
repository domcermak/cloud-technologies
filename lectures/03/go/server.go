package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func reqContext(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	str, ok := ctx.Value("my-auth-token").(string)
	if ok {
		w.Write([]byte(str))
	} else {
		fmt.Fprintf(w, "no auth token")
	}

}

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "my-auth-token", "token")
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func dump(w http.ResponseWriter, req *http.Request) {
	url := req.URL
	fmt.Fprintf(w, "method %v\nscheme %v\n host %v\n port %v\n path %v\n", req.Method, url.Scheme, url.Host, url.Port(), url.Path)
}

func store(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
	w.WriteHeader(http.StatusNoContent)
}

func sleeper(_ http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	select {
	case <-ctx.Done():
	case <-time.After(10 * time.Second):
	}
}

// for more routing options see https://github.com/gorilla/mux
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.Handle("/context", middlewareOne(http.HandlerFunc(reqContext)))
	http.HandleFunc("/dump", dump)
	http.HandleFunc("/store", store)
	http.HandleFunc("/sleeper", sleeper)

	// use DefaultServerMux when handler is nil
	panic(http.ListenAndServe(":8080", nil))
}

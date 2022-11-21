package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		defer r.Body.Close()

		sb := strings.Builder{}
		sb.WriteString(" ====== " + time.Now().Format(time.RFC3339))
		sb.Write(bytes)
		sb.WriteByte('\n')

		for k, v := range r.Header {
			sb.WriteString(k)
			sb.WriteByte(':')
			sb.WriteString(strings.Join(v, ","))
			sb.WriteByte('\n')
		}

		fmt.Println(sb.String())
	}

	http.HandleFunc("/v1/echo", h1)
	http.HandleFunc("/v1/some/echo", h1)

	// ngrok http 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

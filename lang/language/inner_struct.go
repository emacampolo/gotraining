package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type GreetService struct{}

func (GreetService) Greet(name string) string { return "hello " + name }

// curl http://localhost:8080 -d '{"name": "ema"}'
// START OMIT
func main() {
	http.HandleFunc("/", handler(GreetService{}))
	listenAndServe()
}

func handler(s GreetService) http.HandlerFunc {
	type request struct {
		Name string
	}
	type response struct {
		Greeting string `json:"greeting"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var input request
		decode(r, &input)

		greeting := s.Greet(input.Name)
		resp := response{
			Greeting: greeting,
		}

		encode(w, resp)
	}
}

// END OMIT
func listenAndServe() {
	log.Println("listening at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decode(r *http.Request, v interface{}) {
	_ = json.NewDecoder(r.Body).Decode(v)
}

func encode(w http.ResponseWriter, v interface{}) {
	_ = json.NewEncoder(w).Encode(v)
}

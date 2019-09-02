package main

import (
	`encoding/json`
	`log`
	`net/http`
)

//  curl http://localhost:8080 -d "{\"name\":\"emanuel\"}"
func main() {
	http.HandleFunc("/", handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler() http.HandlerFunc {
	type request struct {
		Name string
	}
	type response struct {
		Greeting string `json:"greeting"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var input request
		read(r, &input)

		// Delegate to service layer
		g := greet(input.Name)
		output := response{
			Greeting: g,
		}

		write(w, output)
	}
}

func greet(name string) string {
	return "hello " + name
}

func read(r *http.Request, i interface{}) {
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(i)
}

func write(w http.ResponseWriter, i interface{}) {
	_ = json.NewEncoder(w).Encode(i)
}

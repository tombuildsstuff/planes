package main

import "net/http"

// information handles `/`
func information(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
Hello.

The following routes exist:

* '/airports'
* '/airports/{code}'
* '/routes'
* '/routes/{originCode}/{destinationCode}'

Each API is a GET request and returns either a 200 OK with JSON, or a 400 Bad Request.

`))
}

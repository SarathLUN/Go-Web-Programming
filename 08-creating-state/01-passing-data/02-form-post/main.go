package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
<h1>Submit form with POST</h1>
<form method="post">
<input type="text" name="q" />
<input type="submit" />
</form>
<h1>Submit form with GET</h1>
<form method="GET">
<input type="text" name="q" />
<input type="submit" />
</form>
<br>
YOUR SUBMIT: `+v)
}

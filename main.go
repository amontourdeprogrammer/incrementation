package main
import (
  "log"
  "net/http"
  "fmt"
  "strconv"
  "sync"

)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
  mutex.Lock()
  counter++
  fmt.Fprintf(w, strconv.Itoa(counter))
  mutex.Unlock()
}

func main () {
  http.HandleFunc("/", echoString)

  http.HandleFunc("/increment", incrementCounter)

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

  http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hi")
  })
  log.Fatal(http.ListenAndServe(":8000", nil))
}

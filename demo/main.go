package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

const (
  port = ":4000"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, world!")
}

func GetHelloPHP(w http.ResponseWriter, r *http.Request) {
  // 機器內互打需使用 EXPOSE port
  resp, err := http.Get("http://nginx-service:80/helloworld.php")

  if err != nil {
    // handle error
    fmt.Fprintf(w, "res error")
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    fmt.Fprintf(w, "body error")
  }

  // fmt.Println(string(body))
  fmt.Fprintf(w, string(body))
}

func SayHello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello world! by go.")
}

func init() {
  http.HandleFunc("/getHelloPHP", GetHelloPHP)

  http.HandleFunc("/say-hello", SayHello)

  http.HandleFunc("/", HelloWorld)

  if err := http.ListenAndServe(port, nil); err != nil {
    panic(err)
  }
}

func main() {}

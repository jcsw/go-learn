package main

import (
  "fmt"
  "time"
  "net"
  "net/http"
  "io/ioutil"
)

func main() {

  var netTransport = &http.Transport{
  Dial: (&net.Dialer{
    Timeout: 5 * time.Second,
  }).Dial,
  TLSHandshakeTimeout: 5 * time.Second,
  }

  var netClient = &http.Client{
    Timeout: time.Second * 10,
    Transport: netTransport,
  }

  response, _ := netClient.Get("http://google.com")
  body, _ := ioutil.ReadAll(response.Body)

  fmt.Println("finished request body=", string(body))
}

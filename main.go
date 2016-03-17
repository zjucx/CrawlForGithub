package main

import(
  "fmt"
  "net/http"
  "io/ioutil"
  "regexp"
)

func main() {
  url := "https://github.com/zjucx"
  res, err := http.Get(url)
  if err != nil {
    fmt.Println("http transport error is: ", err)
  }
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println("ioutil error is: ", err)
  }
  defer res.Body.Close()
  //reg := regexp.MustCompile(`fill="(.*)"`)
  reg := regexp.MustCompile("fill\\=\"[\\S\\s]+?\"")
  //reg := regexp.MustCompile("\\<[\\S\\s]+?\\>")
  //reg := regexp.MustCompile(`<rect (.*)`)
  str := reg.FindAllString(string(body), -1)
  str = reg.ReplaceAllString(str, -1)
  //str := reg.FindAllString(string(body), -1)
  fmt.Println(str)
}

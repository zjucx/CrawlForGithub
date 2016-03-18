package main

import(
  "fmt"
  "net/http"
  "io/ioutil"
  "regexp"
  "encoding/json"
  "strings"
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
  bodystr := string(body)
  //reg := regexp.MustCompile(`fill="(.*)"`)
  reg := regexp.MustCompile("fill\\=\"[\\S\\s]+?\"")
  //reg := regexp.MustCompile("\\<[\\S\\s]+?\\>")
  //reg := regexp.MustCompile(`<rect (.*)`)
  tmp := reg.FindAllString(bodystr, -1)
  startindex := 0
  for index, color := range tmp{
    if color != "fill=\"#eeeeee\"" {
      startindex = index
      break
    }
  }
  colorstr := tmp[startindex:]

  reg = regexp.MustCompile("data-count\\=\"[\\S\\s]+?\"")
  tmp = reg.FindAllString(bodystr, -1)
  countstr := tmp[startindex:]

  reg = regexp.MustCompile("data-date\\=\"[\\S\\s]+?\"")
  tmp = reg.FindAllString(bodystr, -1)
  datestr := tmp[startindex:]

  jsondata := make(map[string]interface{})
  
  for index, color := range colorstr{
    jsondata["color"] = strings.Replace(strings.Replace(color, "fill=", "", -1), "\"", "", -1)
    jsondata["count"] = strings.Replace(strings.Replace(countstr[index], "data-count=", "", -1), "\"", "", -1)
    jsondata["date"] = strings.Replace(strings.Replace(datestr[index], "data-date=", "", -1), "\"", "", -1)
  }
   //str := reg.FindAllString(string(body), -1)
   result, _ := json.Marshal(jsondata)
  fmt.Println(string(result))
}

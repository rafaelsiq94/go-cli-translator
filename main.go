package main

import (
  "os"
  "bytes"
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "net/url"

  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file:", err)
    os.Exit(1)
  }

  key := os.Getenv("KEY")
  location := os.Getenv("LOCATION")
  endpoint := "https://api.cognitive.microsofttranslator.com/"
  uri := endpoint + "/translate?api-version=3.0"

  from := os.Args[1]
  to := os.Args[2]
  text := os.Args[3]

  u, _ := url.Parse(uri)
  q := u.Query()
  q.Add("from", from)
  q.Add("to", to)
  u.RawQuery = q.Encode()

  if len(os.Args) < 2 {
    fmt.Println("Please provide a string to translate.")
    os.Exit(1)
  }

  body := []struct {
    Text string
  }{
    {Text: text},
  }
  b, _ := json.Marshal(body)

  req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(b))
  if err != nil {
    log.Fatal(err)
  }
  req.Header.Add("Ocp-Apim-Subscription-Key", key)
  req.Header.Add("Ocp-Apim-Subscription-Region", location)
  req.Header.Add("Content-Type", "application/json")

  res, err := http.DefaultClient.Do(req)
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()

  var result []struct {
    Translations []struct {
      Text string `json:"text"`
      To   string `json:"to"`
    } `json:"translations"`
  }
  if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
    log.Fatal(err)
  }

  if len(result) > 0 && len(result[0].Translations) > 0 {
    fmt.Println(result[0].Translations[0].Text)
  } else {
    fmt.Println("Translation failed.")
  }
}

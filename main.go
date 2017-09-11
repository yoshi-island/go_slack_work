package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    //"fmt"
    "net/url"
    "github.com/ChimeraCoder/anaconda"
    "passwordlist"
)


func twitter_search() []string {
    anaconda.SetConsumerKey(passwordlist.ConsumerKey)
    anaconda.SetConsumerSecret(passwordlist.ConsumerSecret)
    api := anaconda.NewTwitterApi(passwordlist.AccessToken, passwordlist.AccessTokenSecret)

    v := url.Values{}
    v.Set("count", "10")
    searchResult, _ := api.GetSearch("#golang", v)

    twtlst := make([]string, 10)
    for i, tweet := range searchResult.Statuses {
        twtlst[i] = "\n\n" + tweet.CreatedAt + " " + tweet.Text + "\n\n"
        _ = twtlst //avoiding error
        //fmt.Printf("[%s] %s\n", tweet.CreatedAt, tweet.Text)
    }
   return twtlst
}


var (
    IncomingUrl string = passwordlist.Incomingurl
)

type Slack struct {
    Text        string `json:"text"`
    Username    string `json:"username"`
    Icon_emoji  string `json:"icon_emoji"`
    Icon_url    string `json:"icon_url"`
    Channel     string `json:"channel"`
}

func main() {

    output_text := ""
    output_text_2 := ""
    for _, ot := range twitter_search() {

      output_text_2 = output_text + ot
      output_text = output_text_2
      _ = output_text
    }

    //fmt.Print(output_text)

    params, _ := json.Marshal(Slack{
        output_text,
        passwordlist.Username,
        passwordlist.Icon_emoji,
        passwordlist.Icon_url,
        passwordlist.Channel})

    resp, _ := http.PostForm(
        IncomingUrl,
        url.Values{"payload": {string(params)}},
    )

    body, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()

    println(string(body))
}

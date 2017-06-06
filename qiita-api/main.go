package main

import (
	"log"
	"encoding/json"
	"net/http"
	"os"
	"io/ioutil"
)

type QiitaItem struct {
	Title string
}

func GetQiitaItems(token string) []QiitaItem {
	req, err := http.NewRequest("GET", "http://qiita.com/api/v2/authenticated_user/items", nil)
	if err != nil {
		log.Print(err)
	}
	req.Header.Set("Authorization", "Bearer " + token)

	cli := &http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		log.Print(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result []QiitaItem
	err = json.Unmarshal(body, &result)

	return result
}

func main() {
	token := os.Getenv("QIITA_ACCESS_TOKEN")
	items := GetQiitaItems(token)

	for _, item := range items {
		println(item.Title)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type QueryParams struct {
	Repository    string
	Project       string
	Tag           string
	TagLastPushed time.Time
	Digest        string
}

func (q *QueryParams) SetDefault() {
	if len(q.Repository) == 0 {
		q.Repository = "library"
	}
}

func (q *QueryParams) GetQueryURL() string {
	q.SetDefault()
	url := fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/%s/tags/%s", q.Repository, q.Project, q.Tag)
	fmt.Println(url)

	return url
}

func main() {
	params := []QueryParams{
		{Repository: "library", Project: "golang", Tag: "1.22-alpine"},
		{Repository: "library", Project: "node", Tag: "lts-alpine"},
	}

	res, err := http.Get(params[0].GetQueryURL())
	if err != nil {
		fmt.Printf("%v", err)
	}

	var result TagMeta

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Digest %s, LastPushed: %s", result.Digest, result.TagLastPushed)
}

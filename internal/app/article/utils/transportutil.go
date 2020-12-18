package utils

import (
	"blogger-kit/internal/pkg/requests"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

func DecodeArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body err, %v\n", err)
		return nil, err
	}
	var reg requests.Article
	if err = json.Unmarshal(body, &reg); err != nil {
		log.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}

	return &requests.Article{
		CategoryID: reg.CategoryID,
		Summary:    reg.Summary,
		Title:      reg.Title,
		UserName:   reg.UserName,
	}, nil
}

func DecodeArticleInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := r.URL.Query().Get("id")
	if id == "" {
		return nil, ErrorBadRequest
	}
	idInt, _ := strconv.Atoi(id)
	return &requests.ArticleInfo{
		ID: idInt,
	}, nil
}

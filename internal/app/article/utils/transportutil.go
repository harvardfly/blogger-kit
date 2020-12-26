package utils

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"pkg.zpf.com/golang/blogger-kit/internal/pkg/baseerror"

	"pkg.zpf.com/golang/blogger-kit/internal/pkg/requests"
)

var (
	// ErrorBadRequest 参数错误异常
	ErrorBadRequest = errors.New("invalid request parameter")
)

// DecodeArticleRequest Article请求参数验证
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
	// 验证必填参数
	err = baseerror.ParamError(reg)
	if err != nil {
		return nil, err
	}
	return &requests.Article{
		CategoryID: reg.CategoryID,
		Summary:    reg.Summary,
		Title:      reg.Title,
		UserName:   reg.UserName,
	}, nil
}

// DecodeArticleInfoRequest ArticleInfo请求参数验证
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

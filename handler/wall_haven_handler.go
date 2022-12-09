package handler

import (
	"encoding/json"
	"gin-toy/constant"
	"gin-toy/data"
	"gin-toy/util"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	. "github.com/ahmetb/go-linq/v3"
)

func SearchWallHavenPapers(search *data.Search) (*data.SearchResults, error) {
	resp, err := getWithValues("/search", search.ToQuery())
	if err != nil {
		return nil, err
	}
	log.Default().Printf("resp: %v", resp)
	out := &data.SearchResults{}
	processResponse(resp, out)
	urlList := make([]string, 0)
	From(out.Data).SelectT(func(paper data.Wallpaper) string {
		return paper.Path
	}).ToSlice(&urlList)
	go fetchUrl(urlList)
	return out, nil
}

func getWithValues(p string, v url.Values) (*http.Response, error) {
	u, err := url.Parse(constant.BaseURL + p)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	log.Default().Printf("url:%v", u)
	return getAuthedResponse(u.String())
}

func getAuthedResponse(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-Key", "LZARluDNtI4HBHeFBHmljupNNSFiME5t")
	log.Default().Printf("Req:%v", *req)
	return util.Do(req)
}

func fetchUrl(urlList []string) {
	log.Default().Printf("fetch url list:%v", urlList)
	for _, u := range urlList {
		resp, err := util.GET(u)
		if err != nil {
			log.Default().Panicf("err:%v", err)
			return
		}
		parsedURL, err := url.Parse(u)
		if err != nil {
			log.Default().Panicf("parse url err:%v", err)
			return
		}

		log.Default().Printf("parsed url:%v, url path:%v", parsedURL, parsedURL.Path)

		paths := strings.Split(parsedURL.Path, "/")

		download(".\\public\\imgs\\"+paths[len(paths)-1], resp)
	}
}

func download(filepath string, resp *http.Response) error {

	defer resp.Body.Close()

	out, err := os.Create(filepath)
	log.Default().Printf("out file: %s, err:%v", filepath, err)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func processResponse(resp *http.Response, out interface{}) error {
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return json.Unmarshal(byt, out)

}

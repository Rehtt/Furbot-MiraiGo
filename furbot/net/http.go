/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/8 14:31
 */

package net

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func httpGet(urlPath string, headers, query map[string]string) ([]byte, error) {
	Url, err := url.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	urlQuery := url.Values{}
	for k, v := range query {
		urlQuery.Add(k, v)
	}
	Url.RawQuery = urlQuery.Encode()

	req, _ := http.NewRequest("GET", Url.String(), nil)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	out, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return out, nil
}

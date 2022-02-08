/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/8 14:31
 */

package net

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func httpGet(urlPath string, headers, query map[string]string) ([]byte, error) {
	var url strings.Builder
	url.WriteString(urlPath)
	if len(query) != 0 {
		index := 0
		url.WriteByte('?')
		for k, v := range query {
			url.WriteString(k)
			url.WriteByte('=')
			url.WriteString(v)
			if index++; index < len(query) {
				url.WriteByte('&')
			}
		}
	}

	req, _ := http.NewRequest("GET", url.String(), nil)
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

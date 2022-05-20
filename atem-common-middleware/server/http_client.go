package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Get(url string) (response []byte, err error) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http status: %v %v", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//application/json; charset=utf-8
func Post(url string, data interface{}, contentType string) (content []byte, err error) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", contentType)
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// 下载图片信息
func DownLoad(savePath string, url string) (string, error) {
	saveName := savePath
	idx := strings.LastIndex(url, "/")
	if idx < 0 {
		saveName += "/" + url
	} else {
		saveName += url[idx:]
	}
	v, err := http.Get(url)
	if err != nil {
		fmt.Printf("Http get [%v] failed! %v", url, err)
		return "", err
	}
	defer v.Body.Close()
	content, err := ioutil.ReadAll(v.Body)
	if err != nil {
		fmt.Printf("Read http response failed! %v", err)
		return "", err
	}
	err = ioutil.WriteFile(saveName, content, 0666)
	if err != nil {
		fmt.Printf("Save to file failed! %v", err)
		return "", err
	}
	return saveName, nil
}

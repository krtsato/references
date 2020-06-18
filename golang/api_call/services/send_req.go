package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendRequest(param string) ([]byte, error) {
	apiUrl := "https://qiita.com/api/v2/items"

	// URL とクエリを用意する
	reqVals := url.Values{}
	reqVals.Add("page", param)
	reqVals.Add("per_page", "1")

	// リクエストを作成する
	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Errorf("Error: while creating URL")
		return nil, err
	}
	request.URL.RawQuery = reqVals.Encode()

	// リクエストを送信する
	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		fmt.Errorf("Error: while connecting with API")
		return nil, err
	}
	defer response.Body.Close()

	// 簡易的に status error を拾う
	if response.StatusCode != 200 {
		err := fmt.Errorf("Error: %s", response.Status)
		return nil, err
	}

	// 取得結果を変数に格納する
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("Error: while reading response body")
		return nil, err
	}

	return resBody, nil
}

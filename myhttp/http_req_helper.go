package myhttp

import (
	"io"
	"io/ioutil"
	"net/http"
)

// HttpRequest 发送http请求
// url
// method : GET,POST
// json参数： payload := strings.NewReader(`{"name": "xxx"}`)
// example: Request(url, "GET", body)
// body := strings.NewReader(`{"xxx": "xxx"}`)
func HttpRequest(url string, method string, body io.Reader) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}
	//req.Header.Add("token", "xxx")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(result))
	return result, nil
}

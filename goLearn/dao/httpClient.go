package dao

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	hiddenDataUrl = "http://xuhuipatrol.lmding.com/xuhuipatrol/anon/screen/rest/DataRest/getHiddenDangerData/v1_0?" +
		"starttime=2021-01-01%2000:00:00&endtime=2021-02-31%2000:00:00&source=YJJ-DK"

	reviewDataUrl = "http://xuhuipatrol.lmding.com/xuhuipatrol/anon/screen/rest/DataRest/getReviewData/v1_0?" +
		"starttime=2021-01-01%2000:00:00&endtime=2021-01-31%2000:00:00&source=YJJ-DK"
)

func HttpGet() (string, error) {
	resp, err := http.Get(reviewDataUrl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("+++++++++++" + string(body))
	return string(body), nil
}

func HttpPost() {
	resp, err := http.Post(reviewDataUrl, "application/x-www-form-urlencoded", strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

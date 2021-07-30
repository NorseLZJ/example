import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	httpClient = http.Client{Timeout: time.Second * 2}
)

func DoGet(url string) ([]byte, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}
	return body, err
}

func DoPost(url string, params url.Values) ([]byte, int, error) {
	var postReader io.Reader = nil
	if params != nil {
		postReader = strings.NewReader(params.Encode())

	}
	resp, err := httpClient.Post(url, "application/x-www-form-urlencoded", postReader)
	if err != nil {
		fmt.Printf("URL:%s err%s\n", url, err.Error())
		return nil, 0, err

	}
	code := resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("URL:%s err%s\n", url, err.Error())
		return nil, 0, err

	}
	return body, code, err
}


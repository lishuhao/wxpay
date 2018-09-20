package wxpay

import (
	"bytes"
	"io/ioutil"
	"net/http"
)


func HttpPost(url string,postBody []byte) ([]byte,error) {
	httpClient:=http.DefaultClient
	resp,err:=httpClient.Post(url,xmlContentType,bytes.NewReader(postBody))
	if err != nil {
		return nil,err
	}

	return ioutil.ReadAll(resp.Body)
}

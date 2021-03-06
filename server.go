package gopusher

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

//error if keys are not set
//  when? On instantiate or on Trigger()

//post
//get
//auth
//createSignedQueryString

type PusherHttpEndpoint struct {
	url string
}

func (this PusherHttpEndpoint) post(data []byte) (*string, error) {
	httpclient := &http.Client{}
	req, err := http.NewRequest("POST", this.url, bytes.NewBuffer(data))
	resp, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	body := string(responseData)
	return &body, err
}

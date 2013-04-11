package gopusher

import ()

type MockPusherEndpoint struct {
	postResult string
}

func (this MockPusherEndpoint) post(body []byte) (data *string, err error) {
	return &this.postResult, nil
}

package gopusher

import (
	"testing"
)

var (
	aChannel = "mychannel"
	aEvent   = "gopher-stolen"
	someData = struct {
		name    string
		company string
	}{
		"Pieter Joost van de Sande",
		"Wercker",
	}
)

type MockPoster struct {
	dataResult  string
	errorResult error

	sideEffect func(body []byte) (data *string, err error)
}

func (this MockPoster) post(body []byte) (data *string, err error) {
	sideEffect := this.sideEffect

	if sideEffect != nil {
		return sideEffect(body)
	}

	return &this.dataResult, err
}

func TestTriggerPostsCorrectJson(t *testing.T) {
	var postData []byte
	mockedPoster := &MockPoster{
		sideEffect: func(body []byte) (*string, error) {
			postData = body
			result := "ok"
			return &result, nil
		},
	}

	pusher := CreatePusher(aAppId, aAppKey, aAppSecret, mockedPoster, aMarshaller)

	pusher.Trigger(aChannel, aEvent, someData)

	println("data:")
	println(string(postData))
}

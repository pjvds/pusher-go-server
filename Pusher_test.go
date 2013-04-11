package gopusher

import (
	"testing"
)

var (
	aAppId     = "myappid"
	aAppKey    = "myappkey"
	aAppSecret = "myappsecret"
	aPoster    = Poster(nil)
)

func TestCreatePusherInitializesAppId(t *testing.T) {
	theAppId := "theappid"
	result := CreatePusher(theAppId, aAppKey, aAppSecret, aPoster)

	if result.appId != theAppId {
		t.Error("appId not initialized correctly")
	}
}

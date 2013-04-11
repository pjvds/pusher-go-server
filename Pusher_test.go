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

func TestCreatePusherInitializesAppKey(t *testing.T) {
	theAppKey := "theappkey"
	result := CreatePusher(aAppId, theAppKey, aAppSecret, aPoster)

	if result.appKey != theAppKey {
		t.Error("appKey not initialized correctly")
	}
}

func TestCreatePusherInitializesAppSecret(t *testing.T) {
	theAppSecret := "theappsecret"
	result := CreatePusher(aAppId, aAppKey, theAppSecret, aPoster)

	if result.appSecret != theAppSecret {
		t.Error("appSecret not initialized correctly")
	}
}

func TestCreatePusherInitializesPoster(t *testing.T) {
	thePoster := Poster(nil)
	result := CreatePusher(aAppId, aAppKey, aAppSecret, thePoster)

	if result.poster != thePoster {
		t.Error("poster not initialized correctly")
	}
}

package gopusher

import (
	"github.com/pjvds/pusher-go-server/serialization"
	"testing"
)

func TestCreatePusherInitializesAppId(t *testing.T) {
	theAppId := "theappid"
	result := CreatePusher(theAppId, aAppKey, aAppSecret, aPoster, aMarshaller)

	if result.appId != theAppId {
		t.Error("appId not initialized correctly")
	}
}

func TestCreatePusherInitializesAppKey(t *testing.T) {
	theAppKey := "theappkey"
	result := CreatePusher(aAppId, theAppKey, aAppSecret, aPoster, aMarshaller)

	if result.appKey != theAppKey {
		t.Error("appKey not initialized correctly")
	}
}

func TestCreatePusherInitializesAppSecret(t *testing.T) {
	theAppSecret := "theappsecret"
	result := CreatePusher(aAppId, aAppKey, theAppSecret, aPoster, aMarshaller)

	if result.appSecret != theAppSecret {
		t.Error("appSecret not initialized correctly")
	}
}

func TestCreatePusherInitializesPoster(t *testing.T) {
	thePoster := Poster(nil)
	result := CreatePusher(aAppId, aAppKey, aAppSecret, thePoster, aMarshaller)

	if result.poster != thePoster {
		t.Error("poster not initialized correctly")
	}
}

func TestCreatePusherInitializesMarshaller(t *testing.T) {
	theMarshaller := serialization.NewJsonMarshaller()
	result := CreatePusher(aAppId, aAppKey, aAppSecret, aPoster, theMarshaller)

	if result.marshaller != theMarshaller {
		t.Error("marshaller not initialized correctly")
	}
}

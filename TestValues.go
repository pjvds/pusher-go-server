package gopusher

import (
	"github.com/pjvds/pusher-go-server/serialization"
)

var (
	aAppId      = "myappid"
	aAppKey     = "myappkey"
	aAppSecret  = "myappsecret"
	aPoster     = &NullPoster{}
	aMarshaller = serialization.NewJsonMarshaller()
)

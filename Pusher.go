package gopusher

import (
	"github.com/pjvds/pusher-go-server/serialization"
)

// Represents the state of the Pusher context.
type Pusher struct {
	appId      string
	appKey     string
	appSecret  string
	poster     Poster
	marshaller serialization.Marshaller
}

func CreatePusher(appId, appKey, appSecret string, poster Poster, marshaller serialization.Marshaller) *Pusher {
	return &Pusher{
		appId:      appId,
		appKey:     appKey,
		appSecret:  appSecret,
		poster:     poster,
		marshaller: marshaller,
	}
}

func (p Pusher) Trigger(channel, event string, data interface{}) string {
	postData := struct {
		event string
		data  interface{}
	}{
		event,
		data,
	}

	payload, _ := p.marshaller.Marshal(postData)
	_, err := p.poster.post(payload)

	if err != nil {
		return err.Error()
	}
	return "ok"
}

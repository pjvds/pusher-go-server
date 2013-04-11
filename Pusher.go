package gopusher

import("github.com/pjvds/pusher-go-server/serialization")

// Represents the state of the Pusher context.
type Pusher struct {
	appId     string
	appKey    string
	appSecret string
	poster    Poster
    marshaller serialization.Marshaller
}

func CreatePusher(appId, appKey, appSecret string, poster Poster) *Pusher {
	return &Pusher{
		appId:     appId,
		appKey:    appKey,
		appSecret: appSecret,
		poster:    poster}
}

func (p Pusher) Trigger(channel, event string, data interface{}) string {

    map[string]interface{}
    return "ok"
}

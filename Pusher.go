package gopusher

// Represents the state of the Pusher context.
type Pusher struct {
	appId     string
	appKey    string
	appSecret string
	poster    Poster
}

func CreatePusher(appId, appKey, appSecret string, poster Poster) *Pusher {
	return &Pusher{
		appId:     appId,
		appKey:    appKey,
		appSecret: appSecret,
		poster:    poster}
}

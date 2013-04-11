package gopusher

type Poster interface {
	post(body []byte) (data *string, err error)
}

type NullPoster struct {
	result string
	err    error
}

func (this NullPoster) post(body []byte) (data *string, err error) {
	return &this.result, this.err
}

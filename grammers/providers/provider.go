package providers

type provider struct {
	Endpoint string
}

func (this provider) getEndPoint() string {
	return this.Endpoint
}

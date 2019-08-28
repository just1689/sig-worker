package domain

type RemoteImages struct {
	PageURL   string
	ImageURLs []string
}

func (r *RemoteImages) AddImageURL(u string) {
	r.ImageURLs = append(r.ImageURLs, u)
}

package url

import "github.com/google/uuid"

type UrlService struct {
}

var urlMapper map[string]string

func (u *UrlService) Shorten(url string) string {
	if urlMapper == nil {
		urlMapper = make(map[string]string)
	}
	if _, ok := urlMapper[url]; ok {
		return urlMapper[url]
	}
	shortUrl := uuid.New().String()
	urlMapper[shortUrl] = url
	return shortUrl
}

func (u *UrlService) GetUrl(shortUrl string) string {
	return urlMapper[shortUrl]
}

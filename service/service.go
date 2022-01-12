package service

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"strings"
)

const httpPrefix = "http://"

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) processUrl(url string) string {
	if !strings.Contains(url, httpPrefix) {
		url = strings.Join([]string{httpPrefix, url}, "")
	}

	return url
}

func (s *Service) get(path string) (*http.Response, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", s.processUrl(path), nil)
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}

func (s *Service) GetHash(postfix string) (string, error) {
	res, err := s.get(postfix)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	buf := make([]byte, 30*1024)
	md5H := md5.New()
	for {
		var n int
		n, err = res.Body.Read(buf)
		if n > 0 {
			_, err = md5H.Write(buf[:n])
			if err != nil {
				return "", err
			}
			if err == io.EOF {
				break
			}
		}
		if err != nil {
			break
		}

	}

	sum := md5.Sum(nil)
	return hex.EncodeToString(sum[:]), nil
}

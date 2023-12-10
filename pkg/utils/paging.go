package utils

import (
	"errors"
)

type Page struct {
	PageSize   int32
	PageOffset int32
}

func PageFromToken(pageToken string) (*Page, error) {
	var result Page
	if err := DecodeFromBase64(&result, pageToken); err != nil {
		return nil, err
	}
	return &result, nil
}

func (page *Page) GetToken() (string, error) {
	if page == nil {
		return "", errors.New("page is nil")
	}

	token, err := EncodeToBase64(*page)
	if err != nil {
		return "", errors.New("error encoding page")
	}
	return token, nil
}

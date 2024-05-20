package src

import (
	b64 "encoding/base64"
	"encoding/json"
)

type pageParams struct {
	PageSize   int32
	PageOffset int32
}

func constructPageToken(pageSize, pageOffset int32) (string, error) {
	bytes, err := json.Marshal(pageParams{
		PageSize:   pageSize,
		PageOffset: pageOffset,
	})
	if err != nil {
		return "", err
	}

	return b64.StdEncoding.EncodeToString(bytes), nil
}

func deconstructPageToken(pageToken string) (int32, int32, error) {
	var pageParams pageParams
	err := json.Unmarshal([]byte(pageToken), &pageParams)
	if err != nil {
		return 0, 0, err
	}
	return pageParams.PageSize, pageParams.PageOffset, nil
}

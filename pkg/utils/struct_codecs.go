package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"strings"
)

func EncodeToBase64(v interface{}) (encoded string, err error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	defer func(encoder io.WriteCloser) {
		encoderCloseErr := encoder.Close()
		if encoderCloseErr != nil {
			err = errors.Join(err, encoderCloseErr)
		}
	}(encoder)

	err = json.NewEncoder(encoder).Encode(v)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func DecodeFromBase64(v interface{}, enc string) error {
	return json.NewDecoder(base64.NewDecoder(base64.StdEncoding, strings.NewReader(enc))).Decode(v)
}

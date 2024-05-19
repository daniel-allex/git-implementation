package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"io"
)

func ZlibDecompress(compressed string) (string, error) {
	data := bytes.NewReader([]byte(compressed))
	r, err := zlib.NewReader(data)
	if err != nil {
		return "", err
	}

	out, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	err = r.Close()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func ZlibCompress(s string) (string, error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	_, err := w.Write([]byte(s))

	if err != nil {
		return "", err
	}

	err = w.Close()

	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func Sha1Hash(s string) string {
	hash := sha1.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

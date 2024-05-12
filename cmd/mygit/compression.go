package main

import (
	"bytes"
	"compress/zlib"
	"io"
)

func decompress(shah string) string {
	data := bytes.NewReader([]byte(shah))
	r, err := zlib.NewReader(data)
	exceptIfError("Failed to create zlib compression reader", err)

	defer r.Close()

	var out bytes.Buffer
	_, err = io.Copy(&out, r)

	return out.String()
}

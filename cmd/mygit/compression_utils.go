package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"io"
)

func ZlibDecompress(compressed string) string {
	data := bytes.NewReader([]byte(compressed))
	r, err := zlib.NewReader(data)
	ExceptIfError("Failed to create zlib decompression reader", err)

	defer r.Close()

	var out bytes.Buffer
	_, err = io.Copy(&out, r)

	return out.String()
}

func ZlibCompress(s string) string {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	_, err := w.Write([]byte(s))
	ExceptIfError("Failed to write using zlib compression writer", err)

	err = w.Close()
	ExceptIfError("Failed to close zlib writer", err)
	return b.String()
}

func Sha1Hash(s string) string {
	hash := sha1.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

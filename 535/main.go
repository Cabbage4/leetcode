package main

import "encoding/base64"

func main() {
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) encode(longUrl string) string {
	return base64.StdEncoding.EncodeToString([]byte(longUrl))
}

func (c *Codec) decode(shortUrl string) string {
	b, _ := base64.StdEncoding.DecodeString(shortUrl)
	return string(b)
}

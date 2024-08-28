package main

import "strconv"

type Codec struct {
	encodeMap map[string]string
	decodeMap map[string]string
	base      string
}

func Constructor() Codec {
	return Codec{
		encodeMap: make(map[string]string),
		decodeMap: make(map[string]string),
		base:      "http://tinyurl.com/",
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	if shortUrl, exists := this.encodeMap[longUrl]; exists {
		return shortUrl
	}

	shortUrl := this.base + strconv.Itoa(len(this.encodeMap)+1)
	this.encodeMap[longUrl] = shortUrl
	this.decodeMap[shortUrl] = longUrl

	return shortUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.decodeMap[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

package godom

import (
	"math/rand"
	"time"
)

// RandomProxy : 随机返回一个proxy 字符串
func RandomProxy() string {
	rand.Seed(time.Now().UnixNano())
	proxy := ipPool[rand.Intn(len(ipPool))]

	return proxy
}

var ipPool = []string{
	"http://127.0.0.1:8080",
	"socks5://127.0.0.1:1080",
}

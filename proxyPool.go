package godom

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/GuNanHai/toolkit"
)

// RandomProxy : 随机返回一个proxy 字符串
func RandomProxy() string {
	rand.Seed(time.Now().UnixNano())
	proxy := IPPool[rand.Intn(len(IPPool))]

	return proxy
}

func fetchProxyList() []string {
	proxyAPI := "https://guhaiproxy.tk/proxies.json"
	resp, err1 := http.Get(proxyAPI)
	toolkit.CheckErr(err1)
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	toolkit.CheckErr(err2)

	ipPool, err3 := toolkit.ReadStringListFromJSON(body)
	toolkit.CheckErr(err3)
	return ipPool
}

//IPPool : 代理IP列表
var IPPool = fetchProxyList()

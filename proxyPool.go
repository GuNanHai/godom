package godom

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/GuNanHai/toolkit"
)

// RandomProxy : 随机返回一个proxy 字符串
func RandomProxy() string {
	rand.Seed(time.Now().UnixNano())
	proxy := IPPool[rand.Intn(len(IPPool))]

	return proxy
}

func readProxyListFromLocal() []string {
	file := toolkit.GetPkgPath("godom") + "/" + "proxy.txt"
	f, err := os.Open(file)
	toolkit.CheckErr(err)
	fileContent, err2 := ioutil.ReadAll(f)
	toolkit.CheckErr(err2)

	proxyList := strings.Split(string(fileContent), "\n")

	return proxyList[:len(proxyList)-1]
}

// IPPool :
var IPPool = readProxyListFromLocal()

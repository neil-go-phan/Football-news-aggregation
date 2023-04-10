package crawlerhelpers

import (
	"io"
	"log"
	"net/http"
	"strings"
)

// api source: https://docs.proxyscrape.com/
var PROCY_API = "https://api.proxyscrape.com/v2/?request=displayproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all"

func RequestProxyList() (proxyList []string, err error) {
	resp, err := http.Get(PROCY_API)
	if err != nil {
		log.Println("can not get proxy from url:",PROCY_API)
		return proxyList, err
	}
	defer resp.Body.Close()

	responseData, _ := io.ReadAll(resp.Body)
	proxyList = strings.Split(string(responseData), "\r\n")

	return proxyList, nil
}

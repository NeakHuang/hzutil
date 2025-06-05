// Description: http client相关处理
// Datetime:  2025-05-26 23:47
package hzhttp

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func BuildProxy(proxyStr string) func(r *http.Request) (*url.URL, error) {
	if len(proxyStr) <= 0 {
		return nil
	}
	// 判断是否有前缀
	// 无前缀，则默认使用http前缀
	if !strings.HasPrefix(proxyStr, "socks5://") &&
		!strings.HasPrefix(proxyStr, "https://") &&
		!strings.HasPrefix(proxyStr, "http://") {
		proxyStr = fmt.Sprintf("http://%s", proxyStr)
	}
	if proxyUrl, err := url.Parse(proxyStr); err != nil {
		return func(*http.Request) (*url.URL, error) {
			return nil, errors.New(fmt.Sprintf("Analysic proxyUrl[%v] to proxy url fail %v", proxyStr, err))
		}
	} else {
		return http.ProxyURL(proxyUrl)
	}
}

func SetProxy(proxyStr string) (*http.Transport, error) {
	if len(proxyStr) <= 0 {
		return nil, nil
	}
	// 判断是否有前缀
	// 无前缀，则默认使用http前缀
	if !strings.HasPrefix(proxyStr, "socks5://") &&
		!strings.HasPrefix(proxyStr, "https://") &&
		!strings.HasPrefix(proxyStr, "http://") {
		proxyStr = fmt.Sprintf("http://%s", proxyStr)
	}
	if proxyUrl, err := url.Parse(proxyStr); err != nil {
		return nil, errors.New(fmt.Sprintf("Analysic proxyUrl[%v] to proxy url fail %v", proxyStr, err))
	} else {
		proxy := http.ProxyURL(proxyUrl)
		//HttpClient.Transport = &http.Transport{Proxy: proxy}
		return &http.Transport{Proxy: proxy}, nil
	}
}

// CheckProxy 检查代理
func CheckProxy(client *http.Client, url ...string) error {
	ipUrl := "http://myip.ipip.net"
	if len(url) > 0 {
		ipUrl = url[0]
	}
	req, err := http.NewRequest("GET", ipUrl, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Set("Cookie", "_safe=vqd37pjm4p5uodq339yzk6b7jdt6oich; cPNj_2132_saltkey=a5ffFKLd; cPNj_2132_lastvisit=1691752561; cPNj_2132_atarget=1; cPNj_2132_lastfp=6a9c7553e07b9f17e0dfa4f579361059; cPNj_2132_home_diymode=1; cPNj_2132_st_p=0%7C1692189681%7Cda20216937156c4d30b129e0d1859b42; cPNj_2132_viewid=tid_1493743; cPNj_2132_visitedfid=2D141D95D103D155D50D104D125D137D117; cPNj_2132_st_t=0%7C1692237351%7C5d8bcdb21881240dea376ba85b69c733; cPNj_2132_forum_lastvisit=D_152_1689677645D_142_1689771242D_139_1691155745D_37_1691493914D_41_1691756168D_154_1691756172D_36_1691773699D_43_1691773703D_48_1691773707D_49_1691773720D_165_1691773721D_96_1691774864D_145_1691817372D_146_1691829223D_155_1691835420D_109_1691836159D_143_1691836161D_103_1691911090D_117_1691927224D_137_1691927238D_125_1691927977D_104_1691927989D_50_1692067574D_95_1692194604D_141_1692237349D_2_1692237351; cPNj_2132_lastact=1692261728%09index.php%09")

	// 发起 HTTP 请求
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("Response:", string(body))
	return nil
}

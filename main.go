package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	// 从环境变量读取端口号
	port := os.Getenv("PR_PORT")
	defaultPort := 2445

	// 如果端口为空或解析失败，使用默认值
	if port == "" {
		fmt.Printf("PR_PORT 未设置，使用默认端口 %d\n", defaultPort)
		port = strconv.Itoa(defaultPort)
	} else {
		parsedPort, err := strconv.Atoi(port)
		if err != nil || parsedPort < 1 || parsedPort > 65535 {
			fmt.Printf("PR_PORT=%s 无效，使用默认端口 %d\n", port, defaultPort)
			port = strconv.Itoa(defaultPort)
		}
	}

	http.HandleFunc("/", redirectHandler)

	addr := ":" + port
	fmt.Printf("服务器启动，监听端口: %s\n", port)
	fmt.Printf("访问示例: http://localhost%s/?url=http%%3A%%2F%%2Fwww.baidu.com%%2Fs%%3Fwd%%3Dtest\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("服务器启动失败: %v\n", err)
		os.Exit(1)
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// 只处理 GET 请求
	if r.Method != http.MethodGet {
		http.Error(w, "只支持 GET 请求", http.StatusMethodNotAllowed)
		return
	}

	// 获取 url 参数
	encodedURL := r.URL.Query().Get("url")
	if encodedURL == "" {
		http.Error(w, "缺少 url 参数", http.StatusBadRequest)
		return
	}

	// 限制 URL 长度
	const maxURLLength = 2048
	if len(encodedURL) > maxURLLength {
		http.Error(w, "URL 过长", http.StatusBadRequest)
		return
	}

	// 解码 URL
	targetURL, err := url.QueryUnescape(encodedURL)
	if err != nil {
		http.Error(w, "URL 格式无效", http.StatusBadRequest)
		return
	}

	// 验证协议，仅允许 http 和 https
	parsedURL, err := url.Parse(targetURL)
	if err != nil || parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		http.Error(w, "仅支持 http/https 协议", http.StatusBadRequest)
		return
	}

	// 返回 302 重定向
	http.Redirect(w, r, targetURL, http.StatusFound)
}

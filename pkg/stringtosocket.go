package pkg

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

// 将字符串尽可能转化成套接字的二元组即ip和port
func StringToSocket(str string) (string, int, error) {
	str = strings.Trim(str, " \n\t")
	// 获取子串的位置
	httpsPos := strings.Index(str, "https://")
	httpPos := strings.Index(str, "http://")
	if httpsPos == -1 && httpPos == -1 {
		str = "http://" + str
	}
	u, err := url.Parse(str)
	if err != nil {
		return "", 0, err
	}
	ip := u.Hostname()
	port := u.Port()
	scheme := u.Scheme
	if ip == "" {
		return "", 0, errors.New("unable to resolve IP")
	}
	if port == "" {
		if scheme == "https" {
			return ip, 443, nil
		}
		return ip, 80, nil
	}
	return ip, stringToInt(port), err
}

func stringToInt(str string) int {
	num, err := strconv.Atoi(str)
	CheckErr(err)
	return num
}

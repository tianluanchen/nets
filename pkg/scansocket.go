package pkg

import (
	"fmt"
	"net"
	"time"
)

// 扫描指定端口判断端口是否开启
func ScanSocket(ip string, port int, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func StandardScanSocket(ip string, port int) bool {
	return ScanSocket(ip, port, time.Second*3)
}

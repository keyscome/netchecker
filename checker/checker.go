// checker/checker.go
package checker

import (
	"net"
	"time"
)

// CheckConnection 尝试在指定超时时间内建立 TCP 连接
// 成功返回 nil，失败返回错误信息
func CheckConnection(address string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}

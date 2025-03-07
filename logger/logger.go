// logger/logger.go
package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// NewLoggers 创建三个 Logger：
//   - successLogger：成功连接日志，输出到 stdout 和文件 logs/$timestamp/success.log
//   - failureLogger：失败连接日志，输出到 stderr 和文件 logs/$timestamp/failure.log
//   - reportLogger：检测报告日志，输出到 stdout 和文件 logs/$timestamp/report.log
//
// 同时返回 cleanup 函数用于关闭文件句柄
func NewLoggers() (successLogger, failureLogger, reportLogger *log.Logger, cleanup func(), err error) {
	// 使用时间戳创建日志目录，如 logs/20250305164541
	timestamp := time.Now().Format("20060102150405")
	logsDir := fmt.Sprintf("logs/%s", timestamp)
	err = os.MkdirAll(logsDir, 0755)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	successFilename := fmt.Sprintf("%s/success.log", logsDir)
	failureFilename := fmt.Sprintf("%s/failure.log", logsDir)
	reportFilename := fmt.Sprintf("%s/report.log", logsDir)

	successFile, err := os.Create(successFilename)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	failureFile, err := os.Create(failureFilename)
	if err != nil {
		successFile.Close()
		return nil, nil, nil, nil, err
	}

	reportFile, err := os.Create(reportFilename)
	if err != nil {
		successFile.Close()
		failureFile.Close()
		return nil, nil, nil, nil, err
	}

	// 成功日志同时写入 stdout
	successWriter := io.MultiWriter(os.Stdout, successFile)
	// 失败日志写入 stderr
	failureWriter := io.MultiWriter(os.Stderr, failureFile)
	// 检测报告日志同时写入 stdout
	reportWriter := io.MultiWriter(os.Stdout, reportFile)

	successLogger = log.New(successWriter, "SUCCESS: ", log.LstdFlags)
	failureLogger = log.New(failureWriter, "FAILURE: ", log.LstdFlags)
	reportLogger = log.New(reportWriter, "REPORT: ", log.LstdFlags)

	cleanup = func() {
		successFile.Close()
		failureFile.Close()
		reportFile.Close()
	}

	return successLogger, failureLogger, reportLogger, cleanup, nil
}

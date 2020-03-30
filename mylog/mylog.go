package mylog

import (
	"fmt"
	"time"
)
// 日志工具库

// Logger 日志接口
type Logger struct {
}

// NewLog 创建log对象
func NewLog() Logger {
	return Logger{}
}

// Info 打印 info 日志
func (l Logger) Info(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Info] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

// Debug 打印 debug 日志
func (l Logger) Debug(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Debug] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

// Error 打印 error 日志
func (l Logger) Error(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Error] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}
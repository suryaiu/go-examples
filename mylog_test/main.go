package main

import "example.com/demo/mylog"

func main() {
	log := mylog.NewLog()
	log.Info("info")
	log.Error("open file failed")
}

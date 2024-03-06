package main

import (
	"context"
	"fmt"
	"github.com/smiling77877/coredemo/framework"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	// 这个channel负责通知结束
	finish := make(chan struct{}, 1)
	// 这个channel负责通知panic异常
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(time.Second))
	defer cancel()

	go func() {
		//这里增加异常处理
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		//这里做具体的业务
		time.Sleep(10 * time.Second)
		c.Json(200, "ok")
		//新的goroutine结束的时候通过一个finish通道告知父goroutine
		finish <- struct{}{}
	}()

	select {
	// 监听 panic
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
		// 监听结束时间
	case <-finish:
		fmt.Println("finish")
		// 监听超时时间
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "time out")
		c.SetHasTimeout()
	}

	return nil
}

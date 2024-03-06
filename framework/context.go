package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// 自定义Context
type Context struct {
	request        *http.Request
	responsewriter http.ResponseWriter
	ctx            context.Context
	handler        ControllerHandler

	// 是否超时标记位
	hasTimeout bool
	// 写保护机制
	writerMux *sync.Mutex
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responsewriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
	}
}

// base 负责封装基本的函数功能
func (ctx *Context) WriterMux() *sync.Mutex { return ctx.writerMux }

func (ctx *Context) GetRequest() *http.Request { return ctx.request }

func (ctx *Context) GetResponse() http.ResponseWriter { return ctx.responsewriter }

func (ctx *Context) SetHasTimeout() { ctx.hasTimeout = true }

func (ctx *Context) HasTimeout() bool { return ctx.hasTimeout }

// context 实现标准Context接口
func (ctx *Context) BaseContext() context.Context { return ctx.request.Context() }

func (ctx *Context) Deadline() (deadline time.Time, ok bool) { return ctx.BaseContext().Deadline() }

func (ctx *Context) Done() <-chan struct{} { return ctx.BaseContext().Done() }

func (ctx *Context) Err() error { return ctx.BaseContext().Err() }

func (ctx *Context) Value(key any) any { return ctx.BaseContext().Value(key) }

// request  封装了http.Request的对外接口
func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.URL.Query()
	}
	return map[string][]string{}
}

func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

func (ctx *Context) Queryint(key string, def int) int {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			intval, err := strconv.Atoi(vals[len-1])
			if err != nil {
				return def
			}
			return intval
		}
	}
	return def
}

func (ctx *Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return vals[len-1]
		}
	}
	return def
}

func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.PostForm
	}
	return map[string][]string{}
}

func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

func (ctx *Context) FormInt(key string, def int) int {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			intval, err := strconv.Atoi(vals[len-1])
			if err != nil {
				return def
			}
			return intval
		}
	}
	return def
}

func (ctx *Context) FormString(key string, def string) string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return vals[len-1]
		}
		return def
	}
	return def
}

func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request != nil {
		body, err := io.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		ctx.request.Body = io.NopCloser(bytes.NewBuffer(body))

		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

// response  封装了http.ResponseWriter的对外接口
func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.responsewriter.Header().Set("Content-Type", "application/json")
	ctx.responsewriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responsewriter.WriteHeader(500)
	}
	ctx.responsewriter.Write(byt)
	return nil
}

func (ctx *Context) HTML(status int, obj interface{}, template string) error { return nil }

func (ctx *Context) Text(status int, obj string) error { return nil }

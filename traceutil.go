package scalpel

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"runtime"
	"time"
)

const (
	trcid  = "traceid"
	spid   = "spanid"
	method = "method"
	code   = "code"
)

func GetTraceId(r *http.Request) string {
	if r != nil {
		traceId := r.Header.Get("X-Trace-ID")
		if len(traceId) != 0 {
			return traceId
		}
	}

	now := time.Now()
	return now.Format("060102150405X") + fmt.Sprint((now.UnixNano()/1000)%10000000)
}

func GenTraceId(ctx context.Context, r *http.Request) string {
	if c := ctx.Value(trcid); c != nil {
		return fmt.Sprintf("%v", c)
	}

	if r != nil {
		traceId := r.Header.Get("X-Trace-ID")
		if len(traceId) != 0 {
			return traceId
		}
	}

	now := time.Now()
	return now.Format("060102150405X") + fmt.Sprint((now.UnixNano()/1000)%10000000)
}

func GetSpanId() string {
	return fmt.Sprint((time.Now().UnixNano() / 1000) % 10000000)
}

func GetErrCode(ctx context.Context) string {
	if c := ctx.Value(code); c != nil {
		return fmt.Sprintf("%v", c)
	}
	return "1"
}

func GetTrace(ctx context.Context) string {
	return fmt.Sprintf("traceid=%v||spanid=%v", ctx.Value(trcid), ctx.Value(spid))
}

func TracePanic(err interface{}) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\n", err)

	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buf.String()
}

func SetCtx(ctx context.Context, k interface{}, v interface{}) context.Context {
	return context.WithValue(ctx, k, v)
}

func SetTraceId(ctx context.Context, v interface{}) context.Context {
	return SetCtx(ctx, trcid, v)
}

func SetSpanId(ctx context.Context, v interface{}) context.Context {
	return SetCtx(ctx, spid, v)
}

func SetMethod(ctx context.Context, v interface{}) context.Context {
	return SetCtx(ctx, method, v)
}

func SetLang(ctx context.Context, v interface{}) context.Context {
	return SetCtx(ctx, method, v)
}

/**
获取上游ip
*/
func GetClientIp(r *http.Request) string {
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	} else {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
		if ip == `::1` {
			return "127.0.0.1"
		}
		return ip
	}
}

package scalpel

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func HttpInLog(ctx context.Context, r *http.Request, svc string, p interface{}) string {
	f := `%v||service=%s||tag=_http_request_in||ip=%v||uri=%v||method=%v||%+v`
	return fmt.Sprintf(f, GetTrace(ctx), svc, GetClientIp(r), r.URL.Path, r.Method, p)
}

func HttpOutLog(ctx context.Context, r *http.Request, svc string, resp []byte, t time.Time) string {
	c := fmt.Sprint(1) // unknown error code
	var m map[string]interface{}
	if len(resp) > 0 {
		json.Unmarshal(resp, &m)
	}
	if v, ok := m[code]; ok {
		c = fmt.Sprint(v)
	}

	l := len(resp)
	if l > 1000 {
		l = 1000
	}
	f := `%v||service=%s||tag=_http_request_out||code=%v||uri=%v||latency=%dms||%v`
	return fmt.Sprintf(f, GetTrace(ctx), svc, c, r.URL.Path, time.Since(t).Milliseconds(), string(resp[0:l]))
}

func CallInLog(ctx context.Context, callee string, m string, ep string, p interface{}) string {
	f := `%v||tag=_call_in_%s||%s endpoint=%s||%+v`
	return fmt.Sprintf(f, GetTrace(ctx), callee, m, ep, p)
}

func CallOutLog(ctx context.Context, callee string, resp []byte, err error, t time.Time) string {
	l := len(resp)
	if l > 500 {
		l = 500
	}
	f := `%v||tag=_call_out_%s||latency=%dms||resp=%s||err=%v`
	return fmt.Sprintf(f, GetTrace(ctx), callee, time.Since(t).Milliseconds(), resp[0:l], err)
}

package lib

import (
	"fmt"
	slog "github.com/liuzw3018/saber/log"
	"strings"
)

// 通用DLTag常量定义
const (
	SRTagUndefind      = "_undef"
	SRTagMySqlFailed   = "_com_mysql_failure"
	SRTagRedisFailed   = "_com_redis_failure"
	SRTagMySqlSuccess  = "_com_mysql_success"
	SRTagRedisSuccess  = "_com_redis_success"
	SRTagThriftFailed  = "_com_thrift_failure"
	SRTagThriftSuccess = "_com_thrift_success"
	SRTagHTTPSuccess   = "_com_http_success"
	SRTagHTTPFailed    = "_com_http_failure"
	SRTagTCPFailed     = "_com_tcp_failure"
	SRTagRequestIn     = "_com_request_in"
	SRTagRequestOut    = "_com_request_out"
)

const (
	_srTag          = "srtag"
	_traceId        = "traceid"
	_spanId         = "spanid"
	_childSpanId    = "cspanid"
	_srTagBizPrefix = "_com_"
	_srTagBizUndef  = "_com_undef"
)

var Log *Logger

type Logger struct{}

func (l *Logger) TagInfo(trace *TraceContext, srtag string, m map[string]interface{}) {
	m[_srTag] = checkDLTag(srtag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	slog.Info(parseParams(m))
}

func (l *Logger) TagWarn(trace *TraceContext, srtag string, m map[string]interface{}) {
	m[_srTag] = checkDLTag(srtag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	slog.Warn(parseParams(m))
}

func (l *Logger) TagError(trace *TraceContext, srtag string, m map[string]interface{}) {
	m[_srTag] = checkDLTag(srtag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	slog.Error(parseParams(m))
}

func (l *Logger) TagTrace(trace *TraceContext, srtag string, m map[string]interface{}) {
	m[_srTag] = checkDLTag(srtag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	slog.Trace(parseParams(m))
}

func (l *Logger) TagDebug(trace *TraceContext, srtag string, m map[string]interface{}) {
	m[_srTag] = checkDLTag(srtag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	slog.Debug(parseParams(m))
}

func (l *Logger) Close() {
	slog.Close()
}

// CreateBizDLTag 生成业务srtag
func CreateBizDLTag(tagName string) string {
	if tagName == "" {
		return _srTagBizUndef
	}

	return _srTagBizPrefix + tagName
}

// 校验srtag合法性
func checkDLTag(srtag string) string {
	if strings.HasPrefix(srtag, _srTagBizPrefix) {
		return srtag
	}

	if strings.HasPrefix(srtag, "_com_") {
		return srtag
	}

	if srtag == SRTagUndefind {
		return srtag
	}
	return srtag
}

// map格式化为string
func parseParams(m map[string]interface{}) string {
	var srtag string = "_undef"
	if _srtag, _have := m["srtag"]; _have {
		if __val, __ok := _srtag.(string); __ok {
			srtag = __val
		}
	}
	for _key, _val := range m {
		if _key == "srtag" {
			continue
		}
		srtag = srtag + "||" + fmt.Sprintf("%v=%+v", _key, _val)
	}
	srtag = strings.Trim(fmt.Sprintf("%q", srtag), "\"")
	return srtag
}

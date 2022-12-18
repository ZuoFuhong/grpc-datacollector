package document

import (
	"encoding/json"
	"time"
)

// TraceDoc 链路数据文档
type TraceDoc struct {
	TraceId      string `json:"trace_id"`       // 全链路ID
	Cmd          string `json:"cmd"`            // 命令字
	Project      string `json:"project"`        // 被调方
	Source       string `json:"source"`         // 调用方
	ServerIp     string `json:"server_ip"`      // 服务器IP
	Errcode      int64  `json:"errcode"`        // 错误码
	Errmsg       string `json:"errmsg"`         // 异常消息
	Timestamp    int64  `json:"timestamp"`      // 时间戳
	Timecost     int64  `json:"timecost"`       // 耗时，单位：ms
	Reqbody      string `json:"reqbody"`        // 请求参数
	SpanId       int64  `json:"span_id"`        // SpanID
	PatentSpanId int64  `json:"patent_span_id"` // 父 SpanID
	AtTimestamp  string `json:"@timestamp"`     // Kibana 时间搜索（时区：UTC）
}

func ConvertTraceDoc(reqBytes []byte) *TraceDoc {
	doc := new(TraceDoc)
	_ = json.Unmarshal(reqBytes, doc)
	doc.AtTimestamp = time.Unix(doc.Timestamp, 0).Format(time.RFC3339)
	return doc
}

func (c *TraceDoc) Bytes() []byte {
	bytes, _ := json.Marshal(c)
	return bytes
}

package es

import (
	"bytes"
	"context"
	glog "github.com/ZuoFuhong/grpc-datacollector/pkg/log"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type TraceIndex struct {
	esdb *elasticsearch.Client
}

func NewTraceIndex(esDb *elasticsearch.Client) *TraceIndex {
	return &TraceIndex{
		esdb: esDb,
	}
}

// WriteDocument 文档写入
// 这里是用于演示项目流程，逻辑比较简单。如果要用于生产，还有很多问题需要解决：
// 1.写吞吐量很低，将单个文档写入改成 Bulk 批量文档写入，增大写吞吐量。
// 2.生产环境的数据量是很大的，不能全部写入到一个 Index 中，要考虑按日期拆分 Index。
func (t *TraceIndex) WriteDocument(ctx context.Context, index string, buf []byte) error {
	req := esapi.IndexRequest{
		Index: index,
		Body:  bytes.NewBuffer(buf),
	}
	res, err := req.Do(context.Background(), t.esdb)
	if err != nil {
		glog.Errorf("Error getting response: %s", err)
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		glog.Errorf("[%s] Error indexing document", res.Status())
		return err
	}
	return nil
}

package utils

import (
	"context"

	es "github.com/olivere/elastic/v7"

	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/config"
)

const (
	MaxChunkSize   = 10000
	MidChunkSize   = 1000
	IndexChunkInfo = "chunk_info"
)

var (
	Client *es.Client
)

type ChunkList struct {
	AliasName string `json:"alias_name"`
	ChunkId   []int  `json:"chunk_id"`
}

func InitEs(conf config.ESConf) error {
	var options []es.ClientOptionFunc
	options = append(options, es.SetURL(conf.Url...))
	options = append(options, es.SetBasicAuth(conf.Username, conf.Password))
	options = append(options, es.SetSniff(true))
	// options = append(options, es.SetTraceLog(eslog.New(os.Stdout, "", 0)))  //debug

	client, err := es.NewClient(options...)
	if err != nil {
		return err
	}
	Client = client

	return nil
}

func ExistedIndex(ctx context.Context, index string) (bool, error) {
	return es.NewIndicesExistsService(Client).Index([]string{index}).Do(ctx)
}

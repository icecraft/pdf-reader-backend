package utils

import (
	"context"

	es "github.com/olivere/elastic/v7"

	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/config"
)

func InitEs(conf config.ESConf) (*es.Client, error) {
	var options []es.ClientOptionFunc
	options = append(options, es.SetURL(conf.Url...))
	options = append(options, es.SetBasicAuth(conf.Username, conf.Password))
	options = append(options, es.SetSniff(true))
	// options = append(options, es.SetTraceLog(eslog.New(os.Stdout, "", 0)))  //debug

	return es.NewClient(options...)
}

func ExistedIndex(ctx context.Context, client *es.Client, index string) (bool, error) {
	return es.NewIndicesExistsService(client).Index([]string{index}).Do(ctx)
}

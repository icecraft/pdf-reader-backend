package utils

import (
	"context"

	es "github.com/olivere/elastic/v7"

	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/config"
)

func InitEs(conf config.ESConf, mapping, indexName string) (*es.Client, error) {
	var options []es.ClientOptionFunc
	options = append(options, es.SetURL(conf.Url...))
	options = append(options, es.SetBasicAuth(conf.Username, conf.Password))
	options = append(options, es.SetSniff(true))
	// options = append(options, es.SetTraceLog(eslog.New(os.Stdout, "", 0)))  //debug

	client, err := es.NewClient(options...)
	if err != nil {
		return nil, err
	}
	if len(mapping) != 0 && len(indexName) > 0 {
		exists, err := client.IndexExists(indexName).Do(context.TODO())

		if err != nil {
			return nil, err
		}

		if !exists {
			_, err = client.CreateIndex(indexName).Body(mapping).Do(context.TODO())
			if err != nil {
				return nil, err
			}
		}
	}

	return client, nil
}

func ExistedIndex(ctx context.Context, client *es.Client, index string) (bool, error) {
	return es.NewIndicesExistsService(client).Index([]string{index}).Do(ctx)
}

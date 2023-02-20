package svc

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	es "github.com/olivere/elastic/v7"

	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/types"
)

type CibaEsIndex struct {
	indexName string
	cli       *es.Client
}

func NewCibaEsIndex(indexName string, cli *es.Client) *CibaEsIndex {
	return &CibaEsIndex{indexName: indexName, cli: cli}
}

func (o *CibaEsIndex) RetrieveWordTrans(word string) (*types.RetrieveItem, string, error) {
	termQuery := es.NewTermQuery("word", strings.ToLower(word))
	sr, err := o.cli.Search().
		Index(o.indexName).
		Query(termQuery).
		Do(context.Background())

	if err != nil {

		return nil, "", err
	}
	var ret types.RetrieveItem
	if len(sr.Hits.Hits) > 0 {
		if err := json.Unmarshal(sr.Hits.Hits[0].Source, &ret); err != nil {
			return nil, "", err
		}
		return &ret, sr.Hits.Hits[0].Id, nil
	} else {
		return nil, "", nil
	}
}

func (o *CibaEsIndex) Put(record *types.RetrieveItem) error {
	if record == nil {
		return nil
	}
	record.Hits = make([]time.Time, 1)
	record.Hits[0] = time.Now().UTC()

	_, err := o.cli.Index().
		Index(o.indexName).
		BodyJson(record).
		Do(context.Background())
	return err
}

func (o *CibaEsIndex) IncrHit(id string) error {
	ts := time.Now().UTC()

	script := es.NewScript("ctx._source.hits.add(params.ts)").Param("ts", ts)
	_, err := o.cli.Update().Index(o.indexName).Id(id).
		Script(script).
		Do(context.TODO())
	return err
}

/*

type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}


script := elastic.NewScript("ctx._source.retweets += params.num").Param("num", 1)
	update, err := client.Update().Index("twitter").Id("1").
		Script(script).
		Upsert(map[string]interface{}{"retweets": 0}).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}


*/

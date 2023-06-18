package postgres

import (
	"github.com/olivere/elastic/v7"
)

type ElasticHandler interface {
}

type elasticRepository struct {
	es *elastic.Client
}

func NewElasticRepository(es *elastic.Client) ElasticHandler {
	return &elasticRepository{
		es: es,
	}
}

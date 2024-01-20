package repositories

import "github.com/hank8451/cockroach_ca/cockroach/entities"

type CockroachRepository interface {
	InsertCockroachData(in *entities.InsertCockroachDto) error
}
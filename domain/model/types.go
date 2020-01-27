package model

import (
	"errors"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type URI string

func MarshalDateTime(t time.Time) graphql.Marshaler {
	return graphql.MarshalTime(t)
}

func UnmarshalDateTime(v interface{}) (time.Time, error) {
	if s, ok := v.(string); ok {
		return time.Parse(time.RFC3339, s)
	}
	return time.Time{}, errors.New("time should be format RFC3339")
}

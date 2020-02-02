package model

import (
	"fmt"
	"time"

	"regexp"

	"github.com/99designs/gqlgen/graphql"
)

type URI string

func MarshalDateTime(t time.Time) graphql.Marshaler {
	return graphql.MarshalTime(t)
}

var datetimeFormatOnlyDate = regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2}$`)

func UnmarshalDateTime(v interface{}) (time.Time, error) {
	s, ok := v.(string)
	if !ok {
		return time.Time{}, fmt.Errorf("could not parse to datetime(%#v) to string", v)
	}
	switch {
	case datetimeFormatOnlyDate.MatchString(s):
		return time.Parse("2006-01-02", s)
	default:
		return time.Parse(time.RFC3339, s)
	}
}

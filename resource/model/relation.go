package model

import (
	"fmt"

	promModel "github.com/prometheus/common/model"
)

type RelationName struct {
	Name, Source, Target TypeName
}

type RelationChangedRecord struct {
	Since promModel.Time // The time of relation created.
	EndUp promModel.Time // The time of relation finished.
}

// Relation is the data represent dependency of resources
type Relation struct {
	RelationName          `json:",inline" yaml:",inline"`
	SourceUrn, TargetUrn  string
	RelationChangedRecord `json:",inline" yaml:",inline"`
}

func (r Relation) String() string {
	return fmt.Sprintf("%s{%s->%s}", r.Name, r.Source, r.Target)
}

func (r Relation) TypeName() TypeName {
	return TypeName(r.String())
}

type HistoricalRelation struct {
	RelationName         `json:",inline" yaml:",inline"`
	SourceUrn, TargetUrn string
	Records              []*RelationChangedRecord
}

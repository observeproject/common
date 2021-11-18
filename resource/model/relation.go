package model

import (
	"fmt"

	promModel "github.com/prometheus/common/model"
)

type RelationName struct {
	Name           promModel.LabelName
	Source, Target SchemaName
}

func (r RelationName) String() string {
	return fmt.Sprintf("%s_%s_%s", r.Source, r.Name, r.Target)
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

type HistoricalRelation struct {
	RelationName         `json:",inline" yaml:",inline"`
	SourceUrn, TargetUrn string
	Records              []*RelationChangedRecord
}

package model

import (
	promModel "github.com/prometheus/common/model"
)

type RelationName struct {
	Name           promModel.LabelName
	Source, Target SchemaName
}

type RelationChangedRecord struct {
	Since promModel.Time // The time of relation created.
	Endup promModel.Time // The time of relation finished.
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

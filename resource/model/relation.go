package model

import (
	promModel "github.com/prometheus/common/model"
)

type RelationMatcher struct {
	Name          promModel.LabelName
	SourceMatcher ResourceMatcher
	TargetMatcher ResourceMatcher
}

type RelationName struct {
	Name           promModel.LabelName
	Source, Target SchemaName
}

// RelationQuery is the instant query Params
type RelationQuery struct {
	Time      promModel.Time
	Selectors []*RelationMatcher
}

type RelationQueryResponse struct {
	Content map[RelationName][]Relation
}

type RelationQueryRange struct {
	Start     promModel.Time
	End       promModel.Time
	Selectors []*RelationMatcher
}

type RelationQueryRangeResponse struct {
	Content map[SchemaName][]HistoricalRelation
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

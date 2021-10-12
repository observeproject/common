package model

import (
	promModel "github.com/prometheus/common/model"
)

type ResourceSchema struct {
	Name       SchemaName
	Brief      string
	Prefix     string
	UrnPattern string // UrnPattern is a NamedRegex that can be used transform urn from attributes.
	Attributes []*AttributeSchema
}

type AttributeSchema struct {
	Name     promModel.LabelName
	Brief    string
	Type     string
	Required bool
}

type RelationSchema struct {
	Name           promModel.LabelName
	Brief          string
	Source, Target ResourceSchema
	Conditions     []*RelationConditionSchema
}

type RelationConditionSchema struct {
	SourceAttribute promModel.LabelName
	TargetAttribute promModel.LabelName
}

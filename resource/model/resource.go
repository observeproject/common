package model

import (
	promModel "github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/pkg/labels"
)

// Request & Response Section Begin

// ResourceQuery is the instant query Params
type ResourceQuery struct {
	Time      promModel.Time
	Selectors []*ResourceSelector
}

type ResourceSelector struct {
	// SourceMatcher Start point of resource searching, required.
	SourceMatcher ResourceMatcher `json:",omitempty" yaml:",omitempty"`
	// TargetMatcher Endpoint of resource searching, required.
	TargetMatcher ResourceMatcher `json:",omitempty" yaml:",omitempty"`
	// ShowState return the resource with global state.
	ShowState bool
	// Relations return more resources that match the relation of Matcher, optional.
	Relations []*RelationName `json:",omitempty" yaml:",omitempty"`
}

type ResourceMatcher struct {
	Type     SchemaName
	Matchers []*labels.Matcher `json:",omitempty" yaml:",omitempty"`
}

type ResourceQueryResponse struct {
	Resources map[SchemaName][]Resource `json:",omitempty" yaml:",omitempty"`
	Relations map[SchemaName][]Relation `json:",omitempty" yaml:",omitempty"`
}

type ResourceQueryRange struct {
	Start     promModel.Time
	End       promModel.Time
	Selectors []*ResourceSelector
}

type ResourceQueryRangeResponse struct {
	Resources map[SchemaName][]HistoricalResource `json:",omitempty" yaml:",omitempty"`
	Relations map[SchemaName][]HistoricalRelation `json:",omitempty" yaml:",omitempty"`
}

// Request & Response Section End

// Resource Model Section Begin

// Resource is an immutable representation of the entity producing telemetry.
type Resource struct {
	Type           SchemaName   // Related with resource's type, and the type connected with a specification.
	Urn            string       // The Unique resource name of this resources, must be unique with whole scope.
	SecondaryTypes []SchemaName // Additional type of resource, used for observability.
	Attributes     []*Attribute // Attributes of the resource, may be a required or optional.
	States         *State       // State of the resource, name should be unique.
}

type HistoricalResource struct {
	Type           SchemaName   // Related with resource's type, and the type connected with a specification.
	Urn            string       // The Unique resource name of this resources, must be unique with whole scope.
	SecondaryTypes []SchemaName // Additional type of resource, used for observability.

	Attributes []*HistoricalAttribute // Attributes of the resource, may be a required or optional.
	State      *HistoricalState       // State of the resource, name should be unique.
}

// Attribute is a Key-value struct for describe a property of the resource, and can be used for resources selection.
type Attribute struct {
	// The name of an attribute, must be unique for the resource level and cannot be null or empty.
	Name         promModel.LabelName
	StringRecord `yaml:",inline" json:",inline"`
}

type HistoricalAttribute struct {
	Name    promModel.LabelName
	Records []*StringRecord
}

type StringRecord struct {
	Since promModel.Time       // The time of the value change to.
	EndUp promModel.Time       // The time of relation finished.
	Value promModel.LabelValue // The value, cannot be null or empty.
}

// Resource Model Section End

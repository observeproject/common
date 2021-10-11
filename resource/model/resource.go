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
	ResourceMatcher
	ShowStates []*promModel.LabelName
}

type ResourceMatcher struct {
	Type     SchemaName
	Matchers []*labels.Matcher
}

type ResourceQueryResponse struct {
	Content map[SchemaName][]Resource
}

type ResourceQueryRange struct {
	Start     promModel.Time
	End       promModel.Time
	Selectors []*ResourceSelector
}

type ResourceQueryRangeResponse struct {
	Content map[SchemaName][]HistoricalResource
}

// Request & Response Section End

// Resource Model Section Begin

// Resource is an immutable representation of the entity producing telemetry.
type Resource struct {
	Type           SchemaName   // Related with resource's type, and the type connected with a specification.
	Urn            string       // The Unique resource name of this resources, must be unique with whole scope.
	SecondaryTypes []SchemaName // Additional type of resource, used for observability.
	Attributes     []*Attribute // Attributes of the resource, may be a required or optional.
	States         []*State     // State of the resource, name should be unique.
}

type HistoricalResource struct {
	Type           SchemaName   // Related with resource's type, and the type connected with a specification.
	Urn            string       // The Unique resource name of this resources, must be unique with whole scope.
	SecondaryTypes []SchemaName // Additional type of resource, used for observability.

	Attributes []*HistoricalAttribute // Attributes of the resource, may be a required or optional.
	States     []*HistoricalState     // State of the resource, name should be unique.
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
	Value promModel.LabelValue // The value, cannot be null or empty.
}

type StateValue int

const (
	NORMAL StateValue = iota << 2
	INFO
	WARN
	ERROR
	CRITICAL
)

// State is a Key-value struct for describe a status of the resource in some aspect. The value is enum value.
type State struct {
	Name        promModel.LabelName // The name of the state, must be unique in the resource level and cannot be null or empty.
	StateRecord `json:",inline" yaml:",inline"`
}

type StateRecord struct {
	Since promModel.Time // The time of state value change to.
	Value StateValue     // The value of the state, must be match StateValue.
}

type HistoricalState struct {
	Name    promModel.LabelName
	Records []*StateRecord
}

// Resource Model Section End

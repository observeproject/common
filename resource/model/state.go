package model

import (
	promModel "github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/pkg/labels"
)

type StateValue int

const (
	NORMAL StateValue = iota << 2
	INFO
	WARN
	ERROR
	CRITICAL
)

// Request & Response Section Begin

// StateQuery is the instant query Params
type StateQuery struct {
	Time      promModel.Time
	Selectors []*StateSelector
}

type StateSelector struct {
	// Matcher Start point of resource searching, required.
	Matcher ResourceMatcher

	// States selector of stateName
	States labels.Matcher
}

type StateQueryResponse struct {
	Content map[string][]*State
}

type StateQueryRange struct {
	Start   promModel.Time
	End     promModel.Time
	Matcher ResourceMatcher

	// States selector of stateName
	States labels.Matcher
}

type StateQueryRangeResponse struct {
	Content map[string][]*HistoricalState
}
// Request & Response Section End

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

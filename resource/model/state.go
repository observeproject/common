package model

import (
	promModel "github.com/prometheus/common/model"
)

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
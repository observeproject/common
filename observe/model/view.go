package model

import (
	resModel "github.com/observeproject/common/resource/model"
	promModel "github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/pkg/labels"
)

// View Service view is made up with a set of resources, and is used for all observe operations.
type View struct {
	Name      string
	Alias     string
	Scene     string
	Resources map[resModel.SchemaName]*resModel.ResourceMatcher
	Relations []*resModel.RelationName
	Styles    ViewStyles
}

type ViewStyles struct {
	PlainView PlainViewStyle
	TreeView  TreeViewStyle
	StageView StageViewStyle
}

type ViewStyleInf interface {
	Enabled() bool
}
type ViewStyle struct {
	Enabled bool
}

type ResourceMatcherTemplateInf interface {
	Rendering(res resModel.Resource) *resModel.ResourceMatcher
}
type ResourceMatcherTemplate struct {
	ParamType resModel.SchemaName
	Type      resModel.SchemaName
	Matchers  []*MatcherTemplate
}

func (t *ResourceMatcherTemplate) Rendering(res resModel.Resource) *resModel.ResourceMatcher {
	matchers := make([]*labels.Matcher, 0)
	// TODO:

	return &resModel.ResourceMatcher{
		Type:     t.Type,
		Matchers: matchers,
	}
}

type MatcherTemplate struct {
	ParamType resModel.SchemaName
	Type      labels.MatchType
	Name      promModel.LabelName
	Template  string
}

type TreeViewItem struct {
	InputParam         resModel.SchemaName      // Parent node Resource Type
	Relation           resModel.RelationName    // Relation between current type and parent type, used to determain parent node.
	AdditionalMatchers resModel.ResourceMatcher // Additional Matchers used for filter
}

type TreeViewStyle struct {
	ViewStyle
	Items []*TreeViewItem
}

type ViewStage struct {
	Name promModel.LabelName

	InputParam resModel.SchemaName // Parent node Resource Type
	Resources  map[resModel.SchemaName]*resModel.ResourceMatcher
	Relations  []*resModel.RelationName
}

type StageViewStyle struct {
	ViewStyle
	Stages []*ViewStage
}

type PlainViewStyle struct {
	ViewStyle
}

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
	Resources []*resModel.ResourceMatcher `json:",omitempty"`
	Relations []*resModel.RelationName    `json:",omitempty"`
	Styles    *ViewStyles                 `json:",omitempty"`
}

// ViewStyles defines the different show style for view
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

// type TreeViewItem struct {
// 	SourceType resModel.SchemaName       // SourceType is the SchemaName of parent node, optional.
// 	Current    ViewSliceResourceSelector // Current resource matcher wrapper for view use
// }

// type TreeViewStyle struct {
// 	ViewStyle
// 	Items []*TreeViewItem
// }

type ViewStage struct {
	StageName promModel.LabelName // StageName is unique in the view.

	SourceType resModel.TypeName            // FromType SchemaName which is entrance of current stage, required.
	Current    []*ViewSliceResourceSelector // Current resource matcher wrapper for view use.
	Relations  []*resModel.RelationName     // Relations Searching with relations from current resources.
}

// ViewSliceResourceSelector used for resource query or relation query purpose.
type ViewSliceResourceSelector struct {
	TargetType   resModel.TypeName       // TargetType is the SchemaName of current Node(s), required
	RelationName []resModel.RelationName // RelationName combines with Type and additional Params type to be a relation.
	Matchers     []*labels.Matcher       // Matchers is the resource selector for filtering.
}

type StageViewStyle struct {
	ViewStyle
	Stages []*ViewStage
}

type PlainViewStyle struct {
	ViewStyle
}

type TreeViewStyle struct {
	ViewStyle
	Items []*TreeViewItem
}

type TreeViewItem struct {
	Matcher   resModel.ResourceMatcher
	Relations []*resModel.RelationName
	Children  []*TreeViewItem
}

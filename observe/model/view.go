package model

import resModel "github.com/observeproject/common/resource/model"

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
	Context() ViewContext
}

type ViewContext interface {
}

type ViewStyle struct {
	Enabled bool
}

type TreeViewItem struct {
	// TODO：
	Input   resModel.SchemaName
	Content resModel.RelationName
}

type TreeViewStyle struct {
	ViewStyle
	Items []*TreeViewItem
}

type ViewStage struct {
	// TODO：
}

type StageViewStyle struct {
	ViewStyle
	Stages []*ViewStage
}

type PlainViewStyle struct {
	ViewStyle
}

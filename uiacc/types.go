//agov:shared
package uiacc

import "fmt"

// Node 控件树里的一个节点(边界为屏幕像素)。Dump() 返回的原始扁平结构。
type Node struct {
	Text           string
	Desc           string
	ID             string
	Class          string
	Pkg            string
	Clickable      bool
	LongClickable  bool
	Checkable      bool
	Checked        bool
	Scrollable     bool
	Editable       bool
	Enabled        bool
	Selected       bool
	Focusable      bool
	Focused        bool
	Password       bool
	Index          int
	Parent         int // 父节点在 Dump() 结果里的下标(-1=根)
	Self           int // 本节点在 Dump() 结果里的下标
	X1, Y1, X2, Y2 int
}

// CenterX 控件中心的 X 像素坐标。
func (n Node) CenterX() int { return (n.X1 + n.X2) / 2 }

// CenterY 控件中心的 Y 像素坐标。
func (n Node) CenterY() int { return (n.Y1 + n.Y2) / 2 }

// Width 控件像素宽。
func (n Node) Width() int { return n.X2 - n.X1 }

// Height 控件像素高。
func (n Node) Height() int { return n.Y2 - n.Y1 }

// IsRoot 是否为根节点(无父节点)。
func (n Node) IsRoot() bool { return n.Parent < 0 }

// ParentOf 在一次 Dump 的数组里取 n 的父节点;无父节点(根)返回 (Node{}, false)。
func ParentOf(nodes []Node, n Node) (Node, bool) {
	if n.Parent < 0 || n.Parent >= len(nodes) {
		return Node{}, false
	}
	return nodes[n.Parent], true
}

// ChildrenOf 在一次 Dump 的数组里取 n 的全部直接子节点。
func ChildrenOf(nodes []Node, n Node) []Node {
	var out []Node
	for _, c := range nodes {
		if c.Parent == n.Self {
			out = append(out, c)
		}
	}
	return out
}

// Rect 控件包围盒(对标 AutoGo uiacc.Rect)。
type Rect struct {
	Left, Top, Right, Bottom int
	CenterX, CenterY         int
	Width, Height            int
}

// Uiacc 控件选择器:链式叠加条件,全部满足才命中(对标 AutoGo *Uiacc)。用 New() 开始。
//
//	if o := uiacc.New().Text("登录").WaitFor(5000); o != nil { o.Click() }
//	uiacc.New().ClassName("android.widget.EditText").FindOnce().SetText("你好")
type Uiacc struct {
	text, textContains, textStartsWith, textEndsWith, textMatches string
	desc, descContains, descEndsWith, descMatches                 string
	id, idContains, idEndsWith, idMatches                         string
	class, classEndsWith, classMatches                            string
	pkg                                                           string
	clickable, longClickable, checkable, checked, scrollable      *bool
	editable, enabled, selected, focusable, focused, password     *bool
	index                                                         *int
	bounds, boundsInside, boundsContains                          []int
}

// New 新建选择器(对标 AutoGo uiacc.New)。
func New() *Uiacc { return &Uiacc{} }

func bp(v bool) *bool { return &v }

// 以下为链式条件(全部满足才命中)。字符串条件:同一字段可叠加(如 Text+TextContains 同时生效);
// 空字符串表示不限。Matches 系列走 Go 正则(regexp.MatchString)。

// Text 限定 text 精确等于 v。
func (a *Uiacc) Text(v string) *Uiacc { a.text = v; return a }

// TextContains 限定 text 包含子串 v。
func (a *Uiacc) TextContains(v string) *Uiacc { a.textContains = v; return a }

// TextStartsWith 限定 text 以 v 开头。
func (a *Uiacc) TextStartsWith(v string) *Uiacc { a.textStartsWith = v; return a }

// TextEndsWith 限定 text 以 v 结尾。
func (a *Uiacc) TextEndsWith(v string) *Uiacc { a.textEndsWith = v; return a }

// TextMatches 限定 text 匹配正则 v。
func (a *Uiacc) TextMatches(v string) *Uiacc { a.textMatches = v; return a }

// Desc 限定 contentDescription 精确等于 v。
func (a *Uiacc) Desc(v string) *Uiacc { a.desc = v; return a }

// DescContains 限定 contentDescription 包含子串 v。
func (a *Uiacc) DescContains(v string) *Uiacc { a.descContains = v; return a }

// DescEndsWith 限定 contentDescription 以 v 结尾。
func (a *Uiacc) DescEndsWith(v string) *Uiacc { a.descEndsWith = v; return a }

// DescMatches 限定 contentDescription 匹配正则 v。
func (a *Uiacc) DescMatches(v string) *Uiacc { a.descMatches = v; return a }

// Id 限定资源 id(如 com.pkg:id/name)精确等于 v。
func (a *Uiacc) Id(v string) *Uiacc { a.id = v; return a }

// IdContains 限定资源 id 包含子串 v。
func (a *Uiacc) IdContains(v string) *Uiacc { a.idContains = v; return a }

// IdEndsWith 限定资源 id 以 v 结尾。
func (a *Uiacc) IdEndsWith(v string) *Uiacc { a.idEndsWith = v; return a }

// IdMatches 限定资源 id 匹配正则 v。
func (a *Uiacc) IdMatches(v string) *Uiacc { a.idMatches = v; return a }

// ClassName 限定类名(如 android.widget.EditText)精确等于 v。
func (a *Uiacc) ClassName(v string) *Uiacc { a.class = v; return a }

// ClassNameEndsWith 限定类名以 v 结尾。
func (a *Uiacc) ClassNameEndsWith(v string) *Uiacc { a.classEndsWith = v; return a }

// ClassNameMatches 限定类名匹配正则 v。
func (a *Uiacc) ClassNameMatches(v string) *Uiacc { a.classMatches = v; return a }

// PackageName 限定所属应用包名精确等于 v。
func (a *Uiacc) PackageName(v string) *Uiacc { a.pkg = v; return a }

// Clickable 限定 clickable 属性为 v。
func (a *Uiacc) Clickable(v bool) *Uiacc { a.clickable = bp(v); return a }

// LongClickable 限定 longClickable 属性为 v。
func (a *Uiacc) LongClickable(v bool) *Uiacc { a.longClickable = bp(v); return a }

// Checkable 限定 checkable 属性为 v。
func (a *Uiacc) Checkable(v bool) *Uiacc { a.checkable = bp(v); return a }

// Checked 限定 checked 属性为 v。
func (a *Uiacc) Checked(v bool) *Uiacc { a.checked = bp(v); return a }

// Scrollable 限定 scrollable 属性为 v。
func (a *Uiacc) Scrollable(v bool) *Uiacc { a.scrollable = bp(v); return a }

// Editable 限定 editable 属性为 v。
func (a *Uiacc) Editable(v bool) *Uiacc { a.editable = bp(v); return a }

// Enabled 限定 enabled 属性为 v。
func (a *Uiacc) Enabled(v bool) *Uiacc { a.enabled = bp(v); return a }

// Selected 限定 selected 属性为 v。
func (a *Uiacc) Selected(v bool) *Uiacc { a.selected = bp(v); return a }

// Focusable 限定 focusable 属性为 v。
func (a *Uiacc) Focusable(v bool) *Uiacc { a.focusable = bp(v); return a }

// Focused 限定 focused 属性为 v。
func (a *Uiacc) Focused(v bool) *Uiacc { a.focused = bp(v); return a }

// Password 限定 password 属性为 v。
func (a *Uiacc) Password(v bool) *Uiacc { a.password = bp(v); return a }

// Index 限定节点在父容器中的 index 等于 v。
func (a *Uiacc) Index(v int) *Uiacc { a.index = &v; return a }

// Bounds 限定控件包围盒精确等于 [l,t,r,b](屏幕像素)。
func (a *Uiacc) Bounds(l, t, r, b int) *Uiacc { a.bounds = []int{l, t, r, b}; return a }

// BoundsInside 限定控件包围盒完全落在矩形 [l,t,r,b] 内部。
func (a *Uiacc) BoundsInside(l, t, r, b int) *Uiacc { a.boundsInside = []int{l, t, r, b}; return a }

// BoundsContains 限定控件包围盒完全包住矩形 [l,t,r,b]。
func (a *Uiacc) BoundsContains(l, t, r, b int) *Uiacc { a.boundsContains = []int{l, t, r, b}; return a }

// UiObject 一个命中的控件句柄(对标 AutoGo *UiObject)。持有节点快照与同批控件树,支持读属性、导航、动作。
type UiObject struct {
	n   Node
	all []Node
}

// 以下 Get* 读取命中时的节点快照属性(对标 AutoGo Get*)。

// GetText 控件文本。
func (o *UiObject) GetText() string { return o.n.Text }

// GetDesc 控件 contentDescription。
func (o *UiObject) GetDesc() string { return o.n.Desc }

// GetId 控件资源 id。
func (o *UiObject) GetId() string { return o.n.ID }

// GetClassName 控件类名。
func (o *UiObject) GetClassName() string { return o.n.Class }

// GetPackageName 控件所属应用包名。
func (o *UiObject) GetPackageName() string { return o.n.Pkg }

// GetClickable 控件是否可点击。
func (o *UiObject) GetClickable() bool { return o.n.Clickable }

// GetChecked 控件是否已勾选。
func (o *UiObject) GetChecked() bool { return o.n.Checked }

// GetScrollable 控件是否可滚动。
func (o *UiObject) GetScrollable() bool { return o.n.Scrollable }

// GetEditable 控件是否可编辑。
func (o *UiObject) GetEditable() bool { return o.n.Editable }

// GetSelected 控件是否被选中。
func (o *UiObject) GetSelected() bool { return o.n.Selected }

// GetIndex 控件在父容器中的 index。
func (o *UiObject) GetIndex() int { return o.n.Index }

// CenterX 控件中心的 X 像素坐标。
func (o *UiObject) CenterX() int { return o.n.CenterX() }

// CenterY 控件中心的 Y 像素坐标。
func (o *UiObject) CenterY() int { return o.n.CenterY() }

// GetBounds 控件包围盒。
func (o *UiObject) GetBounds() Rect {
	return Rect{Left: o.n.X1, Top: o.n.Y1, Right: o.n.X2, Bottom: o.n.Y2,
		CenterX: o.n.CenterX(), CenterY: o.n.CenterY(), Width: o.n.Width(), Height: o.n.Height()}
}

// GetChildCount 直接子节点数。
func (o *UiObject) GetChildCount() int { return len(ChildrenOf(o.all, o.n)) }

// GetParent 父控件句柄(无则 nil)。
func (o *UiObject) GetParent() *UiObject {
	if p, ok := ParentOf(o.all, o.n); ok {
		return &UiObject{n: p, all: o.all}
	}
	return nil
}

// GetChildren 直接子控件句柄。
func (o *UiObject) GetChildren() []*UiObject {
	cs := ChildrenOf(o.all, o.n)
	out := make([]*UiObject, len(cs))
	for i, c := range cs {
		out[i] = &UiObject{n: c, all: o.all}
	}
	return out
}

// ToString 便于调试。
func (o *UiObject) ToString() string {
	return fmt.Sprintf("UiObject{text=%q id=%q class=%q bounds=[%d,%d,%d,%d]}",
		o.n.Text, o.n.ID, o.n.Class, o.n.X1, o.n.Y1, o.n.X2, o.n.Y2)
}

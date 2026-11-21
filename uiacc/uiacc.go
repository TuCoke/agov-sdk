// Package uiacc 控件树(UiAutomation,不需要开无障碍)。API 形态对标 AutoGo:
//
//	if o := uiacc.New().TextContains("蓝牙").WaitFor(5000); o != nil { o.Click() }
//	uiacc.New().ClassName("android.widget.EditText").FindOnce().SetText("你好")
//	for _, o := range uiacc.New().Clickable(true).Find() { … }
//
// 实现是【胶水】:控件树、注入、performAction 的真身在 utils.dex(二进制);
// Go 发桥命令(单条件直接 performAction)或取树后本地匹配再点中心。
package uiacc

// Dump 取当前整棵控件树(原始扁平数组)。失败返回 nil。
func Dump() []Node { var r0 []Node; return r0 }

// "ok <json>"

// FindOnce 取树后返回第一个命中控件句柄;无则 nil(对标 AutoGo FindOnce)。
func (a *Uiacc) FindOnce() *UiObject { var r0 *UiObject; return r0 }

// Find 取树后返回全部命中控件句柄(对标 AutoGo Find)。
func (a *Uiacc) Find() []*UiObject { var r0 []*UiObject; return r0 }

// WaitFor 反复取树直到命中或超时(timeoutMs),返回句柄;超时 nil(对标 AutoGo WaitFor)。
func (a *Uiacc) WaitFor(timeoutMs int) *UiObject { var r0 *UiObject; return r0 }

// Exists 当前界面是否存在命中控件。
func (a *Uiacc) Exists() bool { var r0 bool; return r0 }

// Click 找到并点击第一个命中控件(便捷;等价 FindOnce().Click())。
func (a *Uiacc) Click() bool { var r0 bool; return r0 }

// bridgeBy 选择器能否用桥单条件表达(恰一个简单条件)→ 走 dex performAction。

// ---- UiObject 动作(优先 dex performAction,退到坐标注入)----

// Click 点击控件(优先 performAction,退到点中心)。
func (o *UiObject) Click() bool { var r0 bool; return r0 }

// LongClick 长按控件。
func (o *UiObject) LongClick() bool { var r0 bool; return r0 }

// SetText 往输入框写文字。
func (o *UiObject) SetText(text string) bool { var r0 bool; return r0 }

// ScrollForward 对控件向前(向下/向右)滚一页。
func (o *UiObject) ScrollForward() bool { var r0 bool; return r0 }

// ScrollBackward 对控件向后(向上/向左)滚一页。
func (o *UiObject) ScrollBackward() bool { var r0 bool; return r0 }

// Collapse 折叠控件(如可展开的列表项)。
func (o *UiObject) Collapse() bool { var r0 bool; return r0 }

// Expand 展开控件(如可展开的列表项)。
func (o *UiObject) Expand() bool { var r0 bool; return r0 }

// Copy 对当前焦点输入框执行复制。
func Copy() bool { var r0 bool; return r0 }

// Cut 对当前焦点输入框执行剪切。
func Cut() bool { var r0 bool; return r0 }

// Paste 对当前焦点输入框执行粘贴。
func Paste() bool { var r0 bool; return r0 }

// SetSelection 设置当前焦点输入框选区 [start, end)。
func SetSelection(start, end int) bool { var r0 bool; return r0 }

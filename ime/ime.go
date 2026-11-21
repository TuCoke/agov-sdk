// Package ime 剪贴板、文本输入与输入法(胶水:框架桥 + shell)。
//
// 能力总览:
//   - 剪贴板:GetClipText / SetClipText 读写系统剪贴板。
//   - 文本输入:InputText 往焦点输入框提交文本(走 commitText,能过联想/搜索框)、
//     PasteText 置剪贴板再粘贴。
//   - 输入法:GetCurrentIME / SetCurrentIME 查/切当前输入法、GetIMEList 列已装输入法。
//     惯例是「用时切、用完切回」,不设默认,免挡人工打字。
package ime

// GetClipText 读系统剪贴板文本。
func GetClipText() string { var r0 string; return r0 }

// "ok <b64>"

// SetClipText 写系统剪贴板。
func SetClipText(text string) bool { var r0 bool; return r0 }

// InputText 往当前焦点输入框输入文字(中文可用)。
func InputText(text string) bool { var r0 bool; return r0 }

// PasteText 写剪贴板再粘贴到焦点输入框。
func PasteText(text string) bool { var r0 bool; return r0 }

// GetIMEList 已安装输入法 ID 列表。
func GetIMEList() []string { var r0 []string; return r0 }

// GetCurrentIME 当前输入法 ID。
func GetCurrentIME() string { var r0 string; return r0 }

// SetCurrentIME 切换输入法(需已启用)。
func SetCurrentIME(imeID string) error { var r0 error; return r0 }

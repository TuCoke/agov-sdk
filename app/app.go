// Package app 应用控制(胶水:monkey/am/pm/cmd/dumpsys)。
//
// 能力总览:
//   - 启动/前台:Launch 按包名启动、CurrentPackage / CurrentActivity 查前台、GetBrowserPackage 默认浏览器。
//   - 安装管理:Install / Uninstall、Enable / Disable、Clear 清数据、ForceStop 强停。
//   - 查询:Installed 是否已装、Version 版本名、Name 显示名、List 列应用(第三方/含系统、可过滤)。
//   - 打开:OpenUrl 打开网址、OpenSetting 应用详情页、ViewFile 用 VIEW Intent 打开本地文件。
//   - Intent:StartActivity / SendBroadcast / StartService 按 IntentOptions 发任意 Intent。
//   - 其他:IgnoreBattOpt 加电池优化白名单(免 Doze)。
package app

// Launch 按包名启动应用(桌面入口)。
func Launch(pkg string) bool { var r0 bool; return r0 }

// CurrentPackage 当前前台应用包名。
func CurrentPackage() string { var r0 string; return r0 }

// CurrentActivity 当前前台 "pkg/Activity"。
func CurrentActivity() string { var r0 string; return r0 }

// Installed 是否已安装。
func Installed(pkg string) bool { var r0 bool; return r0 }

// Name 应用显示名(需 aapt,常取不到时返回包名)。
func Name(pkg string) string { var r0 string; return r0 }

// Version 应用版本名。
func Version(pkg string) string { var r0 string; return r0 }

// List 列应用:includeSystem=false 只列第三方;filter 为子串过滤(可空)。
func List(includeSystem bool, filter string) []AppInfo { var r0 []AppInfo; return r0 }

// Install 安装 APK(-r 覆盖安装)。
func Install(path string) error { var r0 error; return r0 }

// Uninstall 卸载应用。
func Uninstall(pkg string) error { var r0 error; return r0 }

// Clear 清除应用数据(pm clear)。
func Clear(pkg string) error { var r0 error; return r0 }

// ForceStop 强制停止应用(am force-stop)。
func ForceStop(pkg string) error { var r0 error; return r0 }

// Enable 启用应用组件(pm enable)。
func Enable(pkg string) error { var r0 error; return r0 }

// Disable 对当前用户停用应用(pm disable-user)。
func Disable(pkg string) error { var r0 error; return r0 }

// IgnoreBattOpt 把应用加入电池优化白名单(免 Doze 限制)。
func IgnoreBattOpt(pkg string) error { var r0 error; return r0 }

// OpenSetting 打开该应用的「应用详情」系统设置页。
func OpenSetting(pkg string) error { var r0 error; return r0 }

// OpenUrl 用系统默认应用打开一个 URL(VIEW Intent)。
func OpenUrl(url string) error { var r0 error; return r0 }

// ViewFile 用 VIEW Intent 打开本地文件;mime 可空(空则由系统推断)。
func ViewFile(path, mime string) error { var r0 error; return r0 }

// GetBrowserPackage 默认浏览器包名。
func GetBrowserPackage() string { var r0 string; return r0 }

// StartActivity 按 IntentOptions 发 Intent 启动 Activity(am start)。
func StartActivity(o IntentOptions) error { var r0 error; return r0 }

// SendBroadcast 按 IntentOptions 发送广播(am broadcast)。
func SendBroadcast(o IntentOptions) error { var r0 error; return r0 }

// StartService 按 IntentOptions 启动 Service(am startservice)。
func StartService(o IntentOptions) error { var r0 error; return r0 }

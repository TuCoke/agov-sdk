//agov:shared
package app

// AppInfo 一条应用信息。
type AppInfo struct {
	PackageName string
	Label       string
	VersionName string
}

// IntentOptions 发 Intent 的参数(StartActivity / SendBroadcast / StartService)。
type IntentOptions struct {
	Action      string            // 如 android.intent.action.VIEW
	Data        string            // 数据 URI,如 https://… / package: / tel:
	Type        string            // MIME 类型
	PackageName string            // 目标包名
	ClassName   string            // 目标组件全类名(可空)
	Categories  []string          // 类别
	Extras      map[string]string // 字符串 extra
	Flags       []string          // 原样附加的 am 选项,如 "--activity-clear-top"
}

//agov:shared
package motion

// Android KeyEvent 键码(常用子集)。
const (
	KeyHome         = 3
	KeyBack         = 4
	KeyCall         = 5
	KeyEndCall      = 6
	KeyVolumeUp     = 24
	KeyVolumeDown   = 25
	KeyPower        = 26
	KeyCamera       = 27
	KeyTab          = 61
	KeySpace        = 62
	KeyEnter        = 66
	KeyDel          = 67
	KeyMenu         = 82
	KeyNotification = 83
	KeySearch       = 84
	KeyEscape       = 111
	KeyMoveHome     = 122
	KeyMoveEnd      = 123
	KeyAppSwitch    = 187
	KeyPaste        = 279
)

// Back 返回键。
func Back() { Key(KeyBack) }

// Home 主页键。
func Home() { Key(KeyHome) }

// Recents 最近任务。
func Recents() { Key(KeyAppSwitch) }

// Menu 菜单键。
func Menu() { Key(KeyMenu) }

// Enter 回车。
func Enter() { Key(KeyEnter) }

// VolumeUp 音量加。
func VolumeUp() { Key(KeyVolumeUp) }

// VolumeDown 音量减。
func VolumeDown() { Key(KeyVolumeDown) }

// Power 电源键(亮/灭屏)。
func Power() { Key(KeyPower) }

// Notifications 下拉通知栏。
func Notifications() { Key(KeyNotification) }

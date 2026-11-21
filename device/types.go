//agov:shared
package device

// Notification 一条通知栏消息。
type Notification struct {
	Pkg   string
	Title string
	Text  string
}

// 音量流(GetVolume / SetVolume 的 stream 参数)。
const (
	StreamVoice        = 0
	StreamSystem       = 1
	StreamRing         = 2
	StreamMusic        = 3
	StreamAlarm        = 4
	StreamNotification = 5
)

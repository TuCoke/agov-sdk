// Package device 读写设备信息与状态(胶水:shell + 框架桥)。
//
// 能力总览:
//   - 屏幕/显示:Size / CurrentSize 分辨率、Density DPI、Rotation 朝向、
//     IsScreenOn / ScreenUnlocked 亮屏与锁屏、Wake / Sleep 常亮开关、SetDisplayPower 亮灭屏。
//   - 亮度:GetBrightness / SetBrightness(0..255)、GetBrightnessMode / SetBrightnessMode 自动手动。
//   - 音量/震动:GetVolume / SetVolume 按音量流、Vibrate / CancelVibration 震动(走框架桥,破部分 ROM 的音量拦截)。
//   - 设备标识:Model 型号、Serial 序列号、AndroidVersion 系统版本、GetAndroidId、GetWifiMac、Imei(A10+ 非 root 取不到)。
//   - 电量/内存:Battery 电量、BatteryStatus 充电状态、MemInfo 总量/可用。
//   - 网络:GetIp 局域网 IP、GetWifiMac。
//   - 前台/通知:Foreground 当前前台窗口、Notifications 读通知栏。
//   - 电源:Reboot 重启(需 root)。
package device

// firstIntAfter 在 out 里找含 key 的行,取该行第一个整数。

// Size 屏幕物理分辨率(宽, 高),不随旋转变化。
func Size() (int, int) { var r0 int; var r1 int; return r0, r1 }

// CurrentSize 当前朝向下的分辨率(横屏交换)。
func CurrentSize() (int, int) { var r0 int; var r1 int; return r0, r1 }

// Density 屏幕 DPI。
func Density() int { var r0 int; return r0 }

// Rotation 屏幕旋转 0/1/2/3。
func Rotation() int { var r0 int; return r0 }

// Model 设备型号(ro.product.model)。
func Model() string { var r0 string; return r0 }

// Serial 设备序列号(ro.serialno)。
func Serial() string { var r0 string; return r0 }

// AndroidVersion 系统版本号(ro.build.version.release,如 "14")。
func AndroidVersion() string { var r0 string; return r0 }

// Battery 电量百分比。
func Battery() int { var r0 int; return r0 }

// BatteryStatus 电池状态:1 未知 2 充电 3 放电 4 不充 5 满。
func BatteryStatus() int { var r0 int; return r0 }

// IsScreenOn 屏幕是否亮着。
func IsScreenOn() bool { var r0 bool; return r0 }

// ScreenUnlocked 是否已解锁(粗判)。
func ScreenUnlocked() bool { var r0 bool; return r0 }

// Wake 亮屏并保持常亮。
func Wake() {}

// Sleep 取消常亮。
func Sleep() {}

// SetDisplayPower 亮/灭屏。
func SetDisplayPower(on bool) {}

// GetBrightness 当前亮度 0..255。
func GetBrightness() int { var r0 int; return r0 }

// GetBrightnessMode 亮度模式:0 手动 1 自动。
func GetBrightnessMode() int { var r0 int; return r0 }

// SetBrightness 设亮度 0..255(切手动)。
func SetBrightness(v int) {}

// SetBrightnessMode 设自动/手动。
func SetBrightnessMode(auto bool) {}

// GetVolume 某音量流的 (当前, 最大);走框架桥。
func GetVolume(stream int) (int, int) { var r0 int; var r1 int; return r0, r1 }

// "ok cur max"

// SetVolume 设某音量流。
func SetVolume(stream, v int) {}

// Vibrate 震动 ms 毫秒(走框架桥)。
func Vibrate(ms int) error { var r0 error; return r0 }

// CancelVibration 停止震动。
func CancelVibration() error { var r0 error; return r0 }

// GetIp 局域网 IP。
func GetIp() string { var r0 string; return r0 }

// GetWifiMac WiFi MAC。
func GetWifiMac() string { var r0 string; return r0 }

// GetAndroidId ANDROID_ID。
func GetAndroidId() string { var r0 string; return r0 }

// Imei IMEI(Android 10+ 非 root 取不到,返回空)。
func Imei() string { var r0 string; return r0 }

// MemInfo 内存 (总量 KB, 可用 KB)。
func MemInfo() (int, int) { var r0 int; var r1 int; return r0, r1 }

// Foreground 当前前台窗口(pkg/Activity)。
func Foreground() string { var r0 string; return r0 }

// Reboot 重启设备(需 root)。
func Reboot() error { var r0 error; return r0 }

// Notifications 读当前通知栏(best-effort)。
func Notifications() []Notification { var r0 []Notification; return r0 }

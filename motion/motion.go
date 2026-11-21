// Package motion 触控与按键。坐标为屏幕像素。
// 实现是【胶水】:命令发给 utils.dex 框架桥(injectInputEvent,低延迟、多指、跨 ROM);
// 桥不可用时回退 input 命令。拟人滑动曲线在 dex(二进制)里做。
package motion

// Init 探测桥是否就绪(不调也会在首次操作时自动连接/拉起)。返回是否拿到低延迟注入(桥)。
func Init() bool { var r0 bool; return r0 }

// HasInjector 当前是否有低延迟注入(桥可用)。
func HasInjector() bool { var r0 bool; return r0 }

// Tap 点击像素坐标 (x, y)。
func Tap(x, y int) {}

// LongPress 长按;durMs<=0 取 600ms。
func LongPress(x, y, durMs int) {}

// Swipe 从 (x1,y1) 滑到 (x2,y2),历时 durMs 毫秒(<=0 取默认);拟人轨迹在 dex 内生成。
func Swipe(x1, y1, x2, y2, durMs int) {}

// TouchDown 手动分解触摸:手指(id 0)在 (x,y) 按下;桥不可用则无效。
func TouchDown(x, y int) {}

// TouchMove 手动分解触摸:手指(id 0)移动到 (x,y);需先 TouchDown;桥不可用则无效。
func TouchMove(x, y int) {}

// TouchUp 手动分解触摸:抬起手指(id 0);桥不可用则无效。
func TouchUp() {}

// Pinch 双指缩放:两指沿水平轴从 中心±fromR 到 中心±toR(像素半径),分 steps 步。仅桥可用。
func Pinch(cx, cy, fromR, toR, durMs int) bool { var r0 bool; return r0 }

// Key 发送按键(Android KeyEvent 键码,见 Key* 常量)。
func Key(code int) {}

// Text 往当前焦点输入框输入文字(中文可用)。须先点中输入框;失败返回 false。
func Text(s string) bool { var r0 bool; return r0 }

// Paste 写系统剪贴板再粘贴(桥不可用时输中文的兜底)。
func Paste(s string) bool { var r0 bool; return r0 }

// Package sdk 是 autogov 安卓自动化 SDK 的根:操作手机的全部能力按包拆分,用户 import 需要的包即可。
//
//	motion   触控与按键(像素坐标)
//	images   截图 / 找色 / 找图 / 图像处理(*image.NRGBA)
//	ocr      文字识别(PP-OCR)
//	yolo     目标检测(YOLOv8 / ncnn)
//	uiacc    控件树:选择器链 + 节点动作(免无障碍,走 UiAutomation)
//	device   设备信息与状态
//	app      应用控制
//	files    文件系统
//	storages 本地 KV
//	https    网络请求
//	system   进程与 shell
//	ime      剪贴板 / 文本输入 / 输入法切换
//	media    媒体库
//	utils    纯 Go 小工具(Sleep/Random/Log/类型转换)
//
// 架构:这一层(sdk 真身)是普通 Go,和用户脚本一起编成【一个二进制、一个 Go 运行时】,
// 直接调用底层能力包;不涉及 .so 边界或 JSON。真身在编译期由 agov build 换入,
// 用户项目里放的是签名一致的纯 Go 空壳(stub),桌面 go build / IDE 补全对着空壳。
//
// 约定:
//   - 坐标一律屏幕像素(对标 AutoGo);区域 (x1,y1,x2,y2) 中 x2/y2<=0 表示到屏幕边缘。
//   - 颜色:"RRGGBB",可写 "RRGGBB-DRDGDB" 偏色、"A|B" 多选;多点找色见 images 文档。
//   - 相似度 sim 与阈值 thresh:0..1,越大越严格。
package sdk

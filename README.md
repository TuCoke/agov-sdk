# agov-sdk

Android 自动化 SDK 的 **公开 API 层**（Go）。按包拆分，import 需要的包即可。

> ⚠️ **这是 API 壳（stub）**：签名与文档完整，**函数体是空的**。
> 它让你在桌面 `go build` 通过、让 IDE（gopls）补全和跳转可用，
> 但**不能直接在手机上运行** —— 真身由 `agov` 构建器在编译时换入。

## 安装

```bash
go get github.com/TuCoke/agov-sdk@latest
```

国内建议：

```bash
go env -w GOPROXY=https://goproxy.cn,direct
```

## 快速开始

```go
package main

import (
    "github.com/TuCoke/agov-sdk/images"
    "github.com/TuCoke/agov-sdk/motion"
    "github.com/TuCoke/agov-sdk/utils"
)

func main() {
    img := images.Capture(0, 0, 0, 0) // 全屏截图，返回 *image.NRGBA
    if img == nil {
        return
    }
    utils.Log("screen: %v", img.Bounds())

    x, y := images.FindColor(0, 0, 0, 0, "FF0000", 0.9) // 找不到返回 (-1,-1)
    if x < 0 {
        return
    }
    motion.Tap(x, y)
    utils.Sleep(500)
}
```

桌面阶段：

```bash
go build ./...
```

能编过、能补全，但跑不到真机 —— 这是预期行为。

运行阶段（需要 `agov` 构建器）：

```bash
agov build      # 换入真身 + NDK 交叉编译
agov run -d <serial>   # 部署到设备并运行
```

## 包

| 包 | 用途 |
|---|---|
| `motion` | 点击 / 滑动 / 长按 / 多指 / 手势帧 / 按键 |
| `images` | 截图 / 找图 / 找色 / 图像处理（`*image.NRGBA`） |
| `ocr` | 文字识别（PP-OCR） |
| `yolo` | 目标检测（YOLOv8 / ncnn） |
| `uiacc` | 控件树：选择器链 + 节点动作（免无障碍，走 UiAutomation） |
| `device` | 设备信息与状态（尺寸 / 电量 / 亮度 / 旋转 / 网络） |
| `app` | 应用控制（启动 / 停止 / 安装 / 列表 / 当前前台） |
| `files` | 文件系统 |
| `storages` | 本地键值存储 |
| `https` | HTTP 客户端（Get / Post / 下载） |
| `system` | 进程与 shell |
| `ime` | 文本输入 / 输入法 / 剪贴板 |
| `media` | 媒体库 |
| `utils` | 纯 Go 小工具（Sleep / Random / Log / 类型转换） |

## 约定

- **坐标**：一律屏幕像素。区域 `(x1,y1,x2,y2)` 中 `x2/y2 <= 0` 表示到屏幕边缘。
- **颜色**：`"RRGGBB"`，可写 `"RRGGBB-DRDGDB"` 表示偏色、`"A|B"` 表示多候选。
- **相似度 / 阈值**：`0..1`，越大越严格。
- 找图/找色未命中返回 `(-1,-1)`，不是错误。

## 平台

| 平台 | 状态 |
|---|---|
| Android (arm64-v8a / x86_64 / x86) | ✅ 可用 |
| iOS | 🚧 规划中（当前 API 为 Android 形状） |

## 示例

见 [`examples/`](examples/)：

- `screenshot-tap` —— 截图 → 找色 → 点击
- `ocr-find-text` —— OCR 找字 → 点击中心
- `ui-tree` —— 控件树里找可点击节点并点击

这些示例只验证"能编译"，真机效果需要构建器。

## License

[Apache-2.0](LICENSE)

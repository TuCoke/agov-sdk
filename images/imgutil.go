//agov:shared

package images

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// ---- 读 / 写 ----

// Decode 解码 PNG/JPEG 字节为图像对象;失败返回 nil。
func Decode(data []byte) *image.NRGBA {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil
	}
	return ToNRGBA(img)
}

// ReadFile 读图片文件;失败返回 nil。
func ReadFile(path string) *image.NRGBA {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	return Decode(b)
}

// ReadBase64 解码 base64 图片。
func ReadBase64(s string) *image.NRGBA {
	b, err := base64.StdEncoding.DecodeString(strings.TrimSpace(s))
	if err != nil {
		return nil
	}
	return Decode(b)
}

// Encode 编码为 "png" 或 "jpg"(quality 1..100,仅 jpg 用)。
func Encode(img *image.NRGBA, format string, quality int) []byte {
	if img == nil {
		return nil
	}
	var buf bytes.Buffer
	switch strings.ToLower(strings.TrimPrefix(format, ".")) {
	case "jpg", "jpeg":
		if quality <= 0 || quality > 100 {
			quality = 90
		}
		_ = jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
	default:
		_ = png.Encode(&buf, img)
	}
	return buf.Bytes()
}

// EncodeBase64 编码为 base64 字符串。
func EncodeBase64(img *image.NRGBA, format string, quality int) string {
	return base64.StdEncoding.EncodeToString(Encode(img, format, quality))
}

// Save 按扩展名(.png/.jpg)保存到文件;父目录不存在会创建。
func Save(img *image.NRGBA, path string, quality int) error {
	if img == nil {
		return errNilImage
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, Encode(img, filepath.Ext(path), quality), 0o644)
}

// ToNRGBA 任意 image.Image 转 *image.NRGBA(已经是则原样返回)。
func ToNRGBA(img image.Image) *image.NRGBA {
	if img == nil {
		return nil
	}
	if n, ok := img.(*image.NRGBA); ok {
		return n
	}
	b := img.Bounds()
	out := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(out, out.Bounds(), img, b.Min, draw.Src)
	return out
}

// ---- 处理 ----

// Clip 裁剪 (x1,y1,x2,y2)(x2/y2<=0 到边),返回新图。
func Clip(img *image.NRGBA, x1, y1, x2, y2 int) *image.NRGBA {
	if img == nil {
		return nil
	}
	w, h := img.Rect.Dx(), img.Rect.Dy()
	if x2 <= 0 || x2 > w {
		x2 = w
	}
	if y2 <= 0 || y2 > h {
		y2 = h
	}
	if x1 < 0 {
		x1 = 0
	}
	if y1 < 0 {
		y1 = 0
	}
	if x1 >= x2 || y1 >= y2 {
		return nil
	}
	out := image.NewNRGBA(image.Rect(0, 0, x2-x1, y2-y1))
	draw.Draw(out, out.Bounds(), img, image.Pt(img.Rect.Min.X+x1, img.Rect.Min.Y+y1), draw.Src)
	return out
}

// Resize 缩放到 width×height(双线性)。
func Resize(img *image.NRGBA, width, height int) *image.NRGBA {
	if img == nil || width <= 0 || height <= 0 {
		return nil
	}
	sw, sh := img.Rect.Dx(), img.Rect.Dy()
	out := image.NewNRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		fy := (float64(y)+0.5)*float64(sh)/float64(height) - 0.5
		y0 := int(fy)
		if y0 < 0 {
			y0 = 0
		}
		y1 := y0 + 1
		if y1 >= sh {
			y1 = sh - 1
		}
		wy := fy - float64(y0)
		if wy < 0 {
			wy = 0
		}
		for x := 0; x < width; x++ {
			fx := (float64(x)+0.5)*float64(sw)/float64(width) - 0.5
			x0 := int(fx)
			if x0 < 0 {
				x0 = 0
			}
			x1 := x0 + 1
			if x1 >= sw {
				x1 = sw - 1
			}
			wx := fx - float64(x0)
			if wx < 0 {
				wx = 0
			}
			p00 := img.PixOffset(img.Rect.Min.X+x0, img.Rect.Min.Y+y0)
			p10 := img.PixOffset(img.Rect.Min.X+x1, img.Rect.Min.Y+y0)
			p01 := img.PixOffset(img.Rect.Min.X+x0, img.Rect.Min.Y+y1)
			p11 := img.PixOffset(img.Rect.Min.X+x1, img.Rect.Min.Y+y1)
			o := out.PixOffset(x, y)
			for c := 0; c < 4; c++ {
				top := float64(img.Pix[p00+c])*(1-wx) + float64(img.Pix[p10+c])*wx
				bot := float64(img.Pix[p01+c])*(1-wx) + float64(img.Pix[p11+c])*wx
				out.Pix[o+c] = uint8(top*(1-wy) + bot*wy + 0.5)
			}
		}
	}
	return out
}

// Rotate 顺时针旋转 degree 度(取 90 的整数倍;其他值四舍五入到最近的 90 倍)。
func Rotate(img *image.NRGBA, degree int) *image.NRGBA {
	if img == nil {
		return nil
	}
	q := ((degree/90)%4 + 4) % 4
	if degree%90 >= 45 {
		q = (q + 1) % 4
	}
	w, h := img.Rect.Dx(), img.Rect.Dy()
	var out *image.NRGBA
	switch q {
	case 0:
		out = image.NewNRGBA(image.Rect(0, 0, w, h))
		draw.Draw(out, out.Bounds(), img, img.Rect.Min, draw.Src)
		return out
	case 2:
		out = image.NewNRGBA(image.Rect(0, 0, w, h))
	default:
		out = image.NewNRGBA(image.Rect(0, 0, h, w))
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			var nx, ny int
			switch q {
			case 1:
				nx, ny = h-1-y, x
			case 2:
				nx, ny = w-1-x, h-1-y
			case 3:
				nx, ny = y, w-1-x
			}
			copy(out.Pix[out.PixOffset(nx, ny):out.PixOffset(nx, ny)+4],
				img.Pix[img.PixOffset(img.Rect.Min.X+x, img.Rect.Min.Y+y):img.PixOffset(img.Rect.Min.X+x, img.Rect.Min.Y+y)+4])
		}
	}
	return out
}

// Grayscale 转灰度图。
func Grayscale(img *image.NRGBA) *image.Gray {
	if img == nil {
		return nil
	}
	w, h := img.Rect.Dx(), img.Rect.Dy()
	out := image.NewGray(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p := img.PixOffset(img.Rect.Min.X+x, img.Rect.Min.Y+y)
			r, g, b := int(img.Pix[p]), int(img.Pix[p+1]), int(img.Pix[p+2])
			out.Pix[out.PixOffset(x, y)] = uint8((r*299 + g*587 + b*114) / 1000)
		}
	}
	return out
}

// Threshold 二值化:>thresh 的像素置 maxVal,其余 0;inverse=true 反转。
func Threshold(gray *image.Gray, thresh, maxVal int, inverse bool) *image.Gray {
	if gray == nil {
		return nil
	}
	out := image.NewGray(gray.Rect)
	for i, v := range gray.Pix {
		hit := int(v) > thresh
		if inverse {
			hit = !hit
		}
		if hit {
			out.Pix[i] = uint8(maxVal)
		}
	}
	return out
}

// Binarize 二值化(>thresh → 255)。
func Binarize(gray *image.Gray, thresh int) *image.Gray { return Threshold(gray, thresh, 255, false) }

// GrayToNRGBA 灰度图转回彩色对象(给 FindImageIn / ocr.Image 等只收 NRGBA 的函数)。
func GrayToNRGBA(gray *image.Gray) *image.NRGBA {
	if gray == nil {
		return nil
	}
	out := image.NewNRGBA(gray.Rect)
	draw.Draw(out, out.Bounds(), gray, gray.Rect.Min, draw.Src)
	return out
}

// ColorOf 把 "RRGGBB" 解析成 color.NRGBA;非法返回全 0。
func ColorOf(hex string) color.NRGBA {
	hex = strings.TrimPrefix(strings.TrimSpace(hex), "#")
	if len(hex) != 6 {
		return color.NRGBA{}
	}
	v := func(s string) uint8 {
		var n uint8
		for _, c := range s {
			n <<= 4
			switch {
			case c >= '0' && c <= '9':
				n |= uint8(c - '0')
			case c >= 'a' && c <= 'f':
				n |= uint8(c-'a') + 10
			case c >= 'A' && c <= 'F':
				n |= uint8(c-'A') + 10
			}
		}
		return n
	}
	return color.NRGBA{R: v(hex[0:2]), G: v(hex[2:4]), B: v(hex[4:6]), A: 0xff}
}

type imgErr string

func (e imgErr) Error() string { return string(e) }

const errNilImage = imgErr("images: 图像为 nil")

// Package https 网络请求(纯 Go 标准库)。timeoutSec<=0 取默认 30 秒。
//
// 能力总览:Get / Post 基本请求、PostMultipart 表单上传文件、Download 下载到设备文件。
// 均返回状态码,便于判成败。
package https

// Get 发 GET,返回 (状态码, 响应体, 错误)。
func Get(url string, timeoutSec int) (int, []byte, error) {
	var r0 int
	var r1 []byte
	var r2 error
	return r0, r1, r2
}

// Post 发 POST(自定义 body 与头),headers 可空。
func Post(url string, data []byte, headers map[string]string, timeoutSec int) (int, []byte, error) {
	var r0 int
	var r1 []byte
	var r2 error
	return r0, r1, r2
}

// PostMultipart 以 multipart/form-data 上传一个文件(字段名 file)。
func PostMultipart(url, fileName string, fileData []byte, timeoutSec int) (int, []byte, error) {
	var r0 int
	var r1 []byte
	var r2 error
	return r0, r1, r2
}

// Download 下载到设备路径,返回 (状态码, 字节数, 错误)。
func Download(url, path string, timeoutSec int) (int, int64, error) {
	var r0 int
	var r1 int64
	var r2 error
	return r0, r1, r2
}

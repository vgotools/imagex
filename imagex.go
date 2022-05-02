package imagex

import (
	"strings"
)

// 各服务商图片处理标识
const (
	ProcessIdentificationAliCloud  = "?x-oss-process=image/info"
	ProcessIdentificationQiNiu     = "?imageMogr2"
	ProcessIdentificationByteDance = "~tplv-"
)

const (
	BasicInfoSuffixAliCloud  = "?x-oss-process=image/info"
	BasicInfoSuffixQiNiu     = "?imageInfo"
	BasicInfoSuffixByteDance = "~info"
)

// ImageInfo 公共图片信息（七牛云/字节可通用）
// Example: 七牛 http://dn-odum9helk.qbox.me/resource/gogopher.jpg?imageInfo
// Example: 字节 http://p3-imagex.byteimg.com/imagex-rc/preview.jpg~info
type ImageInfo struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Format string `json:"format"`
	Size   int    `json:"size"`
}

// AliOssImageInfo 阿里云图片基本信息
// Example: http://image-demo.oss-cn-hangzhou.aliyuncs.com/example.jpg?x-oss-process=image/info
type AliOssImageInfo struct {
	FileSize struct {
		Value string `json:"value"`
	} `json:"FileSize"`
	Format struct {
		Value string `json:"value"`
	} `json:"Format"`
	ImageHeight struct {
		Value string `json:"value"`
	} `json:"ImageHeight"`
	ImageWidth struct {
		Value string `json:"value"`
	} `json:"ImageWidth"`
}

// GetOssOriginalUrl 获取 OSS 图片未做处理链接
func GetOssOriginalUrl(imgUrl string) string {
	// 去除服务商图片处理参数
	for _, idf := range getProcessIdentifications() {
		if strings.Contains(imgUrl, idf) {
			s := strings.Split(imgUrl, idf)
			if len(s) > 1 {
				return s[0]
			}
		}
	}
	return imgUrl
}

// GetOssOriginalUrlWithIdf 通过 ProcessIdentification 获取 OSS 图片未做处理链接
func GetOssOriginalUrlWithIdf(imgUrl, idf string) string {
	// 去除服务商图片处理参数
	if strings.Contains(imgUrl, idf) {
		s := strings.Split(imgUrl, idf)
		if len(s) > 1 {
			return s[0]
		}
	}
	return imgUrl
}

func getProcessIdentifications() []string {
	return []string{
		ProcessIdentificationAliCloud,
		ProcessIdentificationQiNiu,
		ProcessIdentificationByteDance,
	}
}

// IsWebpUrl 是否是 webp url
func IsWebpUrl(imgUrl string) bool {
	if strings.Contains(imgUrl, ".webp") {
		ss := strings.Split(imgUrl, ".")
		if len(ss) > 0 && ss[len(ss)-1] == "webp" {
			return true
		}
	}
	return false
}

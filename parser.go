package imagex

import (
	"encoding/json"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"strconv"

	"golang.org/x/image/webp"
)

// ParseImage 常规解析图片信息
// 支持 gif/jpeg/png 暂不支持 webp
func ParseImage(imgUrl string) (image.Image, string, error) {
	resp, err := http.Get(imgUrl)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	var (
		img    image.Image
		format string
	)
	if IsWebpUrl(imgUrl) {
		img, err = webp.Decode(resp.Body)
		format = "webp"
	} else {
		img, format, err = image.Decode(resp.Body)
	}

	return img, format, err
}

// ParseAliOssImageInfo 解析阿里云图片基本信息
// imgUrl 需为去除 OSS 图片处理参数后的链接，可参考 GetOssOriginalUrl 去除
// ApiDoc: https://help.aliyun.com/document_detail/44975.html
func ParseAliOssImageInfo(imgUrl string) (*ImageInfo, error) {
	resp, err := http.Get(imgUrl + BasicInfoSuffixAliCloud)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var imageInfo *AliOssImageInfo
	err = json.Unmarshal(body, &imageInfo)
	if err != nil {
		return nil, err
	}

	w, _ := strconv.Atoi(imageInfo.ImageWidth.Value)
	h, _ := strconv.Atoi(imageInfo.ImageHeight.Value)
	s, _ := strconv.Atoi(imageInfo.FileSize.Value)

	return &ImageInfo{
		Width:  w,
		Height: h,
		Format: imageInfo.Format.Value,
		Size:   s,
	}, nil
}

// ParseQiNiuImageInfo 解析七牛云图片基本信息
// imgUrl 需为去除 OSS 图片处理参数后的链接，可参考 GetOssOriginalUrl 去除
// ApiDoc: 七牛云 https://developer.qiniu.com/dora/1269/pictures-basic-information-imageinfo
func ParseQiNiuImageInfo(imgUrl string) (*ImageInfo, error) {
	return ParseOssImageInfo(imgUrl + BasicInfoSuffixQiNiu)
}

// ParseByteDanceImageInfo 解析七牛云/字节图片基本信息
// imgUrl 需为去除 OSS 图片处理参数后的链接，可参考 GetOssOriginalUrl 去除
// ApiDoc: 字节 https://www.volcengine.com/docs/508/64085
func ParseByteDanceImageInfo(imgUrl string) (*ImageInfo, error) {
	return ParseOssImageInfo(imgUrl + BasicInfoSuffixByteDance)
}

// ParseOssImageInfo 解析OSS服务商图片信息
// PS: 适用无图片处理后缀场景
func ParseOssImageInfo(imgUrl string) (*ImageInfo, error) {
	resp, err := http.Get(imgUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var imageInfo *ImageInfo
	err = json.Unmarshal(body, &imageInfo)
	if err != nil {
		return nil, err
	}

	return imageInfo, nil
}

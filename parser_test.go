package imagex

import (
	"strings"
	"testing"
)

func TestParseAliOssImage(t *testing.T) {
	urls := []string{
		"http://image-demo.oss-cn-hangzhou.aliyuncs.com/example.jpg",
	}

	for _, url := range urls {
		image, err := ParseAliOssImageInfo(url)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("format: %s, width: %d, height: %d", image.Format, image.Width, image.Height)
		}
	}
}

func TestParseQiNiuImageInfo(t *testing.T) {
	urls := []string{
		"http://dn-odum9helk.qbox.me/resource/gogopher.jpg",
	}

	for _, url := range urls {
		image, err := ParseQiNiuImageInfo(url)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("format: %s, width: %d, height: %d", image.Format, image.Width, image.Height)
		}
	}
}

func TestParseByteDanceImageInfo(t *testing.T) {
	urls := []string{
		"https://p3-imagex.byteimg.com/imagex-rc/preview.jpg",
	}

	for _, url := range urls {
		image, err := ParseByteDanceImageInfo(url)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("format: %s, width: %d, height: %d", image.Format, image.Width, image.Height)
		}
	}
}

func TestParseImageInfo(t *testing.T) {
	urls := []string{
		"http://image-demo.oss-cn-hangzhou.aliyuncs.com/example.jpg",
		"http://dn-odum9helk.qbox.me/resource/gogopher.jpg",
		"https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimg.jj20.com%2Fup%2Fallimg%2F1113%2F052420110515%2F200524110515-2-1200.jpg&refer=http%3A%2F%2Fimg.jj20.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1654090945&t=4ad8b48ebf362d7730b41049b3fa6d61",
	}

	for _, url := range urls {
		image, format, err := ParseImage(url)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("format: %s, width: %d, height: %d", format, image.Bounds().Dx(), image.Bounds().Dy())
		}
	}
}

func TestGetOssOriginalUrl(t *testing.T) {
	urls := []string{
		"http://image-demo.oss-cn-hangzhou.aliyuncs.com/example.jpg",
		"http://dn-odum9helk.qbox.me/resource/gogopher.jpg?imageMogr2/crop/!300x400a10a10",
		"https://p3-imagex.byteimg.com/imagex-rc/preview.jpg~tplv-image",
	}

	for _, url := range urls {
		tmp := GetOssOriginalUrl(url)
		t.Log(tmp)
	}
}

func TestImageSuffix(t *testing.T) {
	urls := []string{

		"http://image-demo.oss-cn-hangzhou.aliyuncs.com/example.jpg",
		"http://dn-odum9helk.qbox.me/resource/gogopher.jpg?imageMogr2/crop/!300x400a10a10",
		"https://p3-imagex.byteimg.com/imagex-rc/preview.jpg~tplv-image",
		"https://p3-imagex.byteimg.com/imagex-rc/preview.webp~tplv-image",
		"https://p3-imagex.byteimg.com/imagex-rc/preview.webp",
	}
	for _, s := range urls {
		if strings.Contains(s, ".webp") {
			ss := strings.Split(s, ".")
			if len(ss) > 0 && ss[len(ss)-1] == "webp" {
				t.Log(s)
			}
		}
	}
}

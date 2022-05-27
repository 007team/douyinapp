package qiniu

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"strconv"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var (
	AccessKey        = "_TBooutIx4PP6IYLADZ1ue_wBPGc46EC0yv4WlG-"
	SerectKey        = "hC5xVP__weeda_wh_1oJxlMfLlljM8Re8sP_eKyj"
	Bucket           = "007douyin"                    // bucket name
	ImgUrl           = "rchgbnnln.hn-bkt.clouddn.com" // 域名
	ErrorQiniuFailed = errors.New("七牛：视频上传失败")
)

// UploadVideoToQiNiu 将视频上传到七牛云
func UploadVideoToQiNiu(file *multipart.FileHeader, videoId int64) (int, string, error) {
	src, err := file.Open()
	if err != nil {
		log.Fatalln("qiniu put failed 1")
		return 10011, "", ErrorQiniuFailed
	}
	defer src.Close()

	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SerectKey)

	// 获取上传凭证
	upToken := putPlicy.UploadToken(mac)

	// 配置参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, // 华南区
		UseCdnDomains: false,
		UseHTTPS:      false, // 非https
	}
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}        // 上传后返回的结果
	putExtra := storage.PutExtra{} // 额外参数

	key := "videos/" + strconv.Itoa(int(videoId))

	err = formUploader.Put(context.Background(), &ret, upToken, key, src, file.Size, &putExtra)
	if err != nil {
		fmt.Println("qiniu put failed 2", err)
		code := 501
		return code, "", ErrorQiniuFailed
	}

	url := "http://" + ImgUrl + "/" + ret.Key // 返回上传后的文件访问路径
	fmt.Println("视频播放地址： ", url)
	return 0, url, nil

}

func UploadImgToQiNiu(imgName string, loadFile string, video_id int64) (url string) {
	mac := qbox.NewMac(AccessKey, SerectKey)

	// 上传策略
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}

	// 获取上传token
	upToken := putPolicy.UploadToken(mac)

	// 上传Config对象
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan //指定上传的区域
	cfg.UseHTTPS = false           // 是否使用https域名
	cfg.UseCdnDomains = false      //是否使用CDN上传加速

	// 需要上传的文件
	localFile := loadFile + "\\" + imgName

	// 七牛key
	qiniuKey := "cover/" + strconv.Itoa(int(video_id))

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 上传文件
	err := formUploader.PutFile(context.Background(), &ret, upToken, qiniuKey, localFile, nil)
	if err != nil {
		fmt.Println("上传文件失败,原因:", err)
		return
	}
	fmt.Println("上传成功,key为:", ret.Key)

	// 返回视频封面的url

	url = "http://" + ImgUrl + "/" + ret.Key
	fmt.Println(url)
	return url
}

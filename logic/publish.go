package logic

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"os"
	"strconv"
	"sync/atomic"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"

	"github.com/gin-gonic/gin"

	"github.com/007team/douyinapp/dao/mysql"
	"github.com/007team/douyinapp/dao/qiniu"
	"github.com/007team/douyinapp/models"
	"github.com/disintegration/imaging"
)

var (
	VideoPath string = "D:\\GO_WORK\\src\\douyinapp\\public\\video" // 保存视频的路径
	ImgPath   string = "D:\\GO_WORK\\src\\douyinapp\\public\\img"   // 保存图片的路径
)

func PublishList(userId int64) (VideoArr []models.Video) {
	VideoArr = mysql.GetVideoArr(userId)
	return
}

func Publish(c *gin.Context, video *models.Video, data *multipart.FileHeader) (err error) {
	atomic.AddInt64(&models.LastVideoId, 1)
	video.Id = models.LastVideoId // 新视频的videoId

	//先将视频保存到本地
	if err = c.SaveUploadedFile(data, VideoPath+"\\"+strconv.Itoa(int(video.Id))+".mp4"); err != nil {
		fmt.Println("c.SaveUploadedFile failed", err)
		return err
	}
	fmt.Println("保存视频完成")
	// 生成缩略图
	//_, err = GetSnapshot(VideoPath+`\`+strconv.Itoa(int(video.Id))+".mp4", ImgPath, 5, video.Id)
	//if err != nil {
	//	fmt.Println("缩略图生成失败", err)
	//	return err
	//}
	//fmt.Println("生成缩略图完成")

	// 上传视频
	_, fileUrl, err := qiniu.UploadVideoToQiNiu(data, video.Id)
	if err != nil {
		fmt.Println("qiniu upload video failed")
		return err
	}
	fmt.Println("上传视频完成")
	video.PlayUrl = fileUrl
	video.CoverUrl = "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"
	//上传视频封面到七牛云
	//coverUrl := qiniu.UploadImgToQiNiu("data", ImgPath, video.Id)
	//video.CoverUrl = coverUrl
	//fmt.Println("上传视频封面到七牛云完成")
	// 删除本地视频

	// 删除本地视频封面缩略图

	return mysql.CreateNewVideo(video)
}

// 截取视频第一帧 生成缩略图
func GetSnapshot(videoPath, snapshotPath string, frameNum int, video_id int64) (snapshotName string, err error) {
	fmt.Println("进 缩略图生成程序")
	buf := bytes.NewBuffer(nil)

	err = ffmpeg_go.Input(videoPath).
		Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		fmt.Println("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		fmt.Println("生成缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, snapshotPath+`\`+strconv.Itoa(int(video_id))+".jpeg")
	if err != nil {
		fmt.Println("生成缩略图失败：", err)
		return "", err
	}

	// 成功则返回生成的缩略图名
	//names := strings.Split(snapshotPath, `\`)
	snapshotName = strconv.Itoa(int(video_id)) + ".jpeg"
	fmt.Println("缩略图名是：", snapshotName)
	return
}

package logic

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/007team/douyinapp/dao/redis"
	"log"
	"math/rand"

	"github.com/007team/douyinapp/dao/mysql"

	"github.com/007team/douyinapp/models"
	"github.com/007team/douyinapp/pkg/jwt"
)

var letters = []byte("abcdefghjkmnpqrstuvwxyz123456789")

func UserInfo(user *models.User, myUserId int64) (err error) {
	// mysql查询用户具体信息
	if err = mysql.UserInfo(user); err != nil {
		log.Fatalln("mysql.UserInfo failed", err)
		return err
	}
	// redis查询用户的粉丝与关注数
	user.FollowCount, err = redis.UserFollowCount(user.Id)
	if err != nil {
		log.Println("redis.UserFollowCount(user.Id) failed", err)
		return err
	}
	user.FollowerCount, err = redis.UserFollowerCount(user.Id)
	if err != nil {
		log.Println("redis.UserFollowerCount(user.Id) failed", err)
		return err
	}
	// “我”是否关注了这个用户

	user.IsFollow, err = redis.IsFollowUser(user, myUserId)
	if err != nil {
		log.Println("redis.IsFollowUser(user, myUserId) failed", err)
		return err
	}

	//redis.IsFollowUser()
	return nil
}

func encryptPassword(oPassword string, salt string) string {
	h := md5.New()
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
func Register(user *models.User) (token string, err error) {
	//先查询该用户是否存在， 如存在则直接返回错误
	if err = mysql.CheckUserExist(user); err != nil {
		fmt.Println(" 该用户已存在 ")
		return "", err
	}

	// 对用户密码进行加密
	salt := RandLow()          //生成 salt
	oPassword := user.Password // 旧的密码
	newPassword := encryptPassword(oPassword, string(salt))
	user.Salt = string(salt) // 保存salt
	user.Password = newPassword

	// 插入新用户
	if err = mysql.CreateNewUser(user); err != nil {
		return "", err
	}

	// 生成 token
	token, _, err = jwt.GenToken(user.Id)
	if err != nil {
		log.Fatalln("jwt.GenToken  生成token失败", err)
		return "", err
	}

	return token, err
}

// RandLow 生成加密密码用的随机字符串  salt
func RandLow() []byte {
	n := 15
	if n <= 0 {
		return []byte{}
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return []byte{}
	}
	for i, x := range b {
		arc = x & 31
		b[i] = letters[arc]
	}
	return b
}

func Login(user *models.User) (err error) {
	oPassword := user.Password // 未加密的原密码

	// 查询用户
	if err = mysql.Login(user); err != nil {
		return err
	}

	// 进行密码校验
	salt := user.Salt
	newPassword := encryptPassword(oPassword, salt) // 将原密码加密
	if newPassword != user.Password {
		return mysql.ErrorInvalidUserPassword // 用户密码错误
	}

	return nil
}

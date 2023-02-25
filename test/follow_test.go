package test

import (
	"douyin/initialize"
	"douyin/internal/model"
	"douyin/pkg/utils"
	"fmt"
	"testing"
	"time"
)

func TestFollow(t *testing.T) {
	//初始化配置
	initialize.InitConfig("../config/config.toml")

	//初始化mysql
	db := initialize.InitMysql()
	num := 400000
	pwd, err := utils.EncipherPassword("hacker.123")
	if err != nil {
		fmt.Println("密码加密错误", err)
	}
	now := time.Now()
	desc := fmt.Sprintf("我于%d年%d月%d日%d时%d分%d秒注册了抖声，欢迎关注！", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	for i := 367479; i < num; i++ {
		user := model.User{
			Username: fmt.Sprintf("user%d", i),
			Password: pwd,
			//客户端用户注册暂无设置头像、背景图、简介功能，也无修改功能，所以暂时写死
			Avatar:          "https://c-ssl.duitang.com/uploads/blog/202102/08/20210208200511_45cb8.jpg",
			BackgroundImage: "https://article.autotimes.com.cn/wp-content/uploads/2022/04/95f35f8c40454bf1b4f18d7c79b5b948.jpg",
			Signature:       desc,
		}
		err = db.Model(&model.User{}).Create(&user).Error
	}
}

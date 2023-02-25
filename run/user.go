package run

import (
	"LargeDataWrite/global"
	"LargeDataWrite/model"
	"fmt"
	"log"
	"time"
)

func Register() {
	log.Println("用户注册批量模拟数据插入中......")
	num := 10000000
	now := time.Now()
	desc := fmt.Sprintf("我于%d年%d月%d日%d时%d分%d秒注册了抖声，欢迎关注！", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	var user model.User
	var err error
	for i := 0; i < num; i++ {
		user = model.User{
			Username: fmt.Sprintf("user%d", i),
			Password: "$2a$12$0qFTEYMa226BDd2UetrCA.vWGAN5h9voyZAj1J.5tfFb2ulZLQQP.",
			//客户端用户注册暂无设置头像、背景图、简介功能，也无修改功能，所以暂时写死
			Avatar:          "https://c-ssl.duitang.com/uploads/blog/202102/08/20210208200511_45cb8.jpg",
			BackgroundImage: "https://article.autotimes.com.cn/wp-content/uploads/2022/04/95f35f8c40454bf1b4f18d7c79b5b948.jpg",
			Signature:       desc,
		}
		err = global.DB.Model(&model.User{}).Create(&user).Error
		if err != nil {
			fmt.Println("创建用户失败：", err)
		}
	}
}

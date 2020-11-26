package controller

import (
	"encoding/json"
	"fmt"
	"gin_vue/common"
	"gin_vue/dto"
	"gin_vue/model"
	"gin_vue/response"
	"gin_vue/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	//使用map获取请求参数


	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	if len(name) == 0 {
		name = util.RandString(10)
	}

	//手机号已存在
	if isTelephoneExists(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
	}

	//	创建用户
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密失败")
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(fromPassword),
	}
	DB.Create(&newUser)
	//返回结果
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		log.Printf("token generate error:%v", err)
		return
	}
	response.Success(ctx, gin.H{
		"token": token,
	}, "注册成功")

}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//获取数据
	//使用map获取请求参数
	//var requestMap=make(map[string]string);
	//json.NewDecoder(ctx.Request.Body).Decode(&requestMap)

	//使用结构体来封装值
	var requestUser = model.User{}
	json.NewDecoder(ctx.Request.Body).Decode(&requestUser)


	//获取参数
	//telephone := requestUser.Telephone
	//password := requestUser.Password
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//数据验证
	fmt.Println(telephone, "手机号码长度", len(telephone))
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	//判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "用户不存在")
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		log.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": dto.ToUserDto(user.(model.User))},
	})
}

func isTelephoneExists(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"my/blogs/common"
	"my/blogs/dto"
	"my/blogs/model"
	"my/blogs/response"
	"my/blogs/util"
	"net/http"
)

//http处理函数放controller中

func Register(ctx *gin.Context) {
	DB := common.DB
	//使用map获取请求的参数
	//var requestMap = make(map[string]string)
	//json.NewDecoder(ctx.Request.Body).Decode(&requestMap)

	//使用结构体
	var requestUser = model.User{}
	//json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&requestUser)

	//获取参数
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 3 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	//如果名称没有传，给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)
	//判断手机号是否存在

	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
	}
	//创建用户
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(encryptedPassword),
	}
	DB.Create(&newUser)

	//此处为了让用户注册成功之后直接登录，不需要再输入账号密码
	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error:%v", err)
		return
	}
	//返回结果
	response.Success(ctx, gin.H{"token": token}, "注册成功")

}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//使用结构体
	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	//获取参数
	telephone := requestUser.Telephone
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 3 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error:%v", err)
		return
	}
	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"my/blogs/common"
	"my/blogs/model"
	"my/blogs/response"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})
	return CategoryController{DB: db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, "数据验证错误，分类名称必填", nil)
	}
	c.DB.Create(&requestCategory)
	response.Success(ctx, gin.H{"category": requestCategory}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	//绑定body中的参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, "数据验证错误，分类名称必填", nil)
	}
	//获取path中的参数
	//categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var updateCategory model.Category
	//if c.DB.First(&updateCategory, categoryId).{
	//	response.Fail(ctx,"分类不存在",nil)
	//}

	//更新分类
	//可传三种类型：map、struct、name value
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}

//查看

func (c CategoryController) Show(ctx *gin.Context) {
	//获取path中的参数
	//categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var category model.Category
	//if c.DB.First(&updateCategory, categoryId).{
	//	response.Fail(ctx,"分类不存在",nil)
	//}
	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
		response.Fail(ctx, "删除失败，请重试", nil)
		return
	}
	response.Success(ctx, nil, "")
}

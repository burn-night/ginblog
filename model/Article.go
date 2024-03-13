package model

import (
	"fmt"
	"ginblo/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

// CreateArticle 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

// 查询分类下的所有文章
func GetCateArticle(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var cateArtlist []Article
	var total int64
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"cid =?", id).Find(&cateArtlist).Count(&total).Error
	if err != nil {
		fmt.Println("这个错误", err)
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtlist, errmsg.SUCCESS, int(total)
}

// GetCategory 查询单个文章
func GetArticle(id int) (Article, int, int) {
	var art Article
	var total int64
	err := db.Preload("Category").Limit(1).Where("ID = ?", id).Find(&art).Count(&total).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST, 0
	}
	return art, errmsg.SUCCESS, int(total)
}

// GetArticles 查询文章列表
func GetArticles(pageSize int, pageNum int) ([]Article, int) {
	var art []Article
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&art).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return art, errmsg.SUCCESS
}

// EditCategory 编辑文章信息
func EditArticle(id int, data *Article) int {
	var cate Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func DeleteArticle(id int) int {
	var art Article
	err = db.Where("id = ? ", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

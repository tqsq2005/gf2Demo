package model

import "gf2-demo/internal/model/entity"

// RotationCreateUpdateBase 创建/修改内容基类
type RotationCreateUpdateBase struct {
	PicUrl string // 图片链接
	Link   string // 跳转链接
	Sort   int    // 排序
}

// RotationCreateInput 创建内容
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 创建内容返回结果
type RotationCreateOutput struct {
	RotationId uint `json:"rotation_id"`
}

// RotationUpdateInput 修改内容
type RotationUpdateInput struct {
	RotationCreateUpdateBase
	Id uint
}

// RotationListItem 主要用于详情或列表展示
type RotationListItem struct {
	Id     uint   `json:"id"`      // 自增ID
	PicUrl string `json:"pic_url"` // 图片链接
	Link   string `json:"link"`    // 跳转链接
	Sort   int    `json:"sort"`    // 排序
}

// RotationGetDetailInput 详情输入
type RotationGetDetailInput struct {
	Id uint `json:"id"`
}

// RotationGetDetailOutput 详情输出
type RotationGetDetailOutput struct {
	Rotation *entity.RotationInfo `json:"rotation"`
}

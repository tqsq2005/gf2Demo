package backend

import (
	"gf2-demo/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type RotationCreateReq struct {
	g.Meta `path:"/backend/rotation/add" tags:"Rotation" method:"post" summary:"You first rotation api"`
	PicUrl string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort"    dc:"图片排序"`
}
type RotationCreateRes struct {
	// g.Meta `mime:"text/html" example:"string"`
	RotationId uint `json:"rotationId"`
}

type RotationShowReq struct {
	g.Meta `path:"/backend/rotation/show" tags:"Rotation" method:"post" summary:"get rotation detail by id"`
	Id     uint `json:"id" v:"min:1#请输入Id值" dc:"轮播图ID"`
}
type RotationShowRes struct {
	Rotation *entity.RotationInfo
}

type RotationUpdateReq struct {
	g.Meta `path:"/backend/rotation/update/{Id}" method:"post" tags:"Rotation" summary:"修改轮播图接口"`
	Id     uint   `json:"id" v:"min:1#请输入Id值" dc:"轮播图ID"`
	PicUrl string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort"    dc:"图片排序"`
}
type RotationUpdateRes struct{}

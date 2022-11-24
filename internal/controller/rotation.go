package controller

import (
	"context"
	"errors"

	"gf2-demo/api/backend"
	"gf2-demo/internal/model"
	"gf2-demo/internal/service"
)

// Rotation 内容管理
var Rotation = cRotation{}

type cRotation struct{}

func (a *cRotation) Create(ctx context.Context, req *backend.RotationCreateReq) (res *backend.RotationCreateRes, err error) {
	out, err := service.Rotation().Create(ctx, model.RotationCreateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RotationCreateRes{RotationId: out.RotationId}, nil
}

// Show 获取详情
//
//	@receiver a
//	@param ctx
//	@param req
//	@return res
//	@return err
func (a *cRotation) Show(ctx context.Context, req *backend.RotationShowReq) (res *backend.RotationShowRes, err error) {
	out, err := service.Rotation().GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return &backend.RotationShowRes{Rotation: nil}, errors.New("传入的ID值找不到指定记录，请检查参数！")
	}
	return &backend.RotationShowRes{Rotation: out.Rotation}, nil
}

// Update 更新
//
//	@receiver a
//	@param ctx
//	@param req
//	@return res
//	@return err
func (a *cRotation) Update(ctx context.Context, req *backend.RotationUpdateReq) (res *backend.RotationUpdateRes, err error) {
	err = service.Rotation().Update(ctx, model.RotationUpdateInput{
		Id: req.Id,
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	return
}

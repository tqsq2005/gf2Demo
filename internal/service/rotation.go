// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf2-demo/internal/model"
)

type (
	IRotation interface {
		Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error)
		GetDetail(ctx context.Context, id uint) (out *model.RotationGetDetailOutput, err error)
		Update(ctx context.Context, in model.RotationUpdateInput) error
	}
)

var (
	localRotation IRotation
)

func Rotation() IRotation {
	if localRotation == nil {
		panic("implement not found for interface IRotation, forgot register?")
	}
	return localRotation
}

func RegisterRotation(i IRotation) {
	localRotation = i
}

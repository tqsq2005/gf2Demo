package rotation

import (
	"context"
	"gf2-demo/internal/dao"
	"gf2-demo/internal/model"
	"gf2-demo/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
)

// generate: gf gen service
type sRotation struct{}

func init() {
	service.RegisterRotation(New())
}

func New() *sRotation {
	return &sRotation{}
}

// Create 创建内容
func (s *sRotation) Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.RotationInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RotationCreateOutput{RotationId: uint(lastInsertID)}, err
}

// GetDetail 查询详情
func (s *sRotation) GetDetail(ctx context.Context, id uint) (out *model.RotationGetDetailOutput, err error) {
	out = &model.RotationGetDetailOutput{}
	if err := dao.RotationInfo.Ctx(ctx).WherePri(id).Scan(&out.Rotation); err != nil {
		return nil, err
	}
	// 没有数据
	if out.Rotation == nil {
		return nil, nil
	}
	return out, nil
}

// Update 修改
func (s *sRotation) Update(ctx context.Context, in model.RotationUpdateInput) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.RotationInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.RotationInfo.Columns().Id).
			Where(dao.RotationInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

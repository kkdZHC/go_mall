package dao

import (
	"context"
	"go_mall/model"

	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// 通过uid找notice
func (dao *NoticeDao) GetNoticeById(uId uint) (*model.Notice, error) {
	var Notice *model.Notice
	err := dao.DB.Model(&model.Notice{}).Where("id=?", uId).First(&Notice).Error
	return Notice, err
}

package models

import "time"

type Blog struct {
	Base         `xorm:"-" json:"-" form:"-"`
	Id           int64  `xorm:"not null pk autoincr comment('主键ID') INT" json:"id,omitempty" form:"id" `
	Uid          int64  `xorm:"not null default 0 comment('发布人ID') INT" json:"uid,omitempty" form:"uid" binding:"required"`
	Title        string `xorm:"not null default '' comment('文章标题') VARCHAR(255)" json:"title,omitempty" form:"title" binding:"required,min=6,max=100"`
	Content      string `xorm:"not null comment('文章内容') TEXT" json:"content,omitempty" form:"content" binding:"required,min=100"`
	CreateTime   int64  `xorm:"not null default 0 comment('创建时间') INT" json:"create_time,omitempty"`
	UpdateTime   int64  `xorm:"not null default 0 comment('修改时间') INT" json:"update_time,omitempty"`
	Status       int64  `xorm:"not null default 0 comment('状态 0 正常 1 限制访问') SMALLINT" json:"status,-"`
	ViewCount    int64  `xorm:"not null default 0 comment('查看次数') INT" json:"view_count,omitempty"`
	LikeCount    int64  `xorm:"not null default 0 comment('点赞次数') INT" json:"like_count,omitempty"`
	CommentCount int64  `xorm:"not null default 0 comment('评论数') INT" json:"comment_count,omitempty"`
	Tags         string `xorm:"not null default '' comment('标签ID 1,2,3,4') VARCHAR(255)" json:"tags,omitempty" form:"tags" binding:"required"`
	Category     int64  `xorm:"not null default 0 comment('文章分类') INT" json:"category,omitempty" form:"category" binding:"required"`
	Top          int64  `xorm:"not null default 0 comment('是否置顶  0 未置顶 1 已置顶') SMALLINT" json:"top,omitempty"`
	TopTime      int64  `xorm:"not null default 0 comment('置顶时间') INT" json:"top_time,omitempty"`
	TopTimeEnd   int64  `xorm:"not null default 0 comment('置顶结束时间') INT" json:"top_time_end,omitempty"`
}

//TableName 返回自定义表名
func (b Blog) TableName() string {
	return "blog"
}

//NewBlog 实例化blog对象
func NewBlog() *Blog {
	return &Blog{}
}

// AddBlog 发布
func (b *Blog) AddBlog() (bool, error) {
	b.CreateTime = time.Now().Unix()
	b.UpdateTime = time.Now().Unix()
	b.UpdateTime = time.Now().Unix()
	_, err := b.GetDb().Insert(b)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 修改
// 删除
// 查看
// 点赞

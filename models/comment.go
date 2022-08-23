package models

type Comment struct {
	Base 			  `xorm:"-"`
	Id         int    `xorm:"not null pk autoincr comment('主键ID') INT"`
	Uid        int    `xorm:"not null default 0 comment('用户ID') INT"`
	BlogId     int    `xorm:"not null default 0 comment('文章ID') INT"`
	Content    string `xorm:"not null default '' comment('评论内容') VARCHAR(255)"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT"`
	UpdateTime int    `xorm:"not null default 0 comment('修改时间') INT"`
	Status     int    `xorm:"not null default 0 comment('状态 0 正常 1 已删除') SMALLINT"`
	ViewCount  int    `xorm:"not null default 0 comment('查看次数') INT"`
	LikeCount  int    `xorm:"not null default 0 comment('点赞次数') INT"`
	Top        int    `xorm:"not null default 0 comment('是否置顶 0 未置顶 1 已置顶') SMALLINT"`
	TopTime    int    `xorm:"not null default 0 comment('置顶时间') INT"`
	TopTimeEnd int    `xorm:"not null default 0 comment('置顶结束时间') INT"`
}

//TableName 返回表名
func (c Comment) TableName() string {
	return "comment"
}

//NewComment 实例化comment对象
func NewComment() *Comment {
	return &Comment{}
}

// 发布
// 修改
// 删除
// 查看
// 点赞
// 回复

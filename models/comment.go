package models

type Comment struct {
	BaseModel
	UserID    uint   // 用户id
	Content   string // 内容
	PostID    uint   // 文章id
	ReadState bool   `gorm:"default:'0'"` // 阅读状态
	//Replies []*Comment // 评论
	NickName  string `gorm:"-"`
	AvatarUrl string `gorm:"-"`
	GithubUrl string `gorm:"-"`
}

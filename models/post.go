package models

import (
	"database/sql"
	"strconv"
)

type Post struct {
	BaseModel
	JID          string     `json:"id" gorm:"column:jid;unique"` // joplin ID
	Title        string     // title
	Body         string     // body
	View         int        // view count
	Tags         []*Tag     `gorm:"-"` // tags of post
	Comments     []*Comment `gorm:"-"` // comments of post
	CommentTotal int        `gorm:"-"` // count of comment
}

func GetPostByID(id string) (*Post, error) {
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	var post Post
	err = DB.First(&post, "id = ?", pid).Error
	return &post, err
}

func ListAllPost(tag string, pageIndex, pageSize int) ([]*Post, error) {
	return _listPost(tag, true, pageIndex, pageSize)
}

func GetResumePost() (*Post, error) {
	var post Post
	err = DB.First(&post, "title = ?", "Resume").Error
	return &post, err
}

func _listPost(tag string, published bool, pageIndex, pageSize int) (posts []*Post, err error) {
	if len(tag) > 0 {
		_, err := strconv.ParseUint(tag, 10, 64)
		if err != nil {
			return nil, err
		}
		var rows *sql.Rows
		if published {
			//if pageSize > 0 {
			//	rows, err = DB.Raw("select p.* from posts p inner join post_tags pt on p.id = pt.post_id where pt.tag_id = ?  order by created_at desc limit ? offset ?", tagID, true, pageSize, (pageIndex-1)*pageSize).Rows()
			//} else {
			//	rows, err = DB.Raw("select p.* from posts p inner join post_tags pt on p.id = pt.post_id where pt.tag_id = ? and p.is_published = ? order by created_at desc", tagID, true).Rows()
			//}
			rows, err = DB.Raw("select * from posts where posts.title != ?", "Resume").Rows() //
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			for rows.Next() {
				var post Post
				DB.ScanRows(rows, &post)
				posts = append(posts, &post)
			}
		}
	} else {
		err = DB.Order("id desc").Find(&posts).Error
	}
	return
}

func (p *Post) Insert() error {
	return DB.Where(Post{Title: p.Title}).FirstOrCreate(&p).Error // TODO: create or update
}

func (p *Post) UpdateContent() error {
	return DB.Model(&p).Where("jid = ?", p.JID).Updates(map[string]interface{}{
		"title":        p.Title,
		"body":         p.Body,
		"updated_time": p.UpdatedTime,
	}).Error
}

func (p *Post) UpdateView() error {
	return DB.Model(p).Updates(map[string]interface{}{
		"view": p.View,
	}).Error
}

func GetPostByJId(jid string) (*Post, error) {
	var post Post
	err := DB.First(&post, "jid = ?", jid).Error
	return &post, err
}

//  checkout post need update
func CheckPostStatus(jid string, updateTime int64) string {
	var post Post
	err = DB.Where("jid = ?", jid).First(&post).Error
	if err != nil {
		if err.Error() == "record not found" {
			return "create"
		}
	}
	if updateTime > post.UpdatedTime {
		return "update"
	}
	return "noChange"
}

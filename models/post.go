package models

import (
	"database/sql"
	"strconv"
)

type Post struct {
	BaseModel
	Title        string     // title
	Body         string     // body
	View         int        // view count
	Tags         []*Tag     `gorm:"-"` // tags of post
	Comments     []*Comment `gorm:"-"` // comments of post
	CommentTotal int        `gorm:"-"` // count of comment
}

type postTag struct {
	BaseModel
	PostID uint
	TagID  uint
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

func ListAllPost(tag string) ([]*Post, error) {
	return _listPost(tag, true, 0, 0)
}

func _listPost(tag string, published bool, pageIndex, pageSize int) (posts []*Post, err error) {
	if len(tag) > 0 {
		tagID, err := strconv.ParseUint(tag, 10, 64)
		if err != nil {
			return nil, err
		}
		var rows *sql.Rows
		if published {
			if pageSize > 0 {
				rows, err = DB.Raw("select p.* from posts p inner join post_tags pt on p.id = pt.post_id where pt.tag_id = ? and p.is_published = ? order by created_at desc limit ? offset ?", tagID, true, pageSize, (pageIndex-1)*pageSize).Rows()
			} else {
				rows, err = DB.Raw("select p.* from posts p inner join post_tags pt on p.id = pt.post_id where pt.tag_id = ? and p.is_published = ? order by created_at desc", tagID, true).Rows()
			}
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
	}
	return
}

func (p *Post) Insert() error {
	return DB.Where(Post{Title: p.Title}).FirstOrCreate(&p).Error // TODO: create or update
}

func (p *Post) UpdateView() error {
	return DB.Model(p).Updates(map[string]interface{}{
		"view": p.View,
	}).Error
}

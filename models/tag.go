package models

import "strconv"

type Tag struct {
	BaseModel
	Name  string `json:"name"`
	State int    `json:"state"`
}

type postTag struct {
	BaseModel
	PostID uint
	TagID  uint
}

func ListTagsByPostID(id string) (tags []*Tag, err error) {
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	rows, err := DB.Raw("select t.* from tags t inner join post_tags pt on t.id = pt.tag_id where pt.post_id = ?", uint(pid)).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tag Tag
		DB.ScanRows(rows, &tag)
		tags = append(tags, &tag)
	}
	return tags, nil
}

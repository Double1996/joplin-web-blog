package models

type EverNoteAt struct {
	Token               string //
	Secret              string //
	EdamExpires         string // Token到期的日期
	EdamShard           string
	EdamUserId          string
	EdamNoteStoreUrl    string
	EdamWebApiUrlPrefix string
}

// 写入
func (at *EverNoteAt) Insert() error {
	return DB.Create(at).Error
}

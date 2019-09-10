package controller

import (
	"fmt"
	"github.com/double1996/smart-evernote-blog/models"
	"github.com/gin-gonic/gin"
)
import "github.com/double1996/smart-evernote-blog/pkg/evernote"

// 获取印象笔记的回调链接
func GetEverNoteCallback(c *gin.Context) {
	oauthVerifier := c.Query("oauth_verifier")
	fmt.Println(evernote.Rq)
	as, err := evernote.Client.GetAuthorizedToken(evernote.Rq, oauthVerifier)
	if err != nil {
		return
	}
	eas := &models.EverNoteAt{
		Token:               as.Token,
		Secret:              as.Secret,
		EdamExpires:         as.AdditionalData["edam_expires"],
		EdamShard:           as.AdditionalData["edam_shard"],
		EdamUserId:          as.AdditionalData["edam_userId"],
		EdamNoteStoreUrl:    as.AdditionalData["edam_noteStoreUrl"],
		EdamWebApiUrlPrefix: as.AdditionalData["edam_webApiUrlPrefix"],
	}
	ierr := eas.Insert()
	if ierr != nil {
		fmt.Println(ierr)
	}

}

func GetNewEverNoteByWebhook(c *gin.Context) {
	fmt.Println("111111111111")
}

func GetUpdateEvernoteByWebhook(c *gin.Context) {

}

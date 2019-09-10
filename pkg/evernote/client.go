package evernote

import (
	"github.com/double1996/evernote-sdk-golang/client"
	"github.com/mrjones/oauth"
)

const (
	EvernoteKey         string = ""
	EvernoteSecret      string = ""
	EvernoteAuthorToken string = ""
	callBackURL         string = ""
)

var (
	Client *client.EvernoteClient
	Rq     *oauth.RequestToken
)

// 判断数据库存储的AS是否有效
func CheckoutAccessTokenIsExist() {

}

/*
   use webhooks  https://dev.yinxiang.com/doc/articles/polling_notification.php
*/
//
//func SyncEverNoteClient() {
//	clientCtx, _ := context.WithTimeout(context.Background(), time.Duration(15)*time.Second)
//	c := client.NewClient(EvernoteKey, EvernoteSecret, client.YINXIANG)
//	us, err := c.GetUserStore()
//	if err != nil {
//	}
//	c.GetRequestToken(callBackURL)
//	userUrls, err := us.GetUserUrls(clientCtx, EvernoteAuthorToken)
//	if err != nil {
//	}
//	fmt.Println(userUrls)
//}

//func SyncEverNoteClient() {
//	clientCtx, _ := context.WithTimeout(context.Background(), time.Duration(15)*time.Second)
//	Client = client.NewClient(EvernoteKey, EvernoteSecret, client.YINXIANG)
//	us, err := Client.GetUserStore()
//	if err != nil {
//		logger.Error(err.Error())
//	}
//	var url string
//	var err error
//	Rq, url, err = Client.GetRequestToken(callBackURL)
//	if err != nil {
//		logger.Error(err.Error())
//		return
//	}
//	if url != "" && Rq != nil {
//
//	}
//}

//
//func FirstBathchInsertNotes(at *oauth.AccessToken) {
//	clientCtx, _ := context.WithTimeout(context.Background(), time.Duration(15)*time.Second)
//	us, err := Client.GetUserStore()
//	if err != nil {
//		logger.Error(err.Error())
//	}
//	userUrls, err := us.GetUserUrls(clientCtx, at.Token)
//	if err != nil {
//		logger.Error(err.Error())
//	}
//	if userUrls == nil {
//		logger.Error("not found evernote note")
//		return
//	}
//	ns, err := Client.GetNoteStoreWithURL(userUrls.GetNoteStoreUrl())
//	if err != nil {
//		logger.Error(err.Error())
//	}
//	note, err := ns.GetDefaultNotebook(clientCtx, EvernoteAuthorToken)
//	if err != nil {
//		logger.Error(err.Error())
//	}
//	if note == nil {
//		logger.Error(err.Error())
//	}
//}

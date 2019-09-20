package joplin

import (
	"encoding/json"
	"fmt"
	"github.com/double1996/joplin-web-blog/config"
	"github.com/double1996/joplin-web-blog/models"
	"time"

	"io/ioutil"
	"net/http"
)

type joplinTagRes struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type tagNotesRes struct {
	ID string `json:"id"`
}

var client = &http.Client{Timeout: time.Duration(5 * time.Second)}

func requestBlogTags() (tagsRes []joplinTagRes, err error) {
	reqTagsURL := fmt.Sprintf("http://%s/tags?token=%s&fields=id,title", config.Conf.Joplin.Host, config.Conf.Joplin.Token)
	req, err := http.NewRequest("GET", reqTagsURL, nil)
	if err != nil {
		return tagsRes, err
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return tagsRes, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tagsRes, err
	}
	if err := json.Unmarshal(body, &tagsRes); err != nil {
		return tagsRes, err
	}
	return tagsRes, nil
}

func requestPostByTags(id string) (noteRes []tagNotesRes, err error) {
	reqTagsURL := fmt.Sprintf("http://%s/tags/%s/notes?token=%s&fields=id", config.Conf.Joplin.Host, id, config.Conf.Joplin.Token)
	req, err := http.NewRequest("GET", reqTagsURL, nil)
	if err != nil {
		return noteRes, err
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return noteRes, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return noteRes, err
	}
	if err := json.Unmarshal(body, &noteRes); err != nil {
		return noteRes, err
	}
	return noteRes, nil
}

func requestNoteById(id string) {
	reqTagsURL := fmt.Sprintf("http://%s/notes/%s?token=%s&fields=title,body,created_time,updated_time", config.Conf.Joplin.Host, id, config.Conf.Joplin.Token)
	req, err := http.NewRequest("GET", reqTagsURL, nil)
	if err != nil {
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	var post models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
	}
	err = post.Insert()
	if err != nil {
	}
}

func SyncJoplinBlog() {
	tagsRes, err := requestBlogTags()
	if err != nil {

	}
	for _, tag := range tagsRes {
		if tag.Title == "blog" {
			noteRes, err := requestPostByTags(tag.ID)
			if err != nil {
			}
			for _, note := range noteRes {
				requestNoteById(note.ID)
			}
		}
	}
}

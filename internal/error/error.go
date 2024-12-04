package error

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_TAG       = 10001
	ERROR_EXIST_TAG_FAIL  = 10002
	ERROR_NOT_EXIST_TAG   = 10003
	ERROR_GET_TAGS_FAIL   = 10004
	ERROR_COUNT_TAG_FAIL  = 10005
	ERROR_ADD_TAG_FAIL    = 10006
	ERROR_EDIT_TAG_FAIL   = 10007
	ERROR_DELETE_TAG_FAIL = 10008
	ERROR_EXPORT_TAG_FAIL = 10009
	ERROR_IMPORT_TAG_FAIL = 10010

	ERROR_NOT_EXIST_ARTICLE        = 10011
	ERROR_CHECK_EXIST_ARTICLE_FAIL = 10012
	ERROR_ADD_ARTICLE_FAIL         = 10013
	ERROR_DELETE_ARTICLE_FAIL      = 10014
	ERROR_EDIT_ARTICLE_FAIL        = 10015
	ERROR_COUNT_ARTICLE_FAIL       = 10016
	ERROR_GET_ARTICLES_FAIL        = 10017
	ERROR_GET_ARTICLE_FAIL         = 10018
	ERROR_GEN_ARTICLE_POSTER_FAIL  = 10019

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
)

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "Request parameter error",
	ERROR_EXIST_TAG:                 "Tag name already exists",
	ERROR_EXIST_TAG_FAIL:            "Failed to obtain an existing label",
	ERROR_NOT_EXIST_TAG:             "The label does not exist",
	ERROR_GET_TAGS_FAIL:             "Failed to get all tags",
	ERROR_COUNT_TAG_FAIL:            "Statistical label failed",
	ERROR_ADD_TAG_FAIL:              "Failed to add a label",
	ERROR_EDIT_TAG_FAIL:             "Failed to modify the label",
	ERROR_DELETE_TAG_FAIL:           "Failed to delete the label",
	ERROR_EXPORT_TAG_FAIL:           "Export label failed",
	ERROR_IMPORT_TAG_FAIL:           "Failed to import tags",
	ERROR_NOT_EXIST_ARTICLE:         "The article does not exist",
	ERROR_ADD_ARTICLE_FAIL:          "Failed to add an article",
	ERROR_DELETE_ARTICLE_FAIL:       "Failed to delete article",
	ERROR_CHECK_EXIST_ARTICLE_FAIL:  "Check if there is a failure in the article",
	ERROR_EDIT_ARTICLE_FAIL:         "Failed to modify the article",
	ERROR_COUNT_ARTICLE_FAIL:        "Statistical article failed",
	ERROR_GET_ARTICLES_FAIL:         "Failed to get multiple articles",
	ERROR_GET_ARTICLE_FAIL:          "Failed to obtain a single article",
	ERROR_GEN_ARTICLE_POSTER_FAIL:   "Failed to generate an article poster",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token authentication failed",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token timed out",
	ERROR_AUTH_TOKEN:                "Token Generation failed",
	ERROR_AUTH:                      "Token error",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "Failed to save the picture",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "Failed to check the picture",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "Verify that the picture is wrong, there is a problem with the image format or size",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		fmt.Errorf(err.Key, err.Message)
	}

	return
}

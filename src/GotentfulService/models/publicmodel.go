package models

import (
	"time"
)

//
//Public JSON Models
//
type SysVersionPublic struct {
	SysIDBase
	Locale    string    `json:"locale"`
	Version   int       `json:"version"`
	Revision  int       `json:"revision"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type SysSpacePublic struct {
	SysVersionPublic
	Space struct {
		Sys SysLink `json:"sys"`
	} `json:"space"`
}

type SysContentTypePublic struct {
	SysSpacePublic
	ContentType struct {
		Sys SysLink `json:"sys"`
	} `json:"contentType"`
}

type SpacePublic struct {
	Space
}

type LocalePublic struct {
	Locale
	Sys SysSpacePublic `json:"sys"`
}

type ContentTypePublic struct {
	ContentType
	Sys SysSpacePublic `json:"sys"`
}

type EntryPublic struct {
	Entry
	Sys SysContentTypePublic `json:"sys"`
}

type FileContentPublic struct {
	ContentType string `json:"contentType"`
	FileName    string `json:"fileName"`
	Upload      string `json:"upload"`
	Details     struct {
		Image struct {
			Height int `json:"height"`
			Width  int `json:"width"`
		} `json:"image"`
		Size int `json:"size"`
	} `json:"details"`
}

type AssetPublic struct {
	Fields struct {
		Title       string            `json:"title"`
		Description string            `json:"description"`
		File        FileContentPublic `json:"file"`
	} `json:"fields"`
	Sys SysSpacePublic `json:"sys"`
}

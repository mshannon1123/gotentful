package models

import (
	"time"
)

//
//System JSON Models
//
type (
	SysBase struct {
		Type string `json:"type"`
	}

	SysSchema struct {
		SysBase
		LinkType string `json:"linkType"`
	}

	SysIDBase struct {
		SysBase
		ID string `json:"id"`
	}

	SysVersion struct {
		SysIDBase
		Version   int       `json:"version"`
		CreatedAt time.Time `json:"createdAt"`
		CreatedBy struct {
			Sys SysLink `json:"sys"`
		} `json:"createdBy"`
		UpdatedAt time.Time `json:"updatedAt"`
		UpdatedBy struct {
			Sys SysLink `json:"sys"`
		} `json:"updatedBy"`
	}

	SysSpace struct {
		SysVersion
		Space struct {
			Sys SysLink `json:"sys"`
		} `json:"space"`
	}

	SysContentType struct {
		SysSpace
		ContentType struct {
			Sys SysLink `json:"sys"`
		} `json:"contentType"`
	}

	SysArrayResult struct {
		Sys   SysBase       `json:"sys"`
		Total int           `json:"total"`
		Skip  int           `json:"skip"`
		Limit int           `json:"limit"`
		Items []interface{} `json:"items"`
	}

	SysPublish struct {
		FirstPublishedAt time.Time `json:"firstPublishedAt"`
		PublishedCounter int       `json:"publishedCounter"`
		PublishedAt      time.Time `json:"publishedAt"`
		PublishedBy      struct {
			Sys SysLink `json:"sys"`
		} `json:"publishedBy"`
		PublishedVersion int `json:"publishedVersion"`
	}

	SysLink struct {
		SysIDBase
		LinkType string `json:"linkType"`
	}

	Field struct {
		ID          string    `json:"id"`
		Type        string    `json:"type"`
		Name        string    `json:"name"`
		LinkType    string    `json:"linkType"`
		Required    bool      `json:"required"`
		Localized   bool      `json:"localized"`
		Items       SysSchema `json:"items"`
		Disabled    bool      `json:"disabled"`
		Validations []string  `json:"validations"`
	}

	FileContent struct {
		ContentType string `json:"contentType"`
		FileName    string `json:"fileName"`
		Upload      string `json:"upload"`
	}

	Error struct {
		Details struct {
			Sys SysBase `json:"sys"`
		} `json:"details"`
		Message   string    `json:"message"`
		RequestID string    `json:"requestId"`
		Sys       SysIDBase `json:"sys"`
	}
)

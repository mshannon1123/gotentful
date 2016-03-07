package models

type Space struct {
	Name string `json:"name"`
}

type SpaceResult struct {
	Sys SysVersion `json:"sys"`
	Space
}

type Locale struct {
	Name         string  `json:"name"`
}

type LocaleResult struct {
	Sys SysSpace `json:"sys"`
	Locale
}

type ContentType struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	DisplayField string  `json:"displayField"`
	Fields       []Field `json:"fields"`
}

type ContentTypeResult struct {
	Sys SysSpace `json:"sys"`
	ContentType
}

type PublishContentType struct {
	Sys SysPublish `json:"sys"`
	ContentType
}

type Entry struct {
	Fields map[string]map[string]interface{} `json:"fields"`
}

type EntryResult struct {
	Sys SysContentType `json:"sys"`
	Entry
}

type PublishEntry struct {
	Sys SysPublish `json:"sys"`
	Entry
}

type Asset struct {
	Fields struct {
		Title       map[string]string      `json:"title"`
		Description map[string]string      `json:"description"`
		File        map[string]FileContent `json:"file"`
	} `json:"fields"`
}

type AssetResult struct {
	Sys SysContentType `json:"sys"`
	Asset
}

type PublishAsset struct {
	Sys SysPublish `json:"sys"`
	Asset
}

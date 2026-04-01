package models

type Class struct {
	TypeID   any    `json:"type_id"`
	TypeName string `json:"type_name"`
}
type Filter struct {
	Key   string       `json:"key"`
	Name  string       `json:"name"`
	Value []SortOption `json:"value"`
}

// SortOption 筛选值项
type SortOption struct {
	N string `json:"n"`
	V string `json:"v"`
}
type Vod struct {
	VodID       any    `json:"vod_id"`
	VodName     string `json:"vod_name"`
	VodPic      string `json:"vod_pic"`
	VodRemarks  string `json:"vod_remarks,omitempty"`
	VodYear     string `json:"vod_year,omitempty"`
	VodArea     string `json:"vod_area,omitempty"`
	VodActor    string `json:"vod_actor,omitempty"`
	VodDirector string `json:"vod_director,omitempty"`
	VodContent  string `json:"vod_content,omitempty"`
	VodPicSlide string `json:"vod_pic_slide,omitempty"`
	VodPlayFrom string `json:"vod_play_from"`
	VodPlayURL  string `json:"vod_play_url"`
}

type Result struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message,omitempty"`
	Url       string              `json:"url,omitempty"`
	Page      int                 `json:"page,omitempty"`
	PageCount int                 `json:"pagecount,omitempty"`
	Limit     int                 `json:"limit,omitempty"`
	Total     int                 `json:"total,omitempty"`
	List      []Vod               `json:"list,omitempty"`
	Class     []Class             `json:"class,omitempty"`
	Parse     int                 `json:"parse,omitempty"`
	Filters   map[string][]Filter `json:"filters,omitempty"`
	Header    any                 `json:"header,omitempty"`
	Danmaku   string              `json:"danmaku,omitempty"`
	Format    string              `json:"format,omitempty"`
}

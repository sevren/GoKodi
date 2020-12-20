package models

type TVResponse struct {
	Limits  Limit    `json:"limits"`
	TvShows Shows `json:"tvshows"`
}

type Limit struct {
	Start int `json:"start"`
	End   int `json:"end"`
	Total int `json:"total"`
}

type TvShowsT struct {
	Label    string `json:"label"`
	TvShowId int    `json:"tvshowid"`
}

type Shows []TvShowsT


func (tv Shows) String(i int) string {
	return tv[i].Label
}

func (tv Shows) Len() int {
	return len(tv)
}
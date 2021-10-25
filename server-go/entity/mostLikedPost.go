package entity

type MostLikedPost struct {
	Title    string `json:"title"`
	PostDate string `json:"postdate"`
	Likes    int    `json:"likes"`
}

func (MostLikedPost) TableName() string {
	return "post"
}

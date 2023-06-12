package model

type GameModel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

func (GameModel) TableName() string {
	return "games"
}

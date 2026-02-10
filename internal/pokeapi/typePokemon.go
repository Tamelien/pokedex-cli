package pokeapi

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []Stat `json:"stats"`
	Types          []Type `json:"types"`
}

type Stat struct {
	BaseStat int `json:"base_stat"`
	NameStat struct {
		Name string
	} `json:"stat"`
}

type Type struct {
	TypeName struct {
		Name string
	} `json:"type"`
}

package docs

type Item struct {
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	FullUnit  string  `json:"fullUnit,omitempty" bson:"fullUnit,omitempty"`
	FullCount int     `json:"fullCount,omitempty" bson:"fullCount,omitempty"`
	Unit      string  `json:"unit"`
	Count     int     `json:"count"`
	Price     float64 `json:"price"`
	Money     int     `json:"money"`
}

type Bill struct {
	Id string `json:"id" bson:"_id"`
	//Date     string `json:"date"`
	Clerk    string `json:"clerk"`
	Client   string `json:"client"`
	Items    []Item `json:"items"`
	Money    int    `json:"money"`
	Received int    `json:"received"`
	IsReturn bool   `json:"isReturn" bson:"isReturn"`
	CreateAt int    `json:"createAt" bson:"createAt"`
}

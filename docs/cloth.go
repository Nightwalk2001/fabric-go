package docs

type Cloth struct {
	Name   string `json:"name" bson:"_id"`
	Pinyin string `json:"pinyin" bson:"pinyin"`
	//Color  string  `json:"color" bson:"color"`
	Unit string `json:"unit,omitempty" bson:"unit,omitempty"`
	//Price  float64 `json:"price,omitempty" bson:"price,omitempty"`
}

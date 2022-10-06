package docs

type Client struct {
	Name    string `json:"name" bson:"_id"`
	Pinyin  string `json:"pinyin" bson:"pinyin"`
	Phone   string `json:"phone,omitempty" bson:"phone,omitempty"`
	Address string `json:"address,omitempty" bson:"address,omitempty"`
}

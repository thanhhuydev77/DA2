package Model

type Item struct {
	Category            string   `json:"category"`
	Name                string   `json:"name"`
	Subcategory         string   `json:"subCategory"`
	CurrentPrice        float64  `json:"currentPrice"`
	RawPrice            float64  `json:"rawPrice"`
	Currency            string   `json:"currency"`
	Discount            int      `json:"discount"`
	LikesCount          int      `json:"likeCount"`
	Brand               string   `json:"brand"`
	BrandUrl            string   `json:"brandUrl"`
	CodCountry          []string `json:"codCountry"`
	Variation0Color     string   `json:"variation0Color"`
	Variation1Color     string   `json:"variation1Color"`
	Variation0Thumbnail string   `json:"variation0Thumbnail"`
	Variation0Image     string   `json:"variation0Image"`
	Variation1Thumbnail string   `json:"variation1Thumbnail"`
	Variation1Image     string   `json:"variation1Image"`
	ImageUrl            string   `json:"imageUrl"`
	Url                 string   `json:"url"`
	Id                  int      `json:"id"`
	Model               string   `json:"model"`
	IsNew               bool     `json:"isNew"`
}

type ItemCategory struct {
	Id           int      `json:"id"`
	CurrentPrice float64  `json:"currentPrice"`
	RawPrice     float64  `json:"rawPrice"`
	LikesCount   int      `json:"likeCount"`
	IsNew        bool     `json:"isNew"`
	CodCountry   []string `json:"codCountry"`
	Brand        string   `json:"brand"`
	Color        []string `json:"color"`
}
type ItemIdListResult struct {
	ItemIds []string `json:"id"`
}

func RemoveID(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
func RemoveDuplicateID(list []ItemUtility) []ItemUtility {
	for i, value1 := range list {
		for j, value2 := range list {
			if value1.ItemId == value2.ItemId && i != j {
				list = append(list[:j], list[j+1:]...)
			}
		}
	}
	return list
}

type ItemUtility struct {
	ItemId       string
	UtilityValue float64
}

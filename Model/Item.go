package Model

type Item struct {
	Category            string
	Name                string
	Subcategory         string
	CurrentPrice        float64
	RawPrice            float64
	Currency            string
	Discount            int
	LikesCount          int
	Brand               string
	BrandUrl            string
	CodCountry          []string
	Variation0Color     string
	Variation1Color     string
	Variation0Thumbnail string
	Variation0Image     string
	Variation1Thumbnail string
	Variation1Image     string
	ImageUrl            string
	Url                 string
	Id                  int
	Model               string
	IsNew               bool
}

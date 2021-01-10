package Read

import (
	"Project2/Model"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileProductReviewCSV(path string) (map[string][]Model.ItemUserRating, map[string][]Model.UserItemRating) {
	csvFile, err := os.Open(path)
	if err != nil {
		fmt.Print(err)
	}
	lines := csv.NewReader(bufio.NewReader(csvFile))

	itemRating := make(map[string][]Model.ItemUserRating)
	userRating := make(map[string][]Model.UserItemRating)

	for {
		line, error := lines.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		rating, err := strconv.Atoi(line[17])
		if err != nil {
			continue
		}
		ID := line[0]
		//Name:= line[9]
		ReviewUsername := line[23]

		// next if id product or
		if ID == "" || ReviewUsername == "" || rating < 1 || rating > 5 {
			continue
		}

		if len(itemRating[ReviewUsername]) > 0 {
			itemRating[ReviewUsername] = append(itemRating[ReviewUsername], Model.ItemUserRating{
				Item:   ID,
				Rating: rating,
			})
		} else {
			itemRating[ReviewUsername] = []Model.ItemUserRating{
				{
					Item:   ID,
					Rating: rating,
				},
			}
		}

		if len(userRating[ID]) > 0 {
			userRating[ID] = append(userRating[ID], Model.UserItemRating{
				User:   ReviewUsername,
				Rating: rating,
			})
		} else {
			userRating[ID] = []Model.UserItemRating{
				{
					User:   ReviewUsername,
					Rating: rating,
				},
			}
		}
	}
	return itemRating, userRating
}

func existedinList(id string, listid []string) bool {
	for _, value := range listid {
		if id == value {
			return true
		}
	}
	return false
}

func GetItemInfo(listid []string) []Model.ItemOutput {
	var ItemList []Model.ItemOutput
	read := 0
	csvFile, _ := os.Open("storage/shoes.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if read == 10 {
			break
		}

		if existedinList(line[20], listid) {
			//listid.ItemIds = Model.RemoveID(listid.ItemIds, Model.GetIndex(listid.ItemIds, line[20]))
			var rawPrice float64
			var discount int
			var likesCount int
			var id int
			var currentPrice float64
			read += 1
			if tempid, Error4 := strconv.Atoi(line[20]); Error4 != nil {
				continue
			} else {
				id = tempid
			}

			if tempcurrentPrice, Error := strconv.ParseFloat(line[3], 2); Error != nil {
				continue
			} else {
				currentPrice = tempcurrentPrice
			}
			if temprawPrice, Error1 := strconv.ParseFloat(line[4], 2); Error1 != nil {
				continue
			} else {
				rawPrice = temprawPrice
			}

			if tempdiscount, Error2 := strconv.Atoi(line[6]); Error2 != nil {
				continue
			} else {
				discount = tempdiscount
			}

			if templikesCount, Error3 := strconv.Atoi(line[7]); Error3 != nil {
				continue
			} else {
				likesCount = templikesCount
			}
			codCountry := strings.Split(line[11], ",")

			ItemList = append(ItemList, Model.ItemOutput{
				Name:            line[2],
				CurrentPrice:    currentPrice,
				RawPrice:        rawPrice,
				Currency:        line[5],
				Discount:        discount,
				LikesCount:      likesCount,
				Brand:           line[9],
				CodCountry:      codCountry,
				Variation0Color: line[12],
				Variation1Color: line[13],
				ImageUrl:        line[18],
				Id:              id,
			})
		}
	}
	return ItemList
}

func RecommendUserBasedFunc(filename string, userId string) []Model.ItemOutput {
	csvFile, err := os.Open("./Storage/" + filename)
	if err != nil {
		fmt.Print(err)
	}
	lines := csv.NewReader(bufio.NewReader(csvFile))
	lines.Comma = ','
	lines.FieldsPerRecord = -1
	ids := []string{}
	for {
		line, error := lines.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		ID := line[0]
		if ID != userId {
			continue
		} else {
			length := len(line)
			for i := 1; i < length; i++ {
				ids = append(ids, line[i])
			}
		}
		break
	}
	if len(ids) == 0 {
		return []Model.ItemOutput{}
	}
	products := GetItemInfo(ids)
	return products
}

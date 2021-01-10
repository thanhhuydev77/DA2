package processing

import (
	"Project2/Model"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetContentRecommend(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	item := Model.Item{}
	result := Model.ItemIdListResult{}

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		io.WriteString(w, `{"message": "wrong format!"}`)
		return
	}
	//if  splitOk == false {
	//	io.WriteString(w, `{"message": "waiting for minutes, File is processing...!"}`)
	//	return
	//}

	// get item list
	item.Subcategory = FindCategory(item.Id)
	if item.Subcategory == "" {
		io.WriteString(w, `{"message":"item is not available!"}`)
		return
	}
	ItemList := ReadUtilityTable(strconv.Itoa(item.Id), item.Subcategory)
	//get -- sort -- get 10 top record
	ItemList = Model.RemoveDuplicateID(ItemList)
	sort.SliceStable(ItemList, func(i, j int) bool {
		return ItemList[i].UtilityValue > ItemList[j].UtilityValue
	})

	for _, value := range ItemList {
		if len(result.ItemIds) < 10 && value.ItemId != strconv.Itoa(item.Id) {
			result.ItemIds = append(result.ItemIds, value.ItemId)
			fmt.Print(value.ItemId + "\n")
		}
	}
	//return list item id
	jsonresult, _ := json.Marshal(GetItemInfo(result))
	io.WriteString(w, string(jsonresult))
}
func GetItemInfo2(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var item Model.ItemIdListResult
	result := Model.ItemIdListResult{}

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		io.WriteString(w, `{"message": "wrong format!"}`)
		return
	}
	//if  splitOk == false {
	//	io.WriteString(w, `{"message": "waiting for minutes, File is processing...!"}`)
	//	return
	//}

	// get item list
	for i := range item.ItemIds {
		result.ItemIds = append(result.ItemIds, item.ItemIds[i])
	}
	//return list item id
	jsonresult, _ := json.Marshal(GetItemInfo(result))
	io.WriteString(w, string(jsonresult))
}
func existedinList(id string, listid Model.ItemIdListResult) bool {
	for _, value := range listid.ItemIds {
		if id == value {
			return true
		}
	}
	return false
}

func GetItemInfo(listid Model.ItemIdListResult) []Model.ItemOutput {
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

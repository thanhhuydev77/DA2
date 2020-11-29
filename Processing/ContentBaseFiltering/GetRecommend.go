package processing

import (
	"Project2/Model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
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
	jsonresult, _ := json.Marshal(result)
	io.WriteString(w, string(jsonresult))
}

package collaborative

import (
	"Project2/Model"
	"encoding/csv"
	"log"
	"math"
	"os"
	"sort"
)

type Similarity struct {
	userId string
	score  float64
}

func UserSimilarity(itemRating map[string][]Model.ItemUserRating, top int, filename string) {
	dataWrite := [][]string{}
	amount := 5000
	users := GetUsersMap(itemRating)
	length := len(users)

	for i := 0; i < length-1; i++ {
		userSimilarity := []Similarity{}
		score := 0.0
		for j := 1; j < length; j++ {
			if i == j {
				continue
			}
			user1rating := itemRating[users[i]]
			user2rating := itemRating[users[j]]
			tuso := 0
			mauso := 0.0
			sqrt1 := 0.0
			sqrt2 := 0.0
			for _, item1 := range user1rating {
				sqrt1 += math.Pow(float64(item1.Rating), float64(2))
				for _, item2 := range user2rating {
					if item1.Item == item2.Item {
						tuso += item1.Rating * item2.Rating
					}
				}
			}
			for _, item2 := range user2rating {
				sqrt2 += math.Pow(float64(item2.Rating), float64(2))
				for _, item1 := range user1rating {
					if item1.Item == item2.Item {
						tuso += item1.Rating * item2.Rating
					}
				}
			}
			tuso /= 2
			mauso = math.Sqrt(sqrt1) * math.Sqrt(sqrt2)

			if mauso == 0 {
				score = 0
			} else {
				score = float64(tuso) / mauso
			}
			userSimilarity = append(userSimilarity, Similarity{userId: users[j], score: score})
		}

		sort.SliceStable(userSimilarity, func(i, j int) bool {
			return userSimilarity[i].score > userSimilarity[j].score
		})

		userSimilarityLength := len(userSimilarity)
		sl := top
		if userSimilarityLength < sl {
			sl = userSimilarityLength
		} else if userSimilarityLength == 0 {
			sl = 1
		}

		data := []string{users[i]}
		set := make(map[string]bool)
		for j := 0; j < sl; j++ {
			list := itemRating[userSimilarity[j].userId]
			for _, v := range list {
				if set[v.Item] == false {
					data = append(data, v.Item)
					set[v.Item] = true
				}
			}
		}

		dataWrite = append(dataWrite, data)

		amount -= 1
		if amount <= 0 {
			break
		}
	}

	go writeCleanData(dataWrite, filename)
}

func GetUsersMap(mymap map[string][]Model.ItemUserRating) []string {
	i := 0
	keys := make([]string, len(mymap))
	for k := range mymap {
		keys[i] = k
		i++
	}
	return keys
}

func GetItemMap(mymap map[string][]Model.UserItemRating) []string {
	i := 0
	keys := make([]string, len(mymap))
	for k := range mymap {
		keys[i] = k
		i++
	}
	return keys
}

func writeCleanData(data [][]string, path string) {
	csvData, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer csvData.Close()

	csv := csv.NewWriter(csvData)
	defer csv.Flush()

	for _, value := range data {
		errWrite := csv.Write(value)
		if errWrite != nil {
			log.Fatal("Write file err")
		}
	}
}

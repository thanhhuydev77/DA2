// package main

// import (
// 	"fmt"
// )

// type UserRating struct {
// 	Item   string
// 	Rating int
// }

// type ListUserRating struct {
// 	ListUserRating []UserRating
// }

// // func main() {

// // 	x := make(map[string]ListUserRating)

// // 	listRating := ListUserRating{
// // 		ListUserRating: []UserRating{
// // 			UserRating{
// // 				Item:   "item1",
// // 				Rating: 5,
// // 			},
// // 		},
// // 	}
// // 	x["duy"] = listRating
// // 	fmt.Print(x["duy"].ListUserRating[0])
// // }

// // func main() {
// // 	x := make(map[string][]UserRating)

// // 	listRating := []UserRating{
// // 		{
// // 			Item: "item1",
// // 			Rating: 5,
// // 		},
// // 	}
// // 	x["duy"] = listRating
// // 	x["duy"] = append(x["duy"], UserRating{
// // 		Item: "item2",
// // 		Rating: 4,
// // 	})


// // 	fmt.Print(x["duy"])
// // }
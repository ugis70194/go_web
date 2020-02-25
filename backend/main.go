package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type userData struct {
	Affiliation string `json:"affiliation"`
	BirthYear uint `json:"birth_year"`
	Competitions uint `json:"competitions"`
	Country string `json:"country"`
	FormalCountryName string `json:"formal_country_name"`
	HighestRating uint `json:"highest_rating"`
	Rank uint `json:"rank"`
	Rating uint `json:"rating"`
	Updated string `json:"updated"`
	UserColor string `json:"user_color"`
	Wins uint `json:"wins"`
}

type userInfo struct {
	Data userData  `json:"data"` 
}

var x int = 1
func xorShift() int {
	if x == 0 { x = 1 } 
	x ^= (x << 13)
	x ^= (x >> 17)
	x ^= (x << 15)
	return x
}

func userRating(ctx *gin.Context){
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, X-CSRFToken")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET")
	

	var ratings [2]uint = [2]uint{0, 0}
	var usernames []string = ctx.QueryArray("username")

	for i, username := range(usernames) {
		var userinfo userInfo
		var url string = "https://us-central1-atcoderusersapi.cloudfunctions.net/api/info/username/" + username

		timeout := time.Duration(5 * time.Second)
		client := &http.Client{
			Timeout: timeout, 
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			ctx.JSON(200, gin.H{"rating1" : ratings[0], "rating2" : ratings[1]}) 
			log.Print(err) 
		}

		res, err := client.Do(req)
		if err != nil {
			ctx.JSON(200, gin.H{"rating1" : ratings[0], "rating2" : ratings[1]}) 
			log.Print(err) 
		}

		defer res.Body.Close()	
		body, err := ioutil.ReadAll(res.Body)
		if err := json.Unmarshal(body, &userinfo); err != nil { 
			ctx.JSON(200, gin.H{"rating1" : ratings[0], "rating2" : ratings[1]})
			log.Print(err) 
		}
		ratings[i] = userinfo.Data.Rating
	}
	ctx.JSON(200, gin.H{"rating1" : ratings[0], "rating2" : ratings[1]})
}

type hand struct {
	Kind int `josn:"kind"`
}

func janken(ctx *gin.Context) {
	var myhand hand 
	ctx.BindJSON(&myhand)
	hand1 := myhand.Kind 
	hand2 := xorShift() % 3
	result := (hand1-hand2+3) % 3

	ctx.JSON(200, gin.H{ "result" : result })
}

func main() {
    router := gin.Default()

	router.POST("/janken", janken)
	router.GET("/user_rating", userRating)
	
    router.Run(":8000")
}
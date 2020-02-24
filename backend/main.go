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

var x uint32 = 1
func xorShift() uint32 {
	if x == 0 { x = 1 } 
	x ^= (x << 13)
	x ^= (x >> 17)
	x ^= (x << 15)
	return x
}

func userRating(ctx *gin.Context){
	var userinfo userInfo
	var username string = ctx.Query("username")
	if username == "" { username = "tourist" }

	var url string = "https://us-central1-atcoderusersapi.cloudfunctions.net/api/info/username/" + username

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout, 
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil { log.Fatal(err) }
	
	res, err := client.Do(req)
	if err != nil { log.Fatal(err) }

	defer res.Body.Close()	
	body, err := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(body, &userinfo); err != nil { log.Fatal(err) }
	
	ctx.JSON(200, gin.H{"rating" : userinfo.Data.Rating})
}

func janken(ctx *gin.Context) {
	ctx.JSON(200, gin.H{ "num" : xorShift() % 3 })
}

func main() {
    router := gin.Default()

	router.GET("/janken", janken)
	router.GET("/user_rating", userRating)
	
    router.Run()
}
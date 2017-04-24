package models

import (
	"gopkg.in/mgo.v2/bson"
	"nttdata/app/models/mongodb"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
)

type Restaurant struct {
	//businesses []Businesses `json:"businesses"`
	//businesses []Businesses `json:"businesses"`
	Businesses []struct {
		Categories []struct {
			Alias string `json:"alias"`
			Title string `json:"title"`
		} `json:"categories"`

		Coordinates struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"coordinates"`

		DisplayPhone string  `json:"display_phone"`
		Distance     float64 `json:"distance"`
		ID           string  `json:"id"`
		ImageURL     string  `json:"image_url"`
		IsClosed     bool    `json:"is_closed"`
		Location     struct {
			Address1       string   `json:"address1"`
			Address2       string   `json:"address2"`
			Address3       string   `json:"address3"`
			City           string   `json:"city"`
			Country        string   `json:"country"`
			DisplayAddress []string `json:"display_address"`
			State          string   `json:"state"`
			ZipCode        string   `json:"zip_code"`
		} `json:"location"`

		Name         string        `json:"name"`
		Phone        string        `json:"phone"`
		Price        string        `json:"price"`
		Rating       float64       `json:"rating"`
		ReviewCount  int           `json:"review_count"`
		Transactions []interface{} `json:"transactions"`
		URL          string        `json:"url"`
	} `json:"businesses"`

	Region struct {
		Center struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"center"`
	} `json:"region"`

	Total int `json:"total"`
}

func newRestaurantCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("restaurants")
}

// AddPost insert a new Post into database and returns
// last inserted post on success.
func AddRestaurant(m Restaurant) (post Restaurant, err error) {
	c := newRestaurantCollection()
	defer c.Close()
	//m.ID = bson.NewObjectId()
	return m, c.Session.Insert(m)
}

// GetPosts Get all Post from database and returns
// list of Post on success
func GetRestaurants(term string,location string,limit string) ([]Businesses, error) {
	var (
		businesses []Businesses
		restaurant Restaurant
		err   error
	)

	/*c := newRestaurantCollection()
	defer c.Close()

	//err = c.Session.Find(nil).Sort("-createdAt").All(&restaurants)
	err = c.Session.Find(nil).All(&restaurants)*/

	// Call Yelp API to retrieve list of restaurants
	query := "?"
	if term != "" {
		query = query + "term=" + term
	}

	if location != "" {
		query = query + "&location=" + location
	}

	if limit != "" {
		query = query + "&limit=" + limit
	}

	url := "https://api.yelp.com/v3/businesses/search" + query
	fmt.Println("Request URL= " + url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer S7ADV4eUhI-xZdcRjW5zE5xHFA1YKB_Eae9Li0Gff8igf4WuizoIldEIaK6n-sKGWBf_pF27QVd-OskiFhM5ojXei36VLLNn0DCW4D01WocNlcPpVCwUnklqhlf8WHYx")
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	//data := map[string]string{}
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ReadAll: ", err)
	}

	json.Unmarshal(response, &restaurant)
	for _, business := range restaurant.Businesses {
		b := Businesses{}
		b.ID = business.ID
		b.DisplayPhone = business.DisplayPhone
		b.Distance = business.Distance
		b.ImageURL = business.ImageURL
		b.IsClosed = business.IsClosed
		b.Name = business.Name
		b.Phone = business.Phone
		b.Price = business.Price
		b.Rating = business.Rating
		b.ReviewCount = business.ReviewCount
		b.Coordinates.Latitude = business.Coordinates.Latitude
		b.Coordinates.Longitude = business.Coordinates.Longitude

		businesses = append(businesses, b)
	}

	return businesses, err
}

// GetPost Get a Post from database and returns
// a Post on success
func GetRestaurant(id bson.ObjectId) (Restaurant, error) {
	var (
		restaurant Restaurant
		err  error
	)

	c := newRestaurantCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&restaurant)
	return restaurant, err
}

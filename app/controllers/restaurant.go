package controllers

import (
	"encoding/json"
	"errors"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"nttdata/app/models"
	"fmt"
)

type Restaurant struct {
	*revel.Controller
}

func (c Restaurant) Index() revel.Result {
	/*var (
		restaurants []models.Businesses
		err   error
	)
	restaurants, err = models.GetRestaurants()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.Render(restaurants)*/

	return c.Render()
}

func (c Restaurant) Search(term string,location string,limit string) revel.Result {
	fmt.Println("Term to Search 1 " + term)
	var (
		restaurants []models.Businesses
		err   error
	)
	restaurants, err = models.GetRestaurants(term,location,limit)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200

	return c.Render(restaurants, term, location, limit)
}

func (c Restaurant) Show(id string) revel.Result {
	var (
		restaurant   models.Restaurant
		err    error
		restaurantID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid restaurant id format 1"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	restaurantID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid restaurant id format 2"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	restaurant, err = models.GetRestaurant(restaurantID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(restaurant)
}

func (c Restaurant) Create() revel.Result {
	var (
		restaurant models.Restaurant
		err  error
	)

	err = json.NewDecoder(c.Request.Body).Decode(&restaurant)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	restaurant, err = models.AddRestaurant(restaurant)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(restaurant)
}

/*func (c Restaurant) Update() revel.Result {
	var (
		restaurant models.Restaurant
		err  error
	)
	err = json.NewDecoder(c.Request.Body).Decode(&restaurant)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = restaurant.UpdateRestaurant()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.RenderJSON(restaurant)
}

func (c Restaurant) Delete(id string) revel.Result {
	var (
		err    error
		restaurant   models.Restaurant
		restaurantID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid restaurant id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	restaurantID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid restaurant id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	restaurant, err = models.GetRestaurant(restaurantID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = restaurant.DeleteRestaurant()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}*/

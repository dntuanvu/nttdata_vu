// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tRestaurant struct {}
var Restaurant tRestaurant


func (_ tRestaurant) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Restaurant.Index", args).URL
}

func (_ tRestaurant) Search(
		term string,
		location string,
		limit string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "term", term)
	revel.Unbind(args, "location", location)
	revel.Unbind(args, "limit", limit)
	return revel.MainRouter.Reverse("Restaurant.Search", args).URL
}

func (_ tRestaurant) Show(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Restaurant.Show", args).URL
}

func (_ tRestaurant) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Restaurant.Create", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


type tJobs struct {}
var Jobs tJobs


func (_ tJobs) Status(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Jobs.Status", args).URL
}



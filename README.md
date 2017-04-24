Task Assignment - GO 
===============================

The task assignment sample app demonstrates:

* Using an NoSQL (MongoDB) database and configuring the Revel DB module.
* Using the third party [mgo_revel](https://github.com/revel/mgo_revel) 
* Using [validation](../manual/validation) and displaying inline errors


	nttdata_vu/app/
		models		   # Structs and validation.
			restaurant.go
			business.go

		controllers
			init.go    # Register all of the interceptors.
			restaurant.go    # Full restaurant object received from api searching and booking
			business.go  # Business object to display in web page searching and booking

		views
			/restaurant/index.html
			/restaurant/search.html
			/header.html
			/footer.html


## revel Installation
The app uses [revel](https://gopkg.in/revel/revel) framework. 

## mongodb Installation

The app uses [mgo_v2](https://gopkg.in/mgo.v2/bson) database driver. 


### To install on OSX:

1. Install [Homebrew](http://mxcl.github.com/homebrew/) if you don't already have it.
2. Install mongodb: 

~~~
$ brew install mongodb
~~~

## Run application
- run the command "revel run nttdata" in the nttdata directory
- type "http://localhost:9000/" into browser and see the result 
	- 1st page: Search page with 2 text fields (for keying term and location to search) and 1 dropdownlist to choose maximum result will be shown
	- 2nd page: Search's result 

# nttdata_vu

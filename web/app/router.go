package app
import "./controllers"

func Routes() []collect{

	routes := []collect{
		{
			url:"/",
			controller : controllers.HomeController,
		},
		{
			url:"/api/home",
			controller : controllers.MauricioController,
		},
	}
	return routes
}

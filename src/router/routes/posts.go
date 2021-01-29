package routes

import (
	"api/src/controllers"
	"net/http"
)

var routePosts = []Route{
	{
		URI: "/posts",
		Method: http.MethodPost,
		Function: controllers.CreatePost,
		RequiresAuthentication: true,
	},
	{
		URI: "/posts",
		Method: http.MethodGet,
		Function: controllers.ListPost,
		RequiresAuthentication: true,
	},
	{
		URI: "/posts/{id}",
		Method: http.MethodGet,
		Function: controllers.FindPost,
		RequiresAuthentication: true,
	},
	{
		URI: "/posts/{id}",
		Method: http.MethodPut,
		Function: controllers.UpdatePost,
		RequiresAuthentication: true,
	},
	{
		URI: "/posts/{id}",
		Method: http.MethodDelete,
		Function: controllers.DeletePost,
		RequiresAuthentication: true,
	},
	{
		URI: "/posts/{id}/like",
		Method: http.MethodPost,
		Function: controllers.Like,
		RequiresAuthentication: true,
	},
	{
		URI: "/posts/{id}/unlike",
		Method: http.MethodPost,
		Function: controllers.UnLike,
		RequiresAuthentication: true,
	},
}

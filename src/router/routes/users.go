package routes

import (
	"api/src/controllers"
	"net/http"
)

var routeUsers = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUSer,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.ListUsers,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.FindUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{id}/addfriend",
		Method:                 http.MethodPost,
		Function:               controllers.AddFriend,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{id}/removefriend",
		Method:                 http.MethodPost,
		Function:               controllers.RemoveFriend,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{id}/password",
		Method:                 http.MethodPost,
		Function:               controllers.UpdatePassword,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{id}/friends",
		Method:                 http.MethodGet,
		Function:               controllers.ListFriends,
		RequiresAuthentication: false,
	},
}

package routes

import (
	"github.com/golangast/genserv/src/handler/get/home"
	"github.com/golangast/genserv/src/handler/get/loginemail"
	"github.com/golangast/genserv/src/handler/get/profile"
	"github.com/golangast/genserv/src/handler/post/createuser"
	"github.com/golangast/genserv/src/handler/post/userinput"
	"github.com/golangast/genserv/src/handler/restful/post"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	//get
	e.GET("/", home.Home)
	e.GET("/usercreate", profile.Profile)
	e.GET("/loginemail/:email/:sitetoken", loginemail.LoginEmail)

	//post
	e.POST("/usercreate", createuser.Createuser)
	e.POST("/userinput", userinput.UserInput)
	e.POST("/p", post.Posts)

}

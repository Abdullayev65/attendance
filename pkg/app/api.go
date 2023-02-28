package app

import "github.com/gin-gonic/gin"

func (a *App) initRouters() *gin.Engine {
	router := gin.Default()

	user := router.Group("/auth")
	{
		user.POST("/new-user", a.handler.CreateUser)
		user.POST("/sign-in", a.handler.SignIn)
	}

	department := router.Group("/department")
	{
		department.POST("/", a.handler.DepAdd)
		department.GET("/all", a.handler.DepAll)
	}

	position := router.Group("/position")
	{
		position.POST("/", a.handler.PositionAdd)
		position.GET("/all", a.handler.PositionAll)
	}

	attendance := router.Group("/attendance")
	{
		attendance.POST("/",
			a.MW.UserIDFromToken, a.handler.AttendanceAdd)
		attendance.GET("/all", a.handler.AttendanceAll)
		attendance.GET("/list", a.handler.AttendanceList)
		attendance.GET("/:userID",
			a.MW.SetIntFromParam("userID"), a.handler.AttendanceByUserID)
	}

	return router
}

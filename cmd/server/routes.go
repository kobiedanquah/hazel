package main

import (
	"github.com/freekobie/hazel/middlewares"
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// users
	router.POST("/auth/register", app.h.CreateUser)
	router.POST("/auth/login", app.h.LoginUser)
	router.POST("/auth/access", app.h.GetUserAccessToken)
	router.POST("/auth/verify", app.h.VerifyUser)
	router.POST("/auth/verify/request", app.h.RequestVerification)

	protected := router.Group("/")
	protected.Use(middlewares.Authentication())
	{
		//users
		protected.GET("/users/:id", app.h.GetUser)
		protected.PATCH("/users/profile", app.h.UpdateUserData)
		protected.DELETE("/users/:id", app.h.DeleteUser)

		// workspaces
		protected.POST("/workspaces", app.h.CreateWorkspace)
		protected.GET("/workspaces/:id", app.h.GetWorkspace)
		protected.GET("/workspaces/me", app.h.GetUserWorkspaces)
		protected.PATCH("/workspaces/:id", app.h.UpdateWorkspace)
		protected.DELETE("/workspaces/:id", app.h.DeleteWorkspace)
		protected.POST("/workspaces/:id/members", app.h.AddWorkspaceMember)
		protected.DELETE("/workspaces/:id/members/:member_id", app.h.DeleteWorkspaceMember)
		protected.GET("/workspaces/:id/projects", app.h.GetProjectsInWorkspace)

		// projects
		protected.POST("/projects", app.h.CreateProject)
		protected.GET("/projects/:id", app.h.GetProject)
		protected.PATCH("/projects/:id", app.h.UpdateProject)
		protected.DELETE("/projects/:id", app.h.DeleteProject)
		protected.GET("/projects/:id/tasks", app.h.GetProjectTasks)

		// Tasks
		protected.POST("/tasks", app.h.CreateTask)
		protected.GET("/tasks/:id", app.h.GetTask)
		protected.PATCH("/tasks/:id", app.h.UpdateTask)
		protected.DELETE("/tasks/:id", app.h.DeleteTask)
		protected.POST("/tasks/:id/assignments")
		protected.GET("/tasks/:id/assignments")
		protected.DELETE("/tasks/:id/assignments/:member_id")
	}

	return router
}

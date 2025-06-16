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

	authorized := router.Group("/")
	authorized.Use(middlewares.Authentication())
	{
		//users
		authorized.GET("/users/:id", app.h.GetUser)
		authorized.PATCH("/users/profile", app.h.UpdateUserData)
		authorized.DELETE("/users/:id", app.h.DeleteUser)

		// workspaces
		authorized.POST("/workspaces", app.h.CreateWorkspace)
		authorized.GET("/workspaces/:id", app.h.GetWorkspace)
		authorized.GET("/workspaces/me", app.h.GetUserWorkspaces)
		authorized.PATCH("/workspaces/:id", app.h.UpdateWorkspace)
		authorized.DELETE("/workspaces/:id", app.h.DeleteWorkspace)
		authorized.POST("/workspaces/:id/members", app.h.AddWorkspaceMember)
		authorized.DELETE("/workspaces/:id/members/:member_id", app.h.DeleteWorkspaceMember)
		authorized.GET("/workspaces/:id/projects", app.h.GetProjectsInWorkspace)
		// projects
		authorized.POST("/projects", app.h.CreateProject)
		authorized.GET("/projects/:id", app.h.GetProject)
		authorized.PATCH("/projects/:id", app.h.UpdateProject)
		authorized.DELETE("/projects/:id", app.h.DeleteProject)
		authorized.GET("/projects/:id/tasks")

		// Tasks

		// Comments
	}

	return router
}

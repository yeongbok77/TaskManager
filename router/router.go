// Package router 包设置了路由
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yeongbok77/TaskManager/controller"
	"github.com/yeongbok77/TaskManager/logger"
)

// SetUpRouter 函数设置了路由
func SetUpRouter() (r *gin.Engine) {
	r = gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/issue")
	{
		v1.GET("/list", controller.ListIssueHandler)
		v1.GET("/action", controller.ActionIssueHandler)
		v1.GET("/addMilestone", controller.AddMilestoneHandler)
		v1.GET("/addTag", controller.AddTagHandler)
		v1.POST("/addComment", controller.AddCommentHandler)

		v1.GET("/milestone/action", controller.ActionMilestoneHandler)
	}

	return r
}

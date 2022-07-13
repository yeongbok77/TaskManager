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

	// issue 路由组
	v1 := r.Group("/issue")
	{
		v1.GET("/action", controller.ActionIssueHandler)
		v1.GET("/applyMilestone", controller.ApplyMilestoneHandler)
		v1.GET("/applyTag", controller.ApplyTagHandler)
		v1.POST("/addComment", controller.AddCommentHandler)
		v1.GET("/list", controller.ListIssueHandler)
		v1.GET("/listIssueTagFilter", controller.ListIssueTagFilterHandler)
		v1.GET("/listBasisMilestone", controller.ListBasisMilestoneHandler)
	}

	// tag 路由组
	v2 := r.Group("/tag")
	{
		v2.GET("/action", controller.ActionTagHandler)
	}

	// milestone 路由组
	v3 := r.Group("/milestone")
	{
		v3.GET("/action", controller.ActionMilestoneHandler)
	}

	return r
}

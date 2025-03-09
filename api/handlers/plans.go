package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlanFinder interface {
	FetchPlan(*gin.Context) (string, error)
}

func FetchPlan(plans PlanFinder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		

		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Plan Handler Hit!"})
	}
}


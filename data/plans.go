package data

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type Plan struct {
}

type PlanRepository struct {
	// queries *queries.Queries
}

func NewPlanRepository(db *pgx.Conn) *PlanRepository {
	// queries := queries.New(db)

	return &PlanRepository{
		// queries: 	queries,
	}
}

func (r *PlanRepository) FetchPlan(ctx *gin.Context) (string, error) {

	return "Plan Data File Hit!", nil
}
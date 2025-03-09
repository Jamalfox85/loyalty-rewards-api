package api

import (
	"github.com/jamalfox85/loyalty-rewards-api/data"
)

type Application struct {
	Plans *data.PlanRepository
}
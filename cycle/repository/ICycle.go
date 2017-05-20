package repository

import (
	"radiup/cycle"
)

type CycleManager interface {
	CreateCycle(c Cycle)
	UpdateCycle(id int)
	RemoveCycle(id int)
	SearchCycle(id int)
}	
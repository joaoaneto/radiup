package repository

import (
	"radiup/cycle"
)

type CycleManager interface {
	CreateCycle(c Cycle)
	UpdateCycle(/*what to search for cycles?*/)
	RemoveCycle(/*what to search?*/)
	SearchCycle(/*what to search for cycles?*/)
}
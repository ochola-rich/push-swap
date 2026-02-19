// Package models
package models

type MovePlan struct {
	Shared       int
	RemainingA   int
	RemainingB   int
	MoveA, MoveB string
	SharedType   string
	TotalCost    int
}

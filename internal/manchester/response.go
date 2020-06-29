package manchester

import "time"

type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type tireChangeTimeResponse struct {
	ID        uint      `json:"id"`
	Time      time.Time `json:"time"`
	Available bool      `json:"available"`
}

func newTireChangeTimeResponse(entity *tireChangeTimeEntity) *tireChangeTimeResponse {
	return &tireChangeTimeResponse{
		ID:        entity.ID,
		Time:      entity.Time.UTC(),
		Available: entity.Available,
	}
}

type tireChangeTimesResponse []*tireChangeTimeResponse

func newTireChangeTimesResponse(entities []*tireChangeTimeEntity) *tireChangeTimesResponse {
	var availableTimes []*tireChangeTimeResponse

	for _, entity := range entities {
		availableTimes = append(availableTimes, newTireChangeTimeResponse(entity))
	}

	response := tireChangeTimesResponse(availableTimes)

	return &response
}

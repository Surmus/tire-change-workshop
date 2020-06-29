package london

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type tireChangeTimesService struct {
	repository *tireChangeTimeRepository
}

func newTireChangeTimesService(repository *tireChangeTimeRepository) *tireChangeTimesService {
	return &tireChangeTimesService{repository: repository}
}

func (s *tireChangeTimesService) getAvailable(from time.Time, until time.Time) *tireChangeTimesResponse {
	log.Infof("fetching tire change times from %s until %s", from, until)

	tireChangeTimes := s.repository.availableByTimeRange(from, until)

	log.Infof("successfully fetched %d tire change times from %s until %s", len(tireChangeTimes), from, until)

	return newTireChangeTimesResponse(tireChangeTimes)
}

func (s *tireChangeTimesService) book(uuid string, contactInformation string) *tireChangeTimeResponse {
	log.Infof("trying to book tire change time with uuid: %s", uuid)
	tireChangeTime := s.repository.availableByUUID(uuid)
	tireChangeTime.makeBooking(contactInformation)
	tireChangeTime = s.repository.save(tireChangeTime)

	log.Infof("successfully booked tire change time with uuid: %s", uuid)
	return newTireChangeTimeResponse(tireChangeTime.UUID, tireChangeTime.Time)
}

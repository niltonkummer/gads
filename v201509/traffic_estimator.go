package v201509

type TrafficEstimatorService struct {
	Auth
}

func NewTrafficEstimatorService(auth *Auth) *TrafficEstimatorService {
	return &TrafficEstimatorService{Auth: *auth}
}

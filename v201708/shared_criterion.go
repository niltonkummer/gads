package v201708

type SharedCriterionService struct {
	Auth
}

func NewSharedCriterionService(auth *Auth) *SharedCriterionService {
	return &SharedCriterionService{Auth: *auth}
}
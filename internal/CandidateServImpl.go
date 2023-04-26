package internal

import "github.com/dilly3/blood-donor/internal/models"

type ServImpl struct {
	DB IDatabase
}

func NewServImpl(db IDatabase) *ServImpl {
	return &ServImpl{
		DB: db,
	}

}
func (impl *ServImpl) GetById(id string) (*models.Candidate, error) {

	cand, err := impl.DB.GetById(id)
	if err != nil {
		return nil, err
	}
	return cand, nil
}

func (impl *ServImpl) GetAllCandidates() ([]*models.Candidate, error) {

	candidates, err := impl.DB.GetAllCandidates()
	if err != nil {
		return nil, err
	}
	return candidates, nil
}
func (impl *ServImpl) SaveCandidate(cand models.Candidate) (bool, error) {
	cand1 := models.Candidate{}
	cand1.Address = cand.Address
	cand1.Id = models.GetshortId()
	cand1.BloodGroup = cand.BloodGroup
	cand1.Age = cand.Age
	cand1.Email = cand.Email
	cand1.FullName = cand.FullName
	cand1.Mobile = cand.Mobile
	b, err := impl.DB.SaveCandidate(cand1)
	if err != nil {
		return false, err
	}
	return b, nil
}
func (impl *ServImpl) GetByFullname(name string) (*models.Candidate, error) {
	cand, err := impl.DB.GetByFullname(name)
	if err != nil {
		return nil, err
	}
	return cand, nil

}

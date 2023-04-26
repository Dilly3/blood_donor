package internal

import "github.com/dilly3/blood-donor/internal/models"

type ICandidateServ interface {
	GetById(id string) (*models.Candidate, error)
	GetAllCandidates() ([]*models.Candidate, error)
	SaveCandidate(cand models.Candidate) (bool, error)
	GetByFullname(name string) (*models.Candidate, error)
}

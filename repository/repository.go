package repository

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/jgautheron/exago/repository/model"
	"github.com/jgautheron/exago/score"
)

var (
	// Make sure it satisfies the interface.
	_ model.Record = (*Repository)(nil)
)

type Repository struct {
	Name          string     `json:"name"`
	Branch        string     `json:"branch"`
	ExecutionTime string     `json:"execution_time"`
	LastUpdate    time.Time  `json:"last_update"`
	Data          model.Data `json:"results"`
}

func New(name, branch string) *Repository {
	return &Repository{
		Name:   name,
		Branch: branch,
	}
}

// ApplyScore calculates the score based on the repository results.
func (r *Repository) ApplyScore() (err error) {
	val, res := score.Process(r.Data)
	r.Data.Score.Value = val
	r.Data.Score.Details = res
	r.Data.Score.Rank = score.Rank(r.Data.Score.Value)

	log.Infof(
		"[%s] Rank: %s, overall score: %.2f",
		r.GetName(),
		r.Data.Score.Rank,
		r.Data.Score.Value,
	)

	return nil
}

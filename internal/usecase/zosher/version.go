package zosher

import (
	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/logging"
	"github.com/s-yakubovskiy/whereami/internal/metrics"
)

// var _ LocationKeeperRepo = &db.LocationKeeper{}

type UseCase struct {
	cfg *config.AppConfig
	log logging.Logger
	m   metrics.Metrics
	z   ZoshRepo
}

type ZoshVersion struct {
	Version string
	Commit  string
	Date    string
}

type ZoshRepo interface {
	GetVersion() string
	GetCommitHash() string
	GetDate() string
}

func NewZoshUseCase(log logging.Logger, cfg *config.AppConfig, m metrics.Metrics, z ZoshRepo) *UseCase {
	// Register metrics specific to this use case
	// m.RegisterCounter("task_count", "Counts custom task executions", []string{"status"})
	// m.RegisterHistogram("task_latency", "Tracks custom task latencies", []string{"task_type"})
	return &UseCase{
		cfg: cfg,
		log: log,
		m:   m,
		z:   z,
	}
}

func (c *UseCase) Version() *ZoshVersion {
	return &ZoshVersion{
		Version: c.z.GetVersion(),
		Commit:  c.z.GetCommitHash(),
		Date:    c.z.GetDate(),
	}
}

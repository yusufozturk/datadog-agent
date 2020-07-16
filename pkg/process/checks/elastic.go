package checks

import (
	"github.com/pkg/errors"
	"sync"
	"time"

	model "github.com/DataDog/agent-payload/process"
	"github.com/DataDog/datadog-agent/pkg/process/config"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)


// Elasticsearch is a singleton ElasticCheck.
var Elasticsearch = &ElasticCheck{}

// ElasticCheck collects shard + cluster information
type ElasticCheck struct {
	sync.Mutex

	sysInfo         *model.SystemInfo
	lastRun         time.Time
}

// Init initializes the singleton ElasticCheck.
func (e *ElasticCheck) Init(_ *config.AgentConfig, info *model.SystemInfo) {
	e.sysInfo = info
}

// Name returns the name of the ElasticCheck.
func (e *ElasticCheck) Name() string { return "elastic" }

// RealTime indicates if this check only runs in real-time mode.
func (e *ElasticCheck) RealTime() bool { return false }

func (e *ElasticCheck) Run(cfg *config.AgentConfig, groupID int32) ([]model.MessageBody, error) {
	e.Lock()
	defer e.Unlock()

	e.lastRun = time.Now()
	time.Sleep(1*time.Second)

	log.Infof("Collected %n shards in %v", 0, time.Now().Sub(e.lastRun))

	return nil, errors.New("unimplemented")
}

package checks

import (
	"errors"
	"sync"
	"time"

	model "github.com/DataDog/agent-payload/process"
	"github.com/DataDog/datadog-agent/pkg/process/config"
	"github.com/DataDog/datadog-agent/pkg/util/log"

	elasticsearch6 "github.com/elastic/go-elasticsearch/v6"
)


// Elasticsearch is a singleton ElasticCheck.
var Elasticsearch = &ElasticCheck{}

// ElasticCheck collects shard + cluster information
type ElasticCheck struct {
	sync.Mutex

	sysInfo *model.SystemInfo
	lastRun time.Time

	clusterName string

	es *elasticsearch6.Client
}

// Init initializes the singleton ElasticCheck.
func (e *ElasticCheck) Init(_ *config.AgentConfig, info *model.SystemInfo) {
	e.sysInfo = info

	// TODO: Does this work with ES5 if we're only accessing _cat endpoints?
	client, err := elasticsearch6.NewDefaultClient() // Accesses localhost:9200
	if err != nil {
		_ = log.Errorf("failed to create elasticsearch client: %s", err)
		return
	}

	log.Infof("elasticsearch client version: %s", elasticsearch6.Version)
	esInfo, err := client.Info()
	if err != nil {
		_ = log.Errorf("failed to get elasticsearch info: %s", err)
		return
	}


	log.Infof("ES Info: %+v", esInfo)
}

// Name returns the name of the ElasticCheck.
func (e *ElasticCheck) Name() string { return "elastic" }

// RealTime indicates if this check only runs in real-time mode.
func (e *ElasticCheck) RealTime() bool { return false }

func (e *ElasticCheck) Run(cfg *config.AgentConfig, groupID int32) ([]model.MessageBody, error) {
	e.Lock()
	defer e.Unlock()

	// Don't run the check if the client failed to connect
	if e.es == nil {
		return nil, errors.New("no elasticsearch client configured")
	}

	e.lastRun = time.Now()
	time.Sleep(1*time.Second)

	log.Infof("Collected %d shards in %v", 0, time.Now().Sub(e.lastRun))

	return nil, nil
}

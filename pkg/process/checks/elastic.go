package checks

import (
	"bytes"
	"errors"
	"io"
	"sync"
	"time"

	model "github.com/DataDog/agent-payload/process"
	"github.com/DataDog/datadog-agent/pkg/process/config"
	"github.com/DataDog/datadog-agent/pkg/util/log"

	elasticsearch6 "github.com/elastic/go-elasticsearch/v6"
	"github.com/tidwall/gjson"
)


// Elasticsearch is a singleton ElasticCheck.
var Elasticsearch = &ElasticCheck{}

// ElasticCheck collects shard + cluster information
type ElasticCheck struct {
	sync.Mutex

	sysInfo *model.SystemInfo
	lastRun time.Time

	clusterName string
	isLeaderNode bool

	es *elasticsearch6.Client
}

// Init initializes the singleton ElasticCheck.
func (e *ElasticCheck) Init(_ *config.AgentConfig, info *model.SystemInfo) {
	e.sysInfo = info

	log.Infof("elasticsearch client version: %s", elasticsearch6.Version)

	// TODO: Does this work with ES5 if we're only accessing _cat endpoints?
	client, err := elasticsearch6.NewDefaultClient() // Accesses localhost:9200
	if err != nil {
		_ = log.Errorf("failed to create elasticsearch client: %s", err)
		return
	}
	e.es = client

	if e.clusterName, err = e.getClusterName(); err == nil {
		log.Infof("elasticsearch cluster: %s", e.clusterName)
	}
}

// Not my fault: https://github.com/elastic/go-elasticsearch/blob/master/_examples/encoding/gjson.go
func readJsonBlob(r io.Reader) []byte {
	var b bytes.Buffer
	_, _ = b.ReadFrom(r) // TODO: Ayy
	return b.Bytes()
}

func (e *ElasticCheck) getClusterName() (string, error) {
	esInfo, err := e.es.Info()
	if err != nil {
		return "", log.Errorf("failed to get elasticsearch info: %s", err)
	}

	defer esInfo.Body.Close()

	log.Infof("ES Info: %+v", esInfo)

	jsonBlob := readJsonBlob(esInfo.Body)
	cluster := gjson.GetBytes(jsonBlob, "cluster_name").String();
	if cluster == "" {
		return "", log.Errorf("unable to find elasticsearch cluster name")
	}

	return cluster, nil
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

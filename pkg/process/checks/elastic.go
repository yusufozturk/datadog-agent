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
	"github.com/elastic/go-elasticsearch/v6/esapi"
	"github.com/tidwall/gjson"
)


// Elasticsearch is a singleton ElasticCheck.
var Elasticsearch = &ElasticCheck{}

// ElasticCheck collects shard + cluster information
type ElasticCheck struct {
	sync.Mutex

	sysInfo *model.SystemInfo
	lastRun time.Time

	nodeName string
	clusterName string
	clusterUUID string

	isLeaderNode bool

	es *elasticsearch6.Client
}

// Init initializes the singleton ElasticCheck.
func (e *ElasticCheck) Init(_ *config.AgentConfig, info *model.SystemInfo) {
	e.sysInfo = info

	log.Infof("elasticsearch client version: %s", elasticsearch6.Version)

	// TODO: Does this work with ES5 if we're only accessing _cat endpoints?
	// TODO: We should wrap this up in a retrier so that we try and handle intermittent cluster communication failures
	client, err := elasticsearch6.NewDefaultClient() // Accesses localhost:9200
	if err != nil {
		_ = log.Errorf("failed to create elasticsearch client: %s", err)
		return
	}
	e.es = client

	e.getClusterInfo()

	e.isLeaderNode = e.isLeader()
	if e.isLeaderNode {
		log.Infof("this node (%s) is considered the ES leader, tracking shards...", e.nodeName)
	}
}

// Not my fault: https://github.com/elastic/go-elasticsearch/blob/master/_examples/encoding/gjson.go
func readJsonBlob(r io.Reader) []byte {
	var b bytes.Buffer
	_, _ = b.ReadFrom(r) // TODO: Ayy
	return b.Bytes()
}

func (e *ElasticCheck) getClusterInfo()  {
	esInfo, err := e.es.Info()
	if err != nil {
		log.Errorf("failed to get elasticsearch info: %s", err)
		return
	}

	defer esInfo.Body.Close()
	jsonBlob := readJsonBlob(esInfo.Body)
	log.Infof("ES Info: %+v", esInfo)

	// Cluster name
	e.clusterName = gjson.GetBytes(jsonBlob, "cluster_name").String();
	if e.clusterName == "" {
		e.clusterName = "unknown"
		log.Warnf("unable to find elasticsearch cluster name")
	}

	// Cluster UUID
	e.clusterUUID = gjson.GetBytes(jsonBlob, "cluster_uuid").String();
	if e.clusterUUID == "" {
		e.clusterUUID = "unknown"
		log.Warnf("unable to find elasticsearch cluster UUID")
	}

	// Node name
	e.nodeName = gjson.GetBytes(jsonBlob, "name").String();
	if e.nodeName == "" {
		e.nodeName = "unknown"
		log.Warnf("unable to find elasticsearch node name")
	}

	log.Infof("elasticsearch cluster: %s (%s), node: %s", e.clusterName, e.clusterUUID, e.nodeName)
}

// Note: this doesn't really deal with split brains and multiple nodes thinking they're leaders... \o/
func (e *ElasticCheck) isLeader() bool {
	leaderInfo, err := e.es.Cat.Master(func(request *esapi.CatMasterRequest) {
		request.Format = "json"
	})

	if err != nil {
		log.Warnf("failed to get elasticsearch leader info: %s", err)
		return false
	}
	defer leaderInfo.Body.Close()

	jsonBlob := readJsonBlob(leaderInfo.Body)
	leaderNode := gjson.GetBytes(jsonBlob, "node").String();
	if leaderNode == "" {
		log.Warnf("unable to find elasticsearch leader, defaulting to false")
		return false
	}

	return leaderNode == e.nodeName
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

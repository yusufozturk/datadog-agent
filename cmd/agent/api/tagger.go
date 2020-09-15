package api

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/DataDog/datadog-agent/cmd/agent/api/pb"
	"github.com/DataDog/datadog-agent/pkg/tagger"
	"github.com/DataDog/datadog-agent/pkg/tagger/collectors"
)

type taggerServer struct {
	pb.TaggerServer
}

func (s *taggerServer) StreamTags(req *pb.StreamTagsRequest, stream pb.Tagger_StreamTagsServer) error {
	cardinality, err := pb2taggerCardinality(req.Cardinality)
	if err != nil {
		// TODO(juliogreff): wrap errors with gRPC status
		return err
	}

	// TODO(juliogreff): implement filtering

	entities := tagger.List(cardinality)
	for id, entity := range entities.Entities {
		entityID, err := tagger2pbEntityID(id)
		if err != nil {
			// TODO(juliogreff): log and continue
			continue
		}

		err = stream.Send(&pb.StreamTagsResponse{
			Type: pb.EventType_ADDED,
			Entity: &pb.Entity{
				Id:   entityID,
				Tags: entity.Tags,
			},
		})
		if err != nil {
			// TODO(juliogreff): wrap errors with gRPC status
			return err
		}
	}

	return nil
}
func (s *taggerServer) FetchEntity(ctx context.Context, req *pb.FetchEntityRequest) (*pb.Entity, error) {
	entityID := fmt.Sprintf("%s://%s", req.Id.Prefix, req.Id.Uid)
	cardinality, err := pb2taggerCardinality(req.Cardinality)
	if err != nil {
		// TODO(juliogreff): wrap errors with gRPC status
		return nil, err
	}

	tags, err := tagger.Tag(entityID, cardinality)
	if err != nil {
		// TODO(juliogreff): wrap errors with gRPC status
		return nil, err
	}

	return &pb.Entity{
		Id:   req.Id,
		Tags: tags,
	}, nil
}

func tagger2pbEntityID(entityID string) (*pb.EntityId, error) {
	parts := strings.SplitN(entityID, "://", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid entity id %q", entityID)
	}

	return &pb.EntityId{
		Prefix: parts[0],
		Uid:    parts[1],
	}, nil
}

func pb2taggerCardinality(pbCardinality pb.TagCardinality) (collectors.TagCardinality, error) {
	switch pbCardinality {
	case pb.TagCardinality_LOW:
		return collectors.LowCardinality, nil
	case pb.TagCardinality_ORCHESTRATOR:
		return collectors.OrchestratorCardinality, nil
	case pb.TagCardinality_HIGH:
		return collectors.HighCardinality, nil
	}

	return 0, fmt.Errorf("invalid cardinality %q", pbCardinality)
}

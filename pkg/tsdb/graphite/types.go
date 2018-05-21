package graphite

import "github.com/iyeonok/grafana/pkg/tsdb"

type TargetResponseDTO struct {
	Target     string                `json:"target"`
	DataPoints tsdb.TimeSeriesPoints `json:"datapoints"`
}

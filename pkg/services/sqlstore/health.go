package sqlstore

import (
	"github.com/iyeonok/grafana/pkg/bus"
	m "github.com/iyeonok/grafana/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetDBHealthQuery)
}

func GetDBHealthQuery(query *m.GetDBHealthQuery) error {
	return x.Ping()
}

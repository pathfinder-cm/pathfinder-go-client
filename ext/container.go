package ext

import (
	"encoding/json"

	"github.com/pathfinder-cm/pathfinder-go-client/pfmodel"
)

type ContainerRes struct {
	ApiVersion string            `json:"api_version"`
	Data       pfmodel.Container `json:"data"`
}

func NewContainerFromByte(b []byte) (*pfmodel.Container, error) {
	var res ContainerRes
	err := json.Unmarshal(b, &res)

	if err != nil {
		return nil, err
	}

	container := pfmodel.Container{
		Hostname:      res.Data.Hostname,
		Ipaddress:     res.Data.Ipaddress,
		NodeHostname:  res.Data.NodeHostname,
		Status:        res.Data.Status,
		Bootstrappers: res.Data.Bootstrappers,
		Source: pfmodel.Source{
			Type:  res.Data.Source.Type,
			Mode:  res.Data.Source.Mode,
			Alias: res.Data.Source.Alias,
			Remote: pfmodel.Remote{
				Server:      res.Data.Source.Remote.Server,
				Protocol:    res.Data.Source.Remote.Protocol,
				AuthType:    res.Data.Source.Remote.AuthType,
				Certificate: res.Data.Source.Remote.Certificate,
			},
		},
	}
	return &container, nil
}

package ext

import (
	"encoding/json"

	"github.com/pathfinder-cm/pathfinder-go-client/pfmodel"
)

type ContainerListRes struct {
	ApiVersion string               `json:"api_version"`
	Data       ContainerListDataRes `json:"data"`
}

type ContainerListDataRes struct {
	Items []pfmodel.Container `json:"items"`
}

func NewContainerListFromByte(b []byte) (*pfmodel.ContainerList, error) {
	var res ContainerListRes
	err := json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	containers := make(pfmodel.ContainerList, len(res.Data.Items))
	for i, n := range res.Data.Items {
		containers[i] = pfmodel.Container{
			Hostname:     n.Hostname,
			Ipaddress:    n.Ipaddress,
			NodeHostname: n.NodeHostname,
			Status:       n.Status,
			Source: pfmodel.Source{
				Type:        n.Source.Type,
				Mode:        n.Source.Mode,
				Alias:       n.Source.Alias,
				Certificate: n.Source.Certificate,
				Remote: pfmodel.Remote{
					Server:   n.Source.Remote.Server,
					Protocol: n.Source.Remote.Protocol,
					AuthType: n.Source.Remote.AuthType,
				},
			},
		}
	}

	return &containers, nil
}

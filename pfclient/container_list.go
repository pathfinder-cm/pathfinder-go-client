package pfclient

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

	cl := make(pfmodel.ContainerList, len(res.Data.Items))
	for i, c := range res.Data.Items {
		cl[i] = pfmodel.Container{
			Hostname:     c.Hostname,
			Ipaddress:    c.Ipaddress,
			NodeHostname: c.NodeHostname,
			Status:       c.Status,
			Source: pfmodel.Source{
				Type:        c.Source.Type,
				Mode:        c.Source.Mode,
				Alias:       c.Source.Alias,
				Certificate: c.Source.Certificate,
				Remote: pfmodel.Remote{
					Server:   c.Source.Remote.Server,
					Protocol: c.Source.Remote.Protocol,
					AuthType: c.Source.Remote.AuthType,
				},
			},
		}
	}

	return &cl, nil
}

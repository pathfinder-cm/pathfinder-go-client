package ext

import (
	"encoding/json"

	"github.com/pathfinder-cm/pathfinder-go-client/pfmodel"
)

type ContainerRes struct {
	ApiVersion string           `json:"api_version"`
	Data       ContainerDataRes `json:"data"`
}

type ContainerDataRes struct {
	Hostname     string `json:"hostname"`
	Ipaddress    string `json:"ipaddress"`
	Image        string `json:"image"`
	NodeHostname string `json:"node_hostname"`
	Status       string `json:"status"`
}

func NewContainerFromByte(b []byte) (*pfmodel.Container, error) {
	var res ContainerRes
	err := json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	container := pfmodel.Container{
		Hostname:     res.Data.Hostname,
		Ipaddress:    res.Data.Ipaddress,
		Image:        res.Data.Image,
		NodeHostname: res.Data.NodeHostname,
		Status:       res.Data.Status,
	}

	return &container, nil
}

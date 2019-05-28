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
	Hostname      string `json:"hostname"`
	Ipaddress     string `json:"ipaddress"`
	ImageAlias    string `json:"image_alias"`
	ImageServer   string `json:"image_server"`
	ImageProtocol string `json:"image_protocol"`
	NodeHostname  string `json:"node_hostname"`
	Status        string `json:"status"`
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
		ImageAlias:    res.Data.ImageAlias,
		ImageServer:   res.Data.ImageServer,
		ImageProtocol: res.Data.ImageProtocol,
		NodeHostname:  res.Data.NodeHostname,
		Status:        res.Data.Status,
	}

	return &container, nil
}

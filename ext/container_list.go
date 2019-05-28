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
	Items []ContainerListItemRes `json:"items"`
}

type ContainerListItemRes struct {
	Hostname      string `json:"hostname"`
	Ipaddress     string `json:"ipaddress"`
	ImageAlias    string `json:"image_alias"`
	ImageServer   string `json:"image_server"`
	ImageProtocol string `json:"image_protocol"`
	NodeHostname  string `json:"node_hostname"`
	Status        string `json:"status"`
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
			Hostname:      n.Hostname,
			Ipaddress:     n.Ipaddress,
			ImageAlias:    n.ImageAlias,
			ImageServer:   n.ImageServer,
			ImageProtocol: n.ImageProtocol,
			NodeHostname:  n.NodeHostname,
			Status:        n.Status,
		}
	}

	return &containers, nil
}

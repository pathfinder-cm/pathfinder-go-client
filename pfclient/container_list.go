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
	Items []ContainerRes `json:"items"`
}

type ContainerRes struct {
	Hostname      string `json:"hostname"`
	ImageAlias    string `json:"image_alias"`
	ImageServer   string `json:"image_server"`
	ImageProtocol string `json:"image_protocol"`
	Status        string `json:"status"`
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
			Hostname:      c.Hostname,
			ImageAlias:    c.ImageAlias,
			ImageServer:   c.ImageServer,
			ImageProtocol: c.ImageProtocol,
			Status:        c.Status,
		}
	}

	return &cl, nil
}

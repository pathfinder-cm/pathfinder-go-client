package ext

import (
	"encoding/json"

	"github.com/pathfinder-cm/pathfinder-go-client/pfmodel"
)

type NodeRes struct {
	ApiVersion string      `json:"api_version"`
	Data       NodeDataRes `json:"data"`
}

type NodeDataRes struct {
	Hostname   string `json:"hostname"`
	Ipaddress  string `json:"ipaddress"`
	CreatedAt  string `json:"created_at"`
	MemFreeMb  uint64 `json:"mem_free_mb"`
	MemUsedMb  uint64 `json:"mem_used_mb"`
	MemTotalMb uint64 `json:"mem_total_mb"`
}

func NewNodeFromByte(b []byte) (*pfmodel.Node, error) {
	var res NodeRes
	err := json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	node := pfmodel.Node{
		Hostname:   res.Data.Hostname,
		Ipaddress:  res.Data.Ipaddress,
		CreatedAt:  res.Data.CreatedAt,
		MemFreeMb:  res.Data.MemFreeMb,
		MemUsedMb:  res.Data.MemUsedMb,
		MemTotalMb: res.Data.MemTotalMb,
	}

	return &node, nil
}

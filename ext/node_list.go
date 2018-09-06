package ext

import (
	"encoding/json"

	"github.com/pathfinder-cm/pathfinder-go-client/pfmodel"
)

type NodeListRes struct {
	ApiVersion string          `json:"api_version"`
	Data       NodeListDataRes `json:"data"`
}

type NodeListDataRes struct {
	Items []NodeListItemRes `json:"items"`
}

type NodeListItemRes struct {
	Hostname   string `json:"hostname"`
	Ipaddress  string `json:"ipaddress"`
	CreatedAt  string `json:"created_at"`
	MemFreeMb  uint64 `json:"mem_free_mb"`
	MemUsedMb  uint64 `json:"mem_used_mb"`
	MemTotalMb uint64 `json:"mem_total_mb"`
}

func NewNodeListFromByte(b []byte) (*pfmodel.NodeList, error) {
	var res NodeListRes
	err := json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	nodes := make(pfmodel.NodeList, len(res.Data.Items))
	for i, n := range res.Data.Items {
		nodes[i] = pfmodel.Node{
			Hostname:   n.Hostname,
			Ipaddress:  n.Ipaddress,
			CreatedAt:  n.CreatedAt,
			MemFreeMb:  n.MemFreeMb,
			MemUsedMb:  n.MemUsedMb,
			MemTotalMb: n.MemTotalMb,
		}
	}

	return &nodes, nil
}

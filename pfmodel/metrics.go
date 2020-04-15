package pfmodel

type Metrics struct {
	Memory *Memory `json:"memory"`
	Load   *Load   `json:"load"`

	RootDisk *Disk `json:"disk_root"`
	ZFSDisk  *Disk `json:"disk_zfs"`
}

type Memory struct {
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
	Total uint64 `json:"total"`
}

type Load struct {
	Capacity   int     `json:"capacity"`
	LoadAvg1M  float64 `json:"load_avg_1m"`
	LoadAvg5M  float64 `json:"load_avg_5m"`
	LoadAvg15M float64 `json:"load_avg_15m"`
}

type Disk struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
}

package pfmodel

type Container struct {
	Hostname     string  `json:"hostname"`
	Ipaddress    string  `json:"ipaddress"`
	NodeHostname string  `json:"node_hostname"`
	Status       string  `json:"status"`
	Source       Source  `json:"source"`
}

type Source struct {
	Type         string  `json:"source_type"`
	Alias        string  `json:"alias"`
	Certificate  string  `json:"certificate"`
	Mode         string  `json:"mode"`
	Remote       Remote  `json:"remote"`
}

type Remote struct {
	Server       string `json:"server"`
	Protocol     string `json:"protocol"`
	AuthType     string `json:"auth_type"`
}


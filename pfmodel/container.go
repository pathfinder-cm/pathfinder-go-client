package pfmodel

type Container struct {
	Hostname      string `json:"hostname"`
	Ipaddress     string `json:"ipaddress"`
	NodeHostname  string `json:"node_hostname"`
	Status        string `json:"status"`
	Source        Source `json:"source"`
	Bootstrappers []Bootstrapper
}

type Source struct {
	Type   string `json:"source_type"`
	Alias  string `json:"alias"`
	Mode   string `json:"mode"`
	Remote Remote `json:"remote"`
}

type Remote struct {
	Server      string `json:"server"`
	Protocol    string `json:"protocol"`
	AuthType    string `json:"auth_type"`
	Certificate string `json:"certificate"`
}

type Bootstrapper struct {
	Type         string      `json:"bootstrap_type"`
	CookbooksUrl string      `json:"bootstrap_cookbooks_url"`
	Attributes   interface{} `json:"bootstrap_attributes"`
}

package ext

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pathfinder-cm/pathfinder-go-client/pfmodel"
)

func TestGetNodes(t *testing.T) {
	tables := []struct {
		hostname   string
		ipaddress  string
		createdAt  string
		memFreeMb  uint64
		memUsedMb  uint64
		memTotalMb uint64
	}{
		{"test-01", "192.168.1.100", "2018-01-01", 100, 100, 200},
		{"test-02", "192.168.1.101", "2018-01-01", 100, 100, 200},
		{"test-03", "192.168.1.102", "2018-01-01", 100, 100, 200},
	}

	b := []byte(`{
		"api_version": "1.0",
		"data": {
			"items": [
				{"hostname": "test-01", "ipaddress": "192.168.1.100", "created_at": "2018-01-01", "mem_free_mb": 100, "mem_used_mb": 100, "mem_total_mb": 200},
				{"hostname": "test-02", "ipaddress": "192.168.1.101", "created_at": "2018-01-01", "mem_free_mb": 100, "mem_used_mb": 100, "mem_total_mb": 200},
				{"hostname": "test-03", "ipaddress": "192.168.1.102", "created_at": "2018-01-01", "mem_free_mb": 100, "mem_used_mb": 100, "mem_total_mb": 200}
			]
		}
	}`)

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write(b)
	}))
	defer func() { testServer.Close() }()

	client := NewClient("default", "", &http.Client{}, testServer.URL, map[string]string{})
	nodes, _ := client.GetNodes()
	for i, table := range tables {
		if (*nodes)[i].Hostname != table.hostname {
			t.Errorf("Incorrect node hostname generated, got: %s, want: %s.",
				(*nodes)[i].Hostname,
				table.hostname)
		}

		if (*nodes)[i].Ipaddress != table.ipaddress {
			t.Errorf("Incorrect node ipaddress generated, got: %s, want: %s.",
				(*nodes)[i].Ipaddress,
				table.ipaddress)
		}

		if (*nodes)[i].CreatedAt != table.createdAt {
			t.Errorf("Incorrect node CreatedAt generated, got: %s, want: %s.",
				(*nodes)[i].CreatedAt,
				table.createdAt)
		}

		if (*nodes)[i].MemFreeMb != table.memFreeMb {
			t.Errorf("Incorrect node MemFreeMb generated, got: %d, want: %d.",
				(*nodes)[i].MemFreeMb,
				table.memFreeMb)
		}

		if (*nodes)[i].MemUsedMb != table.memUsedMb {
			t.Errorf("Incorrect node MemUsedMb generated, got: %d, want: %d.",
				(*nodes)[i].MemUsedMb,
				table.memUsedMb)
		}

		if (*nodes)[i].MemTotalMb != table.memTotalMb {
			t.Errorf("Incorrect node MemTotalMb generated, got: %d, want: %d.",
				(*nodes)[i].MemTotalMb,
				table.memTotalMb)
		}
	}
}

func TestGetNode(t *testing.T) {
	tables := []struct {
		hostname   string
		ipaddress  string
		createdAt  string
		memFreeMb  uint64
		memUsedMb  uint64
		memTotalMb uint64
	}{
		{"test-01", "192.168.1.100", "2018-01-01", 100, 100, 200},
	}

	b := []byte(`{
		"api_version": "1.0",
		"data": {
			"id": 1,
			"hostname": "test-01",
			"ipaddress": "192.168.1.100",
			"created_at": "2018-01-01",
			"mem_free_mb": 100,
			"mem_used_mb": 100,
			"mem_total_mb": 200
		}
	}`)

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write(b)
	}))
	defer func() { testServer.Close() }()

	client := NewClient("default", "", &http.Client{}, testServer.URL, map[string]string{})
	node, _ := client.GetNode(tables[0].hostname)
	if node.Hostname != tables[0].hostname {
		t.Errorf("Incorrect node hostname generated, got: %s, want: %s.",
			node.Hostname,
			tables[0].hostname)
	}

	if node.Ipaddress != tables[0].ipaddress {
		t.Errorf("Incorrect node ipaddress generated, got: %s, want: %s.",
			node.Ipaddress,
			tables[0].ipaddress)
	}

	if node.CreatedAt != tables[0].createdAt {
		t.Errorf("Incorrect node CreatedAt generated, got: %s, want: %s.",
			node.CreatedAt,
			tables[0].createdAt)
	}

	if node.MemFreeMb != tables[0].memFreeMb {
		t.Errorf("Incorrect node MemFreeMb generated, got: %d, want: %d.",
			node.MemFreeMb,
			tables[0].memFreeMb)
	}

	if node.MemUsedMb != tables[0].memUsedMb {
		t.Errorf("Incorrect node MemUsedMb generated, got: %d, want: %d.",
			node.MemUsedMb,
			tables[0].memUsedMb)
	}

	if node.MemTotalMb != tables[0].memTotalMb {
		t.Errorf("Incorrect node MemTotalMb generated, got: %d, want: %d.",
			node.MemTotalMb,
			tables[0].memTotalMb)
	}
}

func TestGetContainers(t *testing.T) {
	tables := []struct {
		container pfmodel.Container
	}{
		{
			pfmodel.Container{
				Hostname:     "test-01",
				Ipaddress:    "192.168.1.100",
				NodeHostname: "",
				Status:       "PENDING",
				Source: pfmodel.Source{
					Type:        "image",
					Mode:        "pull",
					Alias:       "16.04",
					Certificate: "random",
					Remote: pfmodel.Remote{
						Server:   "https://cloud-images.ubuntu.com/releases",
						Protocol: "simplestream",
						AuthType: "none",
					},
				},
			},
		},
		{
			pfmodel.Container{
				Hostname:     "test-02",
				Ipaddress:    "192.168.1.101",
				NodeHostname: "node-01",
				Status:       "SCHEDULED",
				Source: pfmodel.Source{
					Type:        "image",
					Mode:        "pull",
					Alias:       "16.04",
					Certificate: "random",
					Remote: pfmodel.Remote{
						Server:   "https://cloud-images.ubuntu.com/releases",
						Protocol: "simplestream",
						AuthType: "none",
					},
				},
			},
		},
		{
			pfmodel.Container{
				Hostname:     "test-03",
				Ipaddress:    "192.168.1.102",
				NodeHostname: "",
				Status:       "PENDING",
				Source: pfmodel.Source{
					Type:        "image",
					Mode:        "pull",
					Alias:       "16.04",
					Certificate: "random",
					Remote: pfmodel.Remote{
						Server:   "https://cloud-images.ubuntu.com/releases",
						Protocol: "simplestream",
						AuthType: "none",
					},
				},
			},
		},
	}

	b := []byte(`{
		"api_version": "1.0",
		"data": {
			"items": [
				{
					"hostname": "test-01",
					"ipaddress": "192.168.1.100",
					"source": { 
						"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04", "certificate": "random",
						"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none"} 
					},
					"node_hostname": "",
					"status":"PENDING"
				},
				{
					"hostname": "test-02",
					"ipaddress": "192.168.1.101",
					"source": { 
						"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04", "certificate": "random",
						"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none"} 
					},
					"node_hostname": "node-01",
					"status":"SCHEDULED"
				},
				{
					"hostname": "test-03",
					"ipaddress": "192.168.1.102",
					"source": { 
						"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04", "certificate": "random",
						"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none"} 
					},
					"node_hostname": "",
					"status":"PENDING"
				}
			]
		}
	}`)

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write(b)
	}))
	defer func() { testServer.Close() }()

	client := NewClient("default", "", &http.Client{}, testServer.URL, map[string]string{})
	containers, _ := client.GetContainers()

	for i, table := range tables {
		if (*containers)[i].Hostname != table.container.Hostname {
			t.Errorf("Incorrect container hostname generated, got: %s, want: %s.",
				(*containers)[i].Hostname,
				table.container.Hostname)
		}

		if (*containers)[i].Source.Type != tables[0].container.Source.Type {
			t.Errorf("Incorrect container source type generated, got: %s, want: %s.",
				(*containers)[i].Source.Type,
				tables[0].container.Source.Type)
		}

		if (*containers)[i].Source.Mode != tables[0].container.Source.Mode {
			t.Errorf("Incorrect container source mode generated, got: %s, want: %s.",
				(*containers)[i].Source.Mode,
				tables[0].container.Source.Mode)
		}

		if (*containers)[i].Source.Alias != tables[0].container.Source.Alias {
			t.Errorf("Incorrect container source alias generated, got: %s, want: %s.",
				(*containers)[i].Source.Alias,
				tables[0].container.Source.Alias)
		}

		if (*containers)[i].Source.Certificate != tables[0].container.Source.Certificate {
			t.Errorf("Incorrect container source certificate generated, got: %s, want: %s.",
				(*containers)[i].Source.Certificate,
				tables[0].container.Source.Certificate)
		}

		if (*containers)[i].Source.Remote.Server != tables[0].container.Source.Remote.Server {
			t.Errorf("Incorrect container remote server generated, got: %s, want: %s.",
				(*containers)[i].Source.Remote.Server,
				tables[0].container.Source.Remote.Server)
		}

		if (*containers)[i].Source.Remote.Protocol != tables[0].container.Source.Remote.Protocol {
			t.Errorf("Incorrect container remote protocol generated, got: %s, want: %s.",
				(*containers)[i].Source.Remote.Protocol,
				tables[0].container.Source.Remote.Protocol)
		}

		if (*containers)[i].Source.Remote.AuthType != tables[0].container.Source.Remote.AuthType {
			t.Errorf("Incorrect container remote auth_type generated, got: %s, want: %s.",
				(*containers)[i].Source.Remote.AuthType,
				tables[0].container.Source.Remote.AuthType)
		}

		if (*containers)[i].NodeHostname != table.container.NodeHostname {
			t.Errorf("Incorrect container Node Hostname generated, got: %s, want: %s.",
				(*containers)[i].NodeHostname,
				table.container.NodeHostname)
		}

		if (*containers)[i].Status != table.container.Status {
			t.Errorf("Incorrect container Status generated, got: %s, want: %s.",
				(*containers)[i].Status,
				table.container.Status)
		}
	}
}

func TestGetContainer(t *testing.T) {
	tables := []struct {
		container pfmodel.Container
	}{
		{
			pfmodel.Container{
				Hostname:     "test-02",
				Ipaddress:    "192.168.1.101",
				NodeHostname: "node-01",
				Status:       "SCHEDULED",
				Source: pfmodel.Source{
					Type:        "image",
					Mode:        "pull",
					Alias:       "16.04",
					Certificate: "random",
					Remote: pfmodel.Remote{
						Server:   "https://cloud-images.ubuntu.com/releases",
						Protocol: "simplestream",
						AuthType: "none",
					},
				},
			},
		},
	}

	b := []byte(`{
		"api_version": "1.0",
		"data": {
			"id": "1",
			"hostname": "test-02",
			"ipaddress": "192.168.1.101",
			"source": {
				"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04", "certificate": "random",
				"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none"}
			},
			"node_hostname": "node-01",
			"status":"SCHEDULED"
		}
	}`)

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write(b)
	}))
	defer func() { testServer.Close() }()

	client := NewClient("default", "", &http.Client{}, testServer.URL, map[string]string{})
	container, _ := client.GetContainer(tables[0].container.Hostname)

	if container.Hostname != tables[0].container.Hostname {
		t.Errorf("Incorrect container hostname generated, got: %s, want: %s.",
			container.Hostname,
			tables[0].container.Hostname)
	}

	if container.NodeHostname != tables[0].container.NodeHostname {
		t.Errorf("Incorrect container NodeHostname generated, got: %s, want: %s.",
			container.NodeHostname,
			tables[0].container.NodeHostname)
	}

	if container.Status != tables[0].container.Status {
		t.Errorf("Incorrect container Status generated, got: %s, want: %s.",
			container.Status,
			tables[0].container.Status)
	}

	if container.Ipaddress != tables[0].container.Ipaddress {
		t.Errorf("Incorrect container Ipaddress generated, got: %s, want: %s.",
			container.Ipaddress,
			tables[0].container.Ipaddress)
	}

	if container.Source.Type != tables[0].container.Source.Type {
		t.Errorf("Incorrect container source type generated, got: %s, want: %s.",
			container.Source.Type,
			tables[0].container.Source.Type)
	}

	if container.Source.Mode != tables[0].container.Source.Mode {
		t.Errorf("Incorrect container source mode generated, got: %s, want: %s.",
			container.Source.Mode,
			tables[0].container.Source.Mode)
	}

	if container.Source.Alias != tables[0].container.Source.Alias {
		t.Errorf("Incorrect container source alias generated, got: %s, want: %s.",
			container.Source.Alias,
			tables[0].container.Source.Alias)
	}

	if container.Source.Certificate != tables[0].container.Source.Certificate {
		t.Errorf("Incorrect container source certificate generated, got: %s, want: %s.",
			container.Source.Certificate,
			tables[0].container.Source.Certificate)
	}

	if container.Source.Remote.Server != tables[0].container.Source.Remote.Server {
		t.Errorf("Incorrect container remote server generated, got: %s, want: %s.",
			container.Source.Remote.Server,
			tables[0].container.Source.Remote.Server)
	}

	if container.Source.Remote.Protocol != tables[0].container.Source.Remote.Protocol {
		t.Errorf("Incorrect container remote protocol generated, got: %s, want: %s.",
			container.Source.Remote.Protocol,
			tables[0].container.Source.Remote.Protocol)
	}

	if container.Source.Remote.AuthType != tables[0].container.Source.Remote.AuthType {
		t.Errorf("Incorrect container remote auth_type generated, got: %s, want: %s.",
			container.Source.Remote.AuthType,
			tables[0].container.Source.Remote.AuthType)
	}
}

func TestCreateContainer(t *testing.T) {
	tables := []struct {
		container pfmodel.Container
	}{
		{
			pfmodel.Container{
				Hostname: "test-01",
				Source: pfmodel.Source{
					Type:        "image",
					Mode:        "pull",
					Alias:       "16.04",
					Certificate: "random",
					Remote: pfmodel.Remote{
						Server:   "https://cloud-images.ubuntu.com/releases",
						Protocol: "simplestream",
						AuthType: "none",
					},
				},
			},
		},
	}

	b := []byte(`{
		"api_version": "1.0",
		"data": {
			"hostname": "test-01",
			"source": {
				"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04", "certificate": "random",
				"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none"}
			}
		}
	}`)

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write(b)
	}))
	defer func() { testServer.Close() }()

	client := NewClient("default", "", &http.Client{}, testServer.URL, map[string]string{})
	container, _ := client.CreateContainer(tables[0].container)

	if container.Hostname != tables[0].container.Hostname {
		t.Errorf("Incorrect container hostname generated, got: %s, want: %s.",
			container.Hostname,
			tables[0].container.Hostname)
	}

	if container.Source.Type != tables[0].container.Source.Type {
		t.Errorf("Incorrect container source type generated, got: %s, want: %s.",
			container.Source.Type,
			tables[0].container.Source.Type)
	}

	if container.Source.Mode != tables[0].container.Source.Mode {
		t.Errorf("Incorrect container source mode generated, got: %s, want: %s.",
			container.Source.Mode,
			tables[0].container.Source.Mode)
	}

	if container.Source.Alias != tables[0].container.Source.Alias {
		t.Errorf("Incorrect container source alias generated, got: %s, want: %s.",
			container.Source.Alias,
			tables[0].container.Source.Alias)
	}

	if container.Source.Certificate != tables[0].container.Source.Certificate {
		t.Errorf("Incorrect container source certificate generated, got: %s, want: %s.",
			container.Source.Certificate,
			tables[0].container.Source.Certificate)
	}

	if container.Source.Remote.Server != tables[0].container.Source.Remote.Server {
		t.Errorf("Incorrect container remote server generated, got: %s, want: %s.",
			container.Source.Remote.Server,
			tables[0].container.Source.Remote.Server)
	}

	if container.Source.Remote.Protocol != tables[0].container.Source.Remote.Protocol {
		t.Errorf("Incorrect container remote protocol generated, got: %s, want: %s.",
			container.Source.Remote.Protocol,
			tables[0].container.Source.Remote.Protocol)
	}

	if container.Source.Remote.AuthType != tables[0].container.Source.Remote.AuthType {
		t.Errorf("Incorrect container remote auth_type generated, got: %s, want: %s.",
			container.Source.Remote.AuthType,
			tables[0].container.Source.Remote.AuthType)
	}
}

func TestDeleteContainer(t *testing.T) {
	tables := []struct {
		hostname string
	}{
		{"test-01"},
	}

	b := []byte(`{
		"api_version": "1.0",
		"data": {
			"id": 1,
			"hostname": "test-01",
			"ipaddress": "192.168.1.100",
			"source": {
				"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04", "certificate": "random",
				"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none"}
			},
			"node_hostname": "test-01",
			"status": "SCHEDULE_DELETION"
		}
	}`)

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write(b)
	}))
	defer func() { testServer.Close() }()

	client := NewClient("default", "", &http.Client{}, testServer.URL, map[string]string{})
	container, _ := client.DeleteContainer(tables[0].hostname)

	if container.Hostname != tables[0].hostname {
		t.Errorf("Incorrect container hostname generated, got: %s, want: %s.",
			container.Hostname,
			tables[0].hostname)
	}
}

func TestRescheduleContainer(t *testing.T) {
	tables := []struct {
		hostname string
	}{
		{"test-01"},
	}

	b := []byte(`{
		"api_version": "1.0",
		"data": {
			"id": 1,
			"hostname": "test-01",
			"ipaddress": "192.168.1.100",
			"source": {
				"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04", "certificate": "random",
				"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none"}
			},
			"node_hostname": "test-01",
			"status": "PENDING"
		}
	}`)

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write(b)
	}))
	defer func() { testServer.Close() }()

	client := NewClient("default", "", &http.Client{}, testServer.URL, map[string]string{})
	container, _ := client.RescheduleContainer(tables[0].hostname)

	if container.Hostname != tables[0].hostname {
		t.Errorf("Incorrect container hostname generated, got: %s, want: %s.",
			container.Hostname,
			tables[0].hostname)
	}
}

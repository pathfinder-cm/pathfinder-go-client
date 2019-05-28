package ext

import (
	"net/http"
	"net/http/httptest"
	"testing"
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
		hostname       string
		ipaddress      string
		image_alias    string
		image_server   string
		image_protocol string
		nodeHostname   string
		status         string
	}{
		{"test-01", "192.168.1.100", "18.04", "ubuntu", "simplestream", "", "PENDING"},
		{"test-02", "192.168.1.101", "18.04", "ubuntu", "simplestream", "test-01", "SCHEDULED"},
		{"test-03", "192.168.1.102", "18.04", "ubuntu", "simplestream", "", "PENDING"},
	}

	b := []byte(`{
		"api_version": "1.0",
		"data": {
			"items": [
				{"hostname": "test-01", "ipaddress": "192.168.1.100", "image_alias": "18.04", "image_server": "ubuntu", "image_protocol": "simplestream", "node_hostname": "", "status": "PENDING"},
				{"hostname": "test-02", "ipaddress": "192.168.1.101", "image_alias": "18.04", "image_server": "ubuntu", "image_protocol": "simplestream", "node_hostname": "test-01", "status": "SCHEDULED"},
				{"hostname": "test-03", "ipaddress": "192.168.1.102", "image_alias": "18.04", "image_server": "ubuntu", "image_protocol": "simplestream", "node_hostname": "", "status": "PENDING"}
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
		if (*containers)[i].Hostname != table.hostname {
			t.Errorf("Incorrect container hostname generated, got: %s, want: %s.",
				(*containers)[i].Hostname,
				table.hostname)
		}

		if (*containers)[i].Ipaddress != table.ipaddress {
			t.Errorf("Incorrect container ipaddress generated, got: %s, want: %s.",
				(*containers)[i].Ipaddress,
				table.ipaddress)
		}

		if (*containers)[i].ImageAlias != table.image_alias {
			t.Errorf("Incorrect container Image generated, got: %s, want: %s.",
				(*containers)[i].ImageAlias,
				table.image_alias)
		}

		if (*containers)[i].ImageServer != table.image_server {
			t.Errorf("Incorrect container Image server generated, got: %s, want: %s.",
				(*containers)[i].ImageServer,
				table.image_server)
		}

		if (*containers)[i].ImageProtocol != table.image_protocol {
			t.Errorf("Incorrect container Image protocol generated, got: %s, want: %s.",
				(*containers)[i].ImageProtocol,
				table.image_protocol)
		}

		if (*containers)[i].NodeHostname != table.nodeHostname {
			t.Errorf("Incorrect container Node Hostname generated, got: %s, want: %s.",
				(*containers)[i].NodeHostname,
				table.image_alias)
		}

		if (*containers)[i].Status != table.status {
			t.Errorf("Incorrect container Status generated, got: %s, want: %s.",
				(*containers)[i].Status,
				table.status)
		}
	}
}

func TestGetContainer(t *testing.T) {
	tables := []struct {
		hostname       string
		ipaddress      string
		image_alias    string
		image_server   string
		image_protocol string
		nodeHostname   string
		status         string
	}{
		{"test-02", "192.168.1.101", "18.04", "ubuntu", "simplestream", "test-01", "SCHEDULED"},
	}

	b := []byte(`{
		"api_version": "1.0",
		"data": {
			"id": 1,
			"hostname": "test-02",
			"ipaddress": "192.168.1.101",
			"image_alias": "18.04",
			"image_server": "ubuntu",
			"image_protocol": "simplestream",
			"node_hostname": "test-01",
			"status": "SCHEDULED"
		}
	}`)

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write(b)
	}))
	defer func() { testServer.Close() }()

	client := NewClient("default", "", &http.Client{}, testServer.URL, map[string]string{})
	container, _ := client.GetContainer(tables[0].hostname)
	if container.Hostname != tables[0].hostname {
		t.Errorf("Incorrect container hostname generated, got: %s, want: %s.",
			container.Hostname,
			tables[0].hostname)
	}

	if container.Ipaddress != tables[0].ipaddress {
		t.Errorf("Incorrect container ipaddress generated, got: %s, want: %s.",
			container.Ipaddress,
			tables[0].ipaddress)
	}

	if container.ImageAlias != tables[0].image_alias {
		t.Errorf("Incorrect container Image generated, got: %s, want: %s.",
			container.ImageAlias,
			tables[0].image_alias)
	}

	if container.ImageServer != tables[0].image_server {
		t.Errorf("Incorrect container Image server generated, got: %s, want: %s.",
			container.ImageServer,
			tables[0].image_server)
	}

	if container.ImageProtocol != tables[0].image_protocol {
		t.Errorf("Incorrect container Image protocol generated, got: %s, want: %s.",
			container.ImageProtocol,
			tables[0].image_protocol)
	}

	if container.NodeHostname != tables[0].nodeHostname {
		t.Errorf("Incorrect container Node Hostname generated, got: %s, want: %s.",
			container.NodeHostname,
			tables[0].nodeHostname)
	}

	if container.Status != tables[0].status {
		t.Errorf("Incorrect container Status generated, got: %s, want: %s.",
			container.Status,
			tables[0].status)
	}
}

func TestCreateContainer(t *testing.T) {
	tables := []struct {
		hostname       string
		image_alias    string
		image_server   string
		image_protocol string
	}{
		{"test-01", "18.04", "ubuntu", "simplestream"},
	}

	b := []byte(`{
		"api_version": "1.0",
		"data": {
			"id": 1,
			"hostname": "test-01",
			"ipaddress": "192.168.1.100",
			"image_alias": "18.04",
			"image_server": "ubuntu",
			"image_protocol": "simplestream",
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
	container, _ := client.CreateContainer(tables[0].hostname, tables[0].image_alias, tables[0].image_server, tables[0].image_protocol)

	if container.Hostname != tables[0].hostname {
		t.Errorf("Incorrect container hostname generated, got: %s, want: %s.",
			container.Hostname,
			tables[0].hostname)
	}

	if container.ImageAlias != tables[0].image_alias {
		t.Errorf("Incorrect container Image generated, got: %s, want: %s.",
			container.ImageAlias,
			tables[0].image_alias)
	}

	if container.ImageServer != tables[0].image_server {
		t.Errorf("Incorrect container Image server generated, got: %s, want: %s.",
			container.ImageServer,
			tables[0].image_server)
	}

	if container.ImageProtocol != tables[0].image_protocol {
		t.Errorf("Incorrect container Image protocol generated, got: %s, want: %s.",
			container.ImageProtocol,
			tables[0].image_protocol)
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
			"image_alias": "18.04",
			"image_server": "ubuntu",
			"image_protocol": "simplestream",
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
			"image_alias": "18.04",
			"image_server": "ubuntu",
			"image_protocol": "simplestream",
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

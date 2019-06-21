package pfclient

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pathfinder-cm/pathfinder-go-client/pfmodel"
)

func TestRegister(t *testing.T) {
	node := "test-01"
	ipaddress := "127.0.0.1"

	b := []byte(`{
		"api_version": "1.0",
		"data": {
			"id": 1,
			"cluster_id": 1,
			"cluster_name": "default",
			"hostname": "test-01",
			"authentication_token": "123"
		}
	}`)

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write(b)
	}))
	defer func() { testServer.Close() }()

	pfclient := NewPfclient("default", "", &http.Client{}, testServer.URL, map[string]string{})
	ok, _ := pfclient.Register(node, ipaddress)

	if !ok {
		t.Errorf("Registration unsuccessful")
	}
}

func TestFetchContainersFromServer(t *testing.T) {
	node := "test-01"
	tables := []struct {
		hostname    string
		source_type string
		mode        string
		alias       string
		certificate string
		server      string
		protocol    string
		auth_type   string
		status      string
	}{
		{"test-01", "image", "pull", "16.04", "random", "https://cloud-images.ubuntu.com/releases", "simplestream", "none", "SCHEDULED"},
		{"test-02", "image", "pull", "16.04", "random", "https://cloud-images.ubuntu.com/releases", "simplestream", "none", "SCHEDULED"},
		{"test-03", "image", "pull", "16.04", "random", "https://cloud-images.ubuntu.com/releases", "simplestream", "none", "SCHEDULED"},
	}

	b := []byte(`{
		"api_version": "1.0",
		"data": {
			"items": [
				{
					"hostname": "test-01", 
					"source": { 
						"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04", "certificate": "random",
						"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none"} 
					}, 
					"status":"SCHEDULED"
				},
				{
					"hostname": "test-02", 
					"source": { 
						"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04", "certificate": "random",
						"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none"} 
					}, 
					"status":"SCHEDULED"
				},
				{
					"hostname": "test-03", 
					"source": { 
						"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04", "certificate": "random",
						"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none"} 
					}, 
					"status":"SCHEDULED"
				}
			]
		}
	}`)

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write(b)
	}))
	defer func() { testServer.Close() }()

	pfclient := NewPfclient("default", "", &http.Client{}, testServer.URL, map[string]string{})
	cl, _ := pfclient.FetchContainersFromServer(node)
	for i, table := range tables {
		if (*cl)[i].Hostname != table.hostname {
			t.Errorf("Incorrect container hostname fetched, got: %s, want: %s.",
				(*cl)[i].Hostname,
				table.hostname)
		}

		if (*cl)[i].Source.Type != table.source_type {
			t.Errorf("Incorrect container source type fetched, got: %s, want: %s.",
				(*cl)[i].Source.Type,
				table.source_type)
		}

		if (*cl)[i].Source.Mode != table.mode {
			t.Errorf("Incorrect container source mode fetched, got: %s, want: %s.",
				(*cl)[i].Source.Mode,
				table.mode)
		}

		if (*cl)[i].Source.Alias != table.alias {
			t.Errorf("Incorrect container source alias fetched, got: %s, want: %s.",
				(*cl)[i].Source.Alias,
				table.alias)
		}

		if (*cl)[i].Source.Certificate != table.certificate {
			t.Errorf("Incorrect container source certificate fetched, got: %s, want: %s.",
				(*cl)[i].Source.Certificate,
				table.certificate)
		}

		if (*cl)[i].Source.Remote.Server != table.server {
			t.Errorf("Incorrect container remote server fetched, got: %s, want: %s.",
				(*cl)[i].Source.Remote.Server,
				table.server)
		}

		if (*cl)[i].Source.Remote.Protocol != table.protocol {
			t.Errorf("Incorrect container remote protocol fetched, got: %s, want: %s.",
				(*cl)[i].Source.Remote.Protocol,
				table.protocol)
		}

		if (*cl)[i].Source.Remote.AuthType != table.auth_type {
			t.Errorf("Incorrect container remote auth_type fetched, got: %s, want: %s.",
				(*cl)[i].Source.Remote.AuthType,
				table.auth_type)
		}

		if (*cl)[i].Status != table.status {
			t.Errorf("Incorrect container status fetched, got: %s, want: %s.",
				(*cl)[i].Status,
				table.status)
		}
	}
}

func TestUpdateIpaddress(t *testing.T) {
	tables := []struct {
		node      string
		hostname  string
		ipaddress string
	}{
		{"test-01", "test-c-01", "127.0.0.1"},
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
	}))
	defer func() { testServer.Close() }()

	pfclient := NewPfclient("default", "", &http.Client{}, testServer.URL, map[string]string{})
	ok, _ := pfclient.UpdateIpaddress(tables[0].node, tables[0].hostname, tables[0].ipaddress)
	if ok != true {
		t.Errorf("Error when updating container ipaddress")
	}
}

func TestMarkContainerAsProvisioned(t *testing.T) {
	tables := []struct {
		node     string
		hostname string
	}{
		{"test-01", "test-c-01"},
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
	}))
	defer func() { testServer.Close() }()

	pfclient := NewPfclient("default", "", &http.Client{}, testServer.URL, map[string]string{})
	ok, _ := pfclient.MarkContainerAsProvisioned(tables[0].node, tables[0].hostname)
	if ok != true {
		t.Errorf("Error when marking container as provisioned")
	}
}

func TestMarkContainerAsProvisionError(t *testing.T) {
	tables := []struct {
		node     string
		hostname string
	}{
		{"test-01", "test-c-01"},
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
	}))
	defer func() { testServer.Close() }()

	pfclient := NewPfclient("default", "", &http.Client{}, testServer.URL, map[string]string{})
	ok, _ := pfclient.MarkContainerAsProvisionError(tables[0].node, tables[0].hostname)
	if ok != true {
		t.Errorf("Error when marking container as provision_error")
	}
}

func TestMarkContainerAsDeleted(t *testing.T) {
	tables := []struct {
		node     string
		hostname string
	}{
		{"test-01", "test-c-01"},
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
	}))
	defer func() { testServer.Close() }()

	pfclient := NewPfclient("default", "", &http.Client{}, testServer.URL, map[string]string{})
	ok, _ := pfclient.MarkContainerAsDeleted(tables[0].node, tables[0].hostname)
	if ok != true {
		t.Errorf("Error when marking container as deleted")
	}
}

func TestStoreMetrics(t *testing.T) {
	tables := []struct {
		metrics *pfmodel.Metrics
	}{
		{
			&pfmodel.Metrics{
				Memory: &pfmodel.Memory{
					Used:  100,
					Free:  200,
					Total: 400,
				},
			},
		},
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
	}))
	defer func() { testServer.Close() }()

	pfclient := NewPfclient("default", "", &http.Client{}, testServer.URL, map[string]string{})
	ok, _ := pfclient.StoreMetrics(tables[0].metrics)
	if ok != true {
		t.Errorf("Error when storing metrics")
	}
}

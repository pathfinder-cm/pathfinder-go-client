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
		hostname       string
		image_alias    string
		image_server   string
		image_protocol string
		status         string
	}{
		{"test-01", "16.04", "ubuntu", "simplestream", "SCHEDULED"},
		{"test-02", "16.04", "ubuntu", "simplestream", "SCHEDULED"},
		{"test-03", "16.04", "ubuntu", "simplestream", "SCHEDULED"},
	}

	b := []byte(`{
		"api_version": "1.0",
		"data": {
			"items": [
				{"hostname": "test-01", "image_alias": "16.04", "image_server": "ubuntu", "image_protocol": "simplestream", "status": "SCHEDULED"},
				{"hostname": "test-02", "image_alias": "16.04", "image_server": "ubuntu", "image_protocol": "simplestream", "status": "SCHEDULED"},
				{"hostname": "test-03", "image_alias": "16.04", "image_server": "ubuntu", "image_protocol": "simplestream", "status": "SCHEDULED"}
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

		if (*cl)[i].ImageAlias != table.image_alias {
			t.Errorf("Incorrect container image alias fetched, got: %s, want: %s.",
				(*cl)[i].ImageAlias,
				table.image_alias)
		}

		if (*cl)[i].ImageServer != table.image_server {
			t.Errorf("Incorrect container image server fetched, got: %s, want: %s.",
				(*cl)[i].ImageServer,
				table.image_server)
		}

		if (*cl)[i].ImageProtocol != table.image_protocol {
			t.Errorf("Incorrect container image protocol fetched, got: %s, want: %s.",
				(*cl)[i].ImageProtocol,
				table.image_protocol)
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

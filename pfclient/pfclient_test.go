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

func TestFetchScheduledContainersFromServer(t *testing.T) {
	node := "test-01"
	bootstrappers := []pfmodel.Bootstrapper{
		pfmodel.Bootstrapper{
			Type:         "chef-solo",
			CookbooksUrl: "127.0.0.1",
			Attributes:   "{}",
		},
	}
	tables := []struct {
		container pfmodel.Container
	}{
		{
			pfmodel.Container{
				Hostname:      "test-01",
				Status:        "SCHEDULED",
				Bootstrappers: bootstrappers,
				Source: pfmodel.Source{
					Type:  "image",
					Mode:  "pull",
					Alias: "16.04",
					Remote: pfmodel.Remote{
						Server:      "https://cloud-images.ubuntu.com/releases",
						Protocol:    "simplestream",
						AuthType:    "none",
						Certificate: "random",
					},
				},
			},
		},
		{
			pfmodel.Container{
				Hostname:      "test-02",
				Status:        "SCHEDULED",
				Bootstrappers: bootstrappers,
				Source: pfmodel.Source{
					Type:  "image",
					Mode:  "pull",
					Alias: "16.04",
					Remote: pfmodel.Remote{
						Server:      "https://cloud-images.ubuntu.com/releases",
						Protocol:    "simplestream",
						AuthType:    "none",
						Certificate: "random",
					},
				},
			},
		},
		{
			pfmodel.Container{
				Hostname:      "test-03",
				Status:        "SCHEDULED",
				Bootstrappers: bootstrappers,
				Source: pfmodel.Source{
					Type:  "image",
					Mode:  "pull",
					Alias: "16.04",
					Remote: pfmodel.Remote{
						Server:      "https://cloud-images.ubuntu.com/releases",
						Protocol:    "simplestream",
						AuthType:    "none",
						Certificate: "random",
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
					"source": {
						"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04",
						"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none", "certificate": "random"}
					},
					"status":"SCHEDULED",
					"bootstrappers": [{
						"bootstrap_type":"chef-solo",
						"bootstrap_cookbooks_url":"127.0.0.1",
						"bootstrap_attributes":"{}"
					}]
				},
				{
					"hostname": "test-02",
					"source": {
						"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04",
						"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none", "certificate": "random"}
					},
					"status":"SCHEDULED",
					"bootstrappers": [{
						"bootstrap_type":"chef-solo",
						"bootstrap_cookbooks_url":"127.0.0.1",
						"bootstrap_attributes":"{}"
					}]
				},
				{
					"hostname": "test-03",
					"source": {
						"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04",
						"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none", "certificate": "random"}
					},
					"status":"SCHEDULED",
					"bootstrappers": [{
						"bootstrap_type":"chef-solo",
						"bootstrap_cookbooks_url":"127.0.0.1",
						"bootstrap_attributes":"{}"
					}]
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
	cl, _ := pfclient.FetchScheduledContainersFromServer(node)
	for i, table := range tables {
		if (*cl)[i].Hostname != table.container.Hostname {
			t.Errorf("Incorrect container hostname fetched, got: %s, want: %s.",
				(*cl)[i].Hostname,
				table.container.Hostname)
		}

		if (*cl)[i].Source.Type != table.container.Source.Type {
			t.Errorf("Incorrect container source type fetched, got: %s, want: %s.",
				(*cl)[i].Source.Type,
				table.container.Source.Type)
		}

		if (*cl)[i].Source.Mode != table.container.Source.Mode {
			t.Errorf("Incorrect container source mode fetched, got: %s, want: %s.",
				(*cl)[i].Source.Mode,
				table.container.Source.Mode)
		}

		if (*cl)[i].Source.Alias != table.container.Source.Alias {
			t.Errorf("Incorrect container source alias fetched, got: %s, want: %s.",
				(*cl)[i].Source.Alias,
				table.container.Source.Alias)
		}

		if (*cl)[i].Source.Remote.Server != table.container.Source.Remote.Server {
			t.Errorf("Incorrect container remote server fetched, got: %s, want: %s.",
				(*cl)[i].Source.Remote.Server,
				table.container.Source.Remote.Server)
		}

		if (*cl)[i].Source.Remote.Protocol != table.container.Source.Remote.Protocol {
			t.Errorf("Incorrect container remote protocol fetched, got: %s, want: %s.",
				(*cl)[i].Source.Remote.Protocol,
				table.container.Source.Remote.Protocol)
		}

		if (*cl)[i].Source.Remote.AuthType != table.container.Source.Remote.AuthType {
			t.Errorf("Incorrect container remote auth_type fetched, got: %s, want: %s.",
				(*cl)[i].Source.Remote.AuthType,
				table.container.Source.Remote.AuthType)
		}

		if (*cl)[i].Source.Remote.Certificate != table.container.Source.Remote.Certificate {
			t.Errorf("Incorrect container remote certificate fetched, got: %s, want: %s.",
				(*cl)[i].Source.Remote.Certificate,
				table.container.Source.Remote.Certificate)
		}

		if (*cl)[i].Status != table.container.Status {
			t.Errorf("Incorrect container status fetched, got: %s, want: %s.",
				(*cl)[i].Status,
				table.container.Status)
		}

		if (*cl)[i].Bootstrappers[0].Type != table.container.Bootstrappers[0].Type {
			t.Errorf("Incorrect container bootstrap_type generated, got: %s, want: %s.",
				(*cl)[i].Bootstrappers[0].Type,
				table.container.Bootstrappers[0].Type)
		}

		if (*cl)[i].Bootstrappers[0].CookbooksUrl != table.container.Bootstrappers[0].CookbooksUrl {
			t.Errorf("Incorrect container bootstrap_cookbooks_url generated, got: %s, want: %s.",
				(*cl)[i].Bootstrappers[0].CookbooksUrl,
				table.container.Bootstrappers[0].CookbooksUrl)
		}

		if (*cl)[i].Bootstrappers[0].Attributes != table.container.Bootstrappers[0].Attributes {
			t.Errorf("Incorrect container bootstrap_attributes generated, got: %s, want: %s.",
				(*cl)[i].Bootstrappers[0].Attributes,
				table.container.Bootstrappers[0].Attributes)
		}
	}
}

func TestFetchProvisionedContainersFromServer(t *testing.T) {
	node := "test-01"
	bootstrappers := []pfmodel.Bootstrapper{
		pfmodel.Bootstrapper{
			Type:         "chef-solo",
			CookbooksUrl: "127.0.0.1",
			Attributes:   "{}",
		},
	}
	tables := []struct {
		container pfmodel.Container
	}{
		{
			pfmodel.Container{
				Hostname:      "test-01",
				Status:        "PROVISIONED",
				Bootstrappers: bootstrappers,
				Source: pfmodel.Source{
					Type:  "image",
					Mode:  "pull",
					Alias: "16.04",
					Remote: pfmodel.Remote{
						Server:      "https://cloud-images.ubuntu.com/releases",
						Protocol:    "simplestream",
						AuthType:    "none",
						Certificate: "random",
					},
				},
			},
		},
		{
			pfmodel.Container{
				Hostname:      "test-02",
				Status:        "PROVISIONED",
				Bootstrappers: bootstrappers,
				Source: pfmodel.Source{
					Type:  "image",
					Mode:  "pull",
					Alias: "16.04",
					Remote: pfmodel.Remote{
						Server:      "https://cloud-images.ubuntu.com/releases",
						Protocol:    "simplestream",
						AuthType:    "none",
						Certificate: "random",
					},
				},
			},
		},
		{
			pfmodel.Container{
				Hostname:      "test-03",
				Status:        "PROVISIONED",
				Bootstrappers: bootstrappers,
				Source: pfmodel.Source{
					Type:  "image",
					Mode:  "pull",
					Alias: "16.04",
					Remote: pfmodel.Remote{
						Server:      "https://cloud-images.ubuntu.com/releases",
						Protocol:    "simplestream",
						AuthType:    "none",
						Certificate: "random",
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
					"source": {
						"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04",
						"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none", "certificate": "random"}
					},
					"status":"PROVISIONED",
					"bootstrappers": [{
						"bootstrap_type":"chef-solo",
						"bootstrap_cookbooks_url":"127.0.0.1",
						"bootstrap_attributes":"{}"
					}]
				},
				{
					"hostname": "test-02",
					"source": {
						"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04",
						"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none", "certificate": "random"}
					},
					"status":"PROVISIONED",
					"bootstrappers": [{
						"bootstrap_type":"chef-solo",
						"bootstrap_cookbooks_url":"127.0.0.1",
						"bootstrap_attributes":"{}"
					}]
				},
				{
					"hostname": "test-03",
					"source": {
						"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04",
						"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none", "certificate": "random"}
					},
					"status":"PROVISIONED",
					"bootstrappers": [{
						"bootstrap_type":"chef-solo",
						"bootstrap_cookbooks_url":"127.0.0.1",
						"bootstrap_attributes":"{}"
					}]
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
	cl, _ := pfclient.FetchProvisionedContainersFromServer(node)
	for i, table := range tables {
		if (*cl)[i].Hostname != table.container.Hostname {
			t.Errorf("Incorrect container hostname fetched, got: %s, want: %s.",
				(*cl)[i].Hostname,
				table.container.Hostname)
		}

		if (*cl)[i].Source.Type != table.container.Source.Type {
			t.Errorf("Incorrect container source type fetched, got: %s, want: %s.",
				(*cl)[i].Source.Type,
				table.container.Source.Type)
		}

		if (*cl)[i].Source.Mode != table.container.Source.Mode {
			t.Errorf("Incorrect container source mode fetched, got: %s, want: %s.",
				(*cl)[i].Source.Mode,
				table.container.Source.Mode)
		}

		if (*cl)[i].Source.Alias != table.container.Source.Alias {
			t.Errorf("Incorrect container source alias fetched, got: %s, want: %s.",
				(*cl)[i].Source.Alias,
				table.container.Source.Alias)
		}

		if (*cl)[i].Source.Remote.Server != table.container.Source.Remote.Server {
			t.Errorf("Incorrect container remote server fetched, got: %s, want: %s.",
				(*cl)[i].Source.Remote.Server,
				table.container.Source.Remote.Server)
		}

		if (*cl)[i].Source.Remote.Protocol != table.container.Source.Remote.Protocol {
			t.Errorf("Incorrect container remote protocol fetched, got: %s, want: %s.",
				(*cl)[i].Source.Remote.Protocol,
				table.container.Source.Remote.Protocol)
		}

		if (*cl)[i].Source.Remote.AuthType != table.container.Source.Remote.AuthType {
			t.Errorf("Incorrect container remote auth_type fetched, got: %s, want: %s.",
				(*cl)[i].Source.Remote.AuthType,
				table.container.Source.Remote.AuthType)
		}

		if (*cl)[i].Source.Remote.Certificate != table.container.Source.Remote.Certificate {
			t.Errorf("Incorrect container remote certificate fetched, got: %s, want: %s.",
				(*cl)[i].Source.Remote.Certificate,
				table.container.Source.Remote.Certificate)
		}

		if (*cl)[i].Status != table.container.Status {
			t.Errorf("Incorrect container status fetched, got: %s, want: %s.",
				(*cl)[i].Status,
				table.container.Status)
		}

		if (*cl)[i].Bootstrappers[0].Type != table.container.Bootstrappers[0].Type {
			t.Errorf("Incorrect container bootstrap_type generated, got: %s, want: %s.",
				(*cl)[i].Bootstrappers[0].Type,
				table.container.Bootstrappers[0].Type)
		}

		if (*cl)[i].Bootstrappers[0].CookbooksUrl != table.container.Bootstrappers[0].CookbooksUrl {
			t.Errorf("Incorrect container bootstrap_cookbooks_url generated, got: %s, want: %s.",
				(*cl)[i].Bootstrappers[0].CookbooksUrl,
				table.container.Bootstrappers[0].CookbooksUrl)
		}

		if (*cl)[i].Bootstrappers[0].Attributes != table.container.Bootstrappers[0].Attributes {
			t.Errorf("Incorrect container bootstrap_attributes generated, got: %s, want: %s.",
				(*cl)[i].Bootstrappers[0].Attributes,
				table.container.Bootstrappers[0].Attributes)
		}
	}
}

func TestFetchBootstrappedContainersFromServer(t *testing.T) {
	node := "test-01"
	bootstrappers := []pfmodel.Bootstrapper{
		pfmodel.Bootstrapper{
			Type:         "chef-solo",
			CookbooksUrl: "127.0.0.1",
			Attributes:   "{}",
		},
	}
	tables := []struct {
		container pfmodel.Container
	}{
		{
			pfmodel.Container{
				Hostname:      "test-01",
				Status:        "BOOTSTRAPPED",
				Bootstrappers: bootstrappers,
				Source: pfmodel.Source{
					Type:  "image",
					Mode:  "pull",
					Alias: "16.04",
					Remote: pfmodel.Remote{
						Server:      "https://cloud-images.ubuntu.com/releases",
						Protocol:    "simplestream",
						AuthType:    "none",
						Certificate: "random",
					},
				},
			},
		},
		{
			pfmodel.Container{
				Hostname:      "test-02",
				Status:        "BOOTSTRAPPED",
				Bootstrappers: bootstrappers,
				Source: pfmodel.Source{
					Type:  "image",
					Mode:  "pull",
					Alias: "16.04",
					Remote: pfmodel.Remote{
						Server:      "https://cloud-images.ubuntu.com/releases",
						Protocol:    "simplestream",
						AuthType:    "none",
						Certificate: "random",
					},
				},
			},
		},
		{
			pfmodel.Container{
				Hostname:      "test-03",
				Status:        "HEALTHY",
				Bootstrappers: bootstrappers,
				Source: pfmodel.Source{
					Type:  "image",
					Mode:  "pull",
					Alias: "16.04",
					Remote: pfmodel.Remote{
						Server:      "https://cloud-images.ubuntu.com/releases",
						Protocol:    "simplestream",
						AuthType:    "none",
						Certificate: "random",
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
					"source": {
						"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04",
						"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none", "certificate": "random"}
					},
					"status":"BOOTSTRAPPED",
					"bootstrappers": [{
						"bootstrap_type":"chef-solo",
						"bootstrap_cookbooks_url":"127.0.0.1",
						"bootstrap_attributes":"{}"
					}]
				},
				{
					"hostname": "test-02",
					"source": {
						"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04",
						"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none", "certificate": "random"}
					},
					"status":"BOOTSTRAPPED",
					"bootstrappers": [{
						"bootstrap_type":"chef-solo",
						"bootstrap_cookbooks_url":"127.0.0.1",
						"bootstrap_attributes":"{}"
					}]
				},
				{
					"hostname": "test-03",
					"source": {
						"source_type":"image", "mode":"pull", "fingerprint":"", "alias":"16.04",
						"remote": {"server":"https://cloud-images.ubuntu.com/releases", "protocol":"simplestream", "auth_type":"none", "certificate": "random"}
					},
					"status":"HEALTHY",
					"bootstrappers": [{
						"bootstrap_type":"chef-solo",
						"bootstrap_cookbooks_url":"127.0.0.1",
						"bootstrap_attributes":"{}"
					}]
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
	cl, _ := pfclient.FetchProvisionedContainersFromServer(node)
	for i, table := range tables {
		if (*cl)[i].Hostname != table.container.Hostname {
			t.Errorf("Incorrect container hostname fetched, got: %s, want: %s.",
				(*cl)[i].Hostname,
				table.container.Hostname)
		}

		if (*cl)[i].Source.Type != table.container.Source.Type {
			t.Errorf("Incorrect container source type fetched, got: %s, want: %s.",
				(*cl)[i].Source.Type,
				table.container.Source.Type)
		}

		if (*cl)[i].Source.Mode != table.container.Source.Mode {
			t.Errorf("Incorrect container source mode fetched, got: %s, want: %s.",
				(*cl)[i].Source.Mode,
				table.container.Source.Mode)
		}

		if (*cl)[i].Source.Alias != table.container.Source.Alias {
			t.Errorf("Incorrect container source alias fetched, got: %s, want: %s.",
				(*cl)[i].Source.Alias,
				table.container.Source.Alias)
		}

		if (*cl)[i].Source.Remote.Server != table.container.Source.Remote.Server {
			t.Errorf("Incorrect container remote server fetched, got: %s, want: %s.",
				(*cl)[i].Source.Remote.Server,
				table.container.Source.Remote.Server)
		}

		if (*cl)[i].Source.Remote.Protocol != table.container.Source.Remote.Protocol {
			t.Errorf("Incorrect container remote protocol fetched, got: %s, want: %s.",
				(*cl)[i].Source.Remote.Protocol,
				table.container.Source.Remote.Protocol)
		}

		if (*cl)[i].Source.Remote.AuthType != table.container.Source.Remote.AuthType {
			t.Errorf("Incorrect container remote auth_type fetched, got: %s, want: %s.",
				(*cl)[i].Source.Remote.AuthType,
				table.container.Source.Remote.AuthType)
		}

		if (*cl)[i].Source.Remote.Certificate != table.container.Source.Remote.Certificate {
			t.Errorf("Incorrect container remote certificate fetched, got: %s, want: %s.",
				(*cl)[i].Source.Remote.Certificate,
				table.container.Source.Remote.Certificate)
		}

		if (*cl)[i].Status != table.container.Status {
			t.Errorf("Incorrect container status fetched, got: %s, want: %s.",
				(*cl)[i].Status,
				table.container.Status)
		}

		if (*cl)[i].Bootstrappers[0].Type != table.container.Bootstrappers[0].Type {
			t.Errorf("Incorrect container bootstrap_type generated, got: %s, want: %s.",
				(*cl)[i].Bootstrappers[0].Type,
				table.container.Bootstrappers[0].Type)
		}

		if (*cl)[i].Bootstrappers[0].CookbooksUrl != table.container.Bootstrappers[0].CookbooksUrl {
			t.Errorf("Incorrect container bootstrap_cookbooks_url generated, got: %s, want: %s.",
				(*cl)[i].Bootstrappers[0].CookbooksUrl,
				table.container.Bootstrappers[0].CookbooksUrl)
		}

		if (*cl)[i].Bootstrappers[0].Attributes != table.container.Bootstrappers[0].Attributes {
			t.Errorf("Incorrect container bootstrap_attributes generated, got: %s, want: %s.",
				(*cl)[i].Bootstrappers[0].Attributes,
				table.container.Bootstrappers[0].Attributes)
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

func TestMarkContainerAsBootstrapStarted(t *testing.T) {
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
	ok, _ := pfclient.MarkContainerAsBootstrapStarted(tables[0].node, tables[0].hostname)
	if ok != true {
		t.Errorf("Error when marking container as bootstrapped")
	}
}

func TestMarkContainerAsBootstrapped(t *testing.T) {
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
	ok, _ := pfclient.MarkContainerAsBootstrapped(tables[0].node, tables[0].hostname)
	if ok != true {
		t.Errorf("Error when marking container as bootstrapped")
	}
}

func TestMarkContainerAsHealthy(t *testing.T) {
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
	ok, _ := pfclient.MarkContainerAsHealthy(tables[0].node, tables[0].hostname)
	if ok != true {
		t.Errorf("Error when marking container as healthy")
	}
}

func TestMarkContainerAsBootstrapError(t *testing.T) {
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
	ok, _ := pfclient.MarkContainerAsBootstrapError(tables[0].node, tables[0].hostname)
	if ok != true {
		t.Errorf("Error when marking container as bootstrap_error")
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

func TestUpdateContainerStatus(t *testing.T) {
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

	pfclient := pfclient{
		cluster:         "default",
		clusterPassword: "",
		httpClient:      &http.Client{},
		pfServerAddr:    testServer.URL,
		pfApiPath:       map[string]string{},
	}
	ok, _ := updateContainerStatus(&pfclient, tables[0].node, tables[0].hostname, "api/v2/node/containers/mark_bootstrapped")
	if ok != true {
		t.Errorf("Error when updating container status")
	}
}

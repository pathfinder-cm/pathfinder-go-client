package pfclient

import (
	"testing"
)

func TestNewContainerListFromByte(t *testing.T) {
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
	cl, _ := NewContainerListFromByte(b)

	for i, table := range tables {
		if (*cl)[i].Hostname != table.hostname {
			t.Errorf("Incorrect container hostname generated, got: %s, want: %s.",
				(*cl)[i].Hostname,
				table.hostname)
		}

		if (*cl)[i].ImageAlias != table.image_alias {
			t.Errorf("Incorrect container image alias generated, got: %s, want: %s.",
				(*cl)[i].ImageAlias,
				table.image_alias)
		}

		if (*cl)[i].ImageServer != table.image_server {
			t.Errorf("Incorrect container image server generated, got: %s, want: %s.",
				(*cl)[i].ImageServer,
				table.image_server)
		}

		if (*cl)[i].ImageProtocol != table.image_protocol {
			t.Errorf("Incorrect container image protocol generated, got: %s, want: %s.",
				(*cl)[i].ImageProtocol,
				table.image_protocol)
		}

		if (*cl)[i].Status != table.status {
			t.Errorf("Incorrect container status generated, got: %s, want: %s.",
				(*cl)[i].Status,
				table.status)
		}
	}
}

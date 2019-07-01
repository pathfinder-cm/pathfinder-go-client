package pfclient

import (
	"testing"

	"github.com/pathfinder-cm/pathfinder-go-client/pfmodel"
)

func TestNewContainerListFromByte(t *testing.T) {
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

	cl, _ := NewContainerListFromByte(b)
	for i, table := range tables {
		if (*cl)[i].Hostname != table.container.Hostname {
			t.Errorf("Incorrect container hostname generated, got: %s, want: %s.",
				(*cl)[i].Hostname,
				table.container.Hostname)
		}

		if (*cl)[i].Source.Type != table.container.Source.Type {
			t.Errorf("Incorrect container source type generated, got: %s, want: %s.",
				(*cl)[i].Source.Type,
				table.container.Source.Type)
		}

		if (*cl)[i].Source.Mode != table.container.Source.Mode {
			t.Errorf("Incorrect container source mode generated, got: %s, want: %s.",
				(*cl)[i].Source.Mode,
				table.container.Source.Mode)
		}

		if (*cl)[i].Source.Alias != table.container.Source.Alias {
			t.Errorf("Incorrect container source alias generated, got: %s, want: %s.",
				(*cl)[i].Source.Alias,
				table.container.Source.Alias)
		}

		if (*cl)[i].Source.Remote.Server != table.container.Source.Remote.Server {
			t.Errorf("Incorrect container remote server generated, got: %s, want: %s.",
				(*cl)[i].Source.Remote.Server,
				table.container.Source.Remote.Server)
		}

		if (*cl)[i].Source.Remote.Protocol != table.container.Source.Remote.Protocol {
			t.Errorf("Incorrect container remote protocol generated, got: %s, want: %s.",
				(*cl)[i].Source.Remote.Protocol,
				table.container.Source.Remote.Protocol)
		}

		if (*cl)[i].Source.Remote.AuthType != table.container.Source.Remote.AuthType {
			t.Errorf("Incorrect container remote auth_type generated, got: %s, want: %s.",
				(*cl)[i].Source.Remote.AuthType,
				table.container.Source.Remote.AuthType)
		}

		if (*cl)[i].Source.Remote.Certificate != table.container.Source.Remote.Certificate {
			t.Errorf("Incorrect container remote certificate generated, got: %s, want: %s.",
				(*cl)[i].Source.Remote.Certificate,
				table.container.Source.Remote.Certificate)
		}

		if (*cl)[i].Status != table.container.Status {
			t.Errorf("Incorrect container status generated, got: %s, want: %s.",
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

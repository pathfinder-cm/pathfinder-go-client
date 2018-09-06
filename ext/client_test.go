package ext

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchContainersFromServer(t *testing.T) {
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

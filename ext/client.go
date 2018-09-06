package ext

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pathfinder-cm/pathfinder-go-client/pfmodel"
	log "github.com/sirupsen/logrus"
)

type Client interface {
	GetNodes() (*pfmodel.NodeList, error)
}

type client struct {
	cluster      string
	token        string
	httpClient   *http.Client
	pfServerAddr string
	pfApiPath    map[string]string
}

func NewClient(
	cluster string,
	token string,
	httpClient *http.Client,
	pfServerAddr string,
	pfApiPath map[string]string) Client {

	return &client{
		cluster:      cluster,
		token:        token,
		httpClient:   httpClient,
		pfServerAddr: pfServerAddr,
		pfApiPath:    pfApiPath,
	}
}

func (c *client) GetNodes() (*pfmodel.NodeList, error) {
	addr := fmt.Sprintf("%s/%s", c.pfServerAddr, c.pfApiPath["GetNodes"])
	u, err := url.Parse(addr)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	q := u.Query()
	q.Set("cluster_name", c.cluster)
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	req.Header.Set("X-Auth-Token", c.token)
	res, err := c.httpClient.Do(req)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(res.Body)
		log.Error(string(b))
		return nil, errors.New(string(b))
	}

	b, _ := ioutil.ReadAll(res.Body)
	nodes, err := NewNodeListFromByte(b)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return nodes, nil
}

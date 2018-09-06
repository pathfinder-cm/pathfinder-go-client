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
	GetNode(string) (*pfmodel.Node, error)
	GetContainers() (*pfmodel.ContainerList, error)
	GetContainer(string) (*pfmodel.Container, error)
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

func (c *client) GetNode(nodeHostname string) (*pfmodel.Node, error) {
	addr := fmt.Sprintf("%s/%s/%s", c.pfServerAddr, c.pfApiPath["GetNode"], nodeHostname)
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
	node, err := NewNodeFromByte(b)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return node, nil
}

func (c *client) GetContainers() (*pfmodel.ContainerList, error) {
	addr := fmt.Sprintf("%s/%s", c.pfServerAddr, c.pfApiPath["GetContainers"])
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
	nodes, err := NewContainerListFromByte(b)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return nodes, nil
}

func (c *client) GetContainer(containerHostname string) (*pfmodel.Container, error) {
	addr := fmt.Sprintf("%s/%s/%s", c.pfServerAddr, c.pfApiPath["GetContainer"], containerHostname)
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
	container, err := NewContainerFromByte(b)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return container, nil
}

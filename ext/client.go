package ext

import (
	"bytes"
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
	CreateContainer(pfmodel.Container) (*pfmodel.Container, error)
	DeleteContainer(string) (*pfmodel.Container, error)
	RescheduleContainer(string) (*pfmodel.Container, error)
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
	containers, err := NewContainerListFromByte(b)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return containers, nil
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

func (c *client) CreateContainer(cntr pfmodel.Container) (*pfmodel.Container, error) {
	addr := fmt.Sprintf("%s/%s", c.pfServerAddr, c.pfApiPath["CreateContainer"])
	u, err := url.Parse(addr)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	q := u.Query()
	q.Set("cluster_name", c.cluster)
	u.RawQuery = q.Encode()

	form := url.Values{}
	form.Set("container[hostname]", cntr.Hostname)
	form.Set("container[source][source_type]", cntr.Source.Type)
	form.Set("container[source][alias]", cntr.Source.Alias)
	form.Set("container[source][mode]", cntr.Source.Mode)
	form.Set("container[source][remote][server]", cntr.Source.Remote.Server)
	form.Set("container[source][remote][protocol]", cntr.Source.Remote.Protocol)
	form.Set("container[source][remote][certificate]", cntr.Source.Remote.Certificate)
	body := bytes.NewBufferString(form.Encode())

	req, err := http.NewRequest(http.MethodPost, u.String(), body)
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
	cntrRes, err := NewContainerFromByte(b)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return cntrRes, nil
}

func (c *client) DeleteContainer(hostname string) (*pfmodel.Container, error) {
	addr := fmt.Sprintf("%s/%s/%s/%s", c.pfServerAddr, c.pfApiPath["DeleteContainer"], hostname, "schedule_deletion")
	u, err := url.Parse(addr)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	q := u.Query()
	q.Set("cluster_name", c.cluster)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
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

func (c *client) RescheduleContainer(hostname string) (*pfmodel.Container, error) {
	addr := fmt.Sprintf("%s/%s/%s/%s", c.pfServerAddr, c.pfApiPath["RescheduleContainer"], hostname, "reschedule")
	u, err := url.Parse(addr)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	q := u.Query()
	q.Set("cluster_name", c.cluster)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
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

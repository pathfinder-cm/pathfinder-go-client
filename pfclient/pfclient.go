package pfclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pathfinder-cm/pathfinder-go-client/pfmodel"
	log "github.com/sirupsen/logrus"
)

type Pfclient interface {
	Register(node string, ipaddress string) (bool, error)
	FetchScheduledContainersFromServer(node string) (*pfmodel.ContainerList, error)
	FetchProvisionedContainersFromServer(node string) (*pfmodel.ContainerList, error)
	UpdateIpaddress(node string, hostname string, ipaddress string) (bool, error)
	MarkContainerAsProvisioned(node string, hostname string) (bool, error)
	MarkContainerAsProvisionError(node string, hostname string) (bool, error)
	MarkContainerAsBootstrapped(node string, hostname string) (bool, error)
	MarkContainerAsBootstrapError(node string, hostname string) (bool, error)
	MarkContainerAsDeleted(node string, hostname string) (bool, error)
	StoreMetrics(collectedMetrics *pfmodel.Metrics) (bool, error)
}

type pfclient struct {
	cluster         string
	clusterPassword string
	token           string
	httpClient      *http.Client
	pfServerAddr    string
	pfApiPath       map[string]string
}

func NewPfclient(
	cluster string,
	clusterPassword string,
	httpClient *http.Client,
	pfServerAddr string,
	pfApiPath map[string]string) Pfclient {

	return &pfclient{
		cluster:         cluster,
		clusterPassword: clusterPassword,
		httpClient:      httpClient,
		pfServerAddr:    pfServerAddr,
		pfApiPath:       pfApiPath,
	}
}

func (p *pfclient) Register(node string, ipaddress string) (bool, error) {
	addr := fmt.Sprintf("%s/%s", p.pfServerAddr, p.pfApiPath["Register"])
	u, err := url.Parse(addr)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}
	q := u.Query()
	q.Set("cluster_name", p.cluster)
	q.Set("node_hostname", node)
	q.Set("node_ipaddress", ipaddress)
	u.RawQuery = q.Encode()

	form := url.Values{}
	form.Set("password", p.clusterPassword)
	body := bytes.NewBufferString(form.Encode())

	req, err := http.NewRequest(http.MethodPost, u.String(), body)
	res, err := p.httpClient.Do(req)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(res.Body)
		log.Error(string(b))
		return false, errors.New(string(b))
	}

	b, _ := ioutil.ReadAll(res.Body)
	register, err := NewRegisterFromByte(b)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}

	p.token = register.AuthenticationToken
	return true, nil
}

func (p *pfclient) FetchScheduledContainersFromServer(node string) (*pfmodel.ContainerList, error) {
	addr := fmt.Sprintf("%s/%s", p.pfServerAddr, p.pfApiPath["ListScheduledContainers"])
	u, err := url.Parse(addr)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	q := u.Query()
	q.Set("cluster_name", p.cluster)
	q.Set("node_hostname", node)
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	req.Header.Set("X-Auth-Token", p.token)
	res, err := p.httpClient.Do(req)
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
	serverContainers, err := NewContainerListFromByte(b)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return serverContainers, nil
}

func (p *pfclient) FetchProvisionedContainersFromServer(node string) (*pfmodel.ContainerList, error) {
	addr := fmt.Sprintf("%s/%s", p.pfServerAddr, p.pfApiPath["ListProvisionedContainers"])
	u, err := url.Parse(addr)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	q := u.Query()
	q.Set("cluster_name", p.cluster)
	q.Set("node_hostname", node)
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	req.Header.Set("X-Auth-Token", p.token)
	res, err := p.httpClient.Do(req)
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
	serverContainers, err := NewContainerListFromByte(b)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return serverContainers, nil
}

func (p *pfclient) UpdateIpaddress(node string, hostname string, ipaddress string) (bool, error) {
	addr := fmt.Sprintf("%s/%s", p.pfServerAddr, p.pfApiPath["UpdateIpaddress"])
	u, err := url.Parse(addr)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}
	q := u.Query()
	q.Set("cluster_name", p.cluster)
	q.Set("node_hostname", node)
	q.Set("hostname", hostname)
	u.RawQuery = q.Encode()

	form := url.Values{}
	form.Set("ipaddress", ipaddress)
	body := bytes.NewBufferString(form.Encode())

	req, err := http.NewRequest(http.MethodPost, u.String(), body)
	req.Header.Set("X-Auth-Token", p.token)
	res, err := p.httpClient.Do(req)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(res.Body)
		log.Error(string(b))
		return false, errors.New(string(b))
	}

	return true, nil
}

func (p *pfclient) MarkContainerAsProvisioned(node string, hostname string) (bool, error) {
	addr := fmt.Sprintf("%s/%s", p.pfServerAddr, p.pfApiPath["MarkProvisioned"])
	u, err := url.Parse(addr)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}
	q := u.Query()
	q.Set("cluster_name", p.cluster)
	q.Set("node_hostname", node)
	q.Set("hostname", hostname)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	req.Header.Set("X-Auth-Token", p.token)
	res, err := p.httpClient.Do(req)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(res.Body)
		log.Error(string(b))
		return false, errors.New(string(b))
	}

	return true, nil
}

func (p *pfclient) MarkContainerAsProvisionError(node string, hostname string) (bool, error) {
	addr := fmt.Sprintf("%s/%s", p.pfServerAddr, p.pfApiPath["MarkProvisionError"])
	u, err := url.Parse(addr)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}
	q := u.Query()
	q.Set("cluster_name", p.cluster)
	q.Set("node_hostname", node)
	q.Set("hostname", hostname)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	req.Header.Set("X-Auth-Token", p.token)
	res, err := p.httpClient.Do(req)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(res.Body)
		log.Error(string(b))
		return false, errors.New(string(b))
	}

	return true, nil
}

func (p *pfclient) MarkContainerAsBootstrapped(node string, hostname string) (bool, error) {
	addr := fmt.Sprintf("%s/%s", p.pfServerAddr, p.pfApiPath["MarkBootstrapped"])
	u, err := url.Parse(addr)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}
	q := u.Query()
	q.Set("cluster_name", p.cluster)
	q.Set("node_hostname", node)
	q.Set("hostname", hostname)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	req.Header.Set("X-Auth-Token", p.token)
	res, err := p.httpClient.Do(req)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(res.Body)
		log.Error(string(b))
		return false, errors.New(string(b))
	}

	return true, nil
}

func (p *pfclient) MarkContainerAsBootstrapError(node string, hostname string) (bool, error) {
	addr := fmt.Sprintf("%s/%s", p.pfServerAddr, p.pfApiPath["MarkBootstrapError"])
	u, err := url.Parse(addr)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}
	q := u.Query()
	q.Set("cluster_name", p.cluster)
	q.Set("node_hostname", node)
	q.Set("hostname", hostname)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	req.Header.Set("X-Auth-Token", p.token)
	res, err := p.httpClient.Do(req)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(res.Body)
		log.Error(string(b))
		return false, errors.New(string(b))
	}

	return true, nil
}

func (p *pfclient) MarkContainerAsDeleted(node string, hostname string) (bool, error) {
	addr := fmt.Sprintf("%s/%s", p.pfServerAddr, p.pfApiPath["MarkDeleted"])
	u, err := url.Parse(addr)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}
	q := u.Query()
	q.Set("cluster_name", p.cluster)
	q.Set("node_hostname", node)
	q.Set("hostname", hostname)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	req.Header.Set("X-Auth-Token", p.token)
	res, err := p.httpClient.Do(req)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(res.Body)
		log.Error(string(b))
		return false, errors.New(string(b))
	}

	return true, nil
}

func (p *pfclient) StoreMetrics(metrics *pfmodel.Metrics) (bool, error) {
	// Setup address and query params
	addr := fmt.Sprintf("%s/%s", p.pfServerAddr, p.pfApiPath["StoreMetrics"])
	u, err := url.Parse(addr)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}
	q := u.Query()
	q.Set("cluster_name", p.cluster)
	u.RawQuery = q.Encode()

	// Setup request body
	b, err := json.Marshal(metrics)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}

	// Create the request and execute it
	req, _ := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(b))
	req.Header.Set("X-Auth-Token", p.token)
	req.Header.Set("Content-Type", "application/json")
	res, err := p.httpClient.Do(req)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(res.Body)
		log.Error(string(b))
		return false, errors.New(string(b))
	}

	return true, nil
}

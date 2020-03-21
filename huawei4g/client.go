package huawei4g

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	baseURL string
}

type TrafficStatistics struct {
	CurrentConnectTime  int
	CurrentUpload       int
	CurrentUploadRate   int
	CurrentDownload     int
	CurrentDownloadRate int
	TotalUpload         int
	TotalDownload       int
	TotalConnectTime    int
}

func (c Client) GetTrafficStatistics() (*TrafficStatistics, error) {
	res, err := http.Get(fmt.Sprintf(c.baseURL + "/api/monitoring/traffic-statistics"))

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get traffic stats: %q", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to get traffic stats: %v", err)
	}

	var stats TrafficStatistics

	err = xml.Unmarshal(body, &stats)
	if err != nil {
		return nil, fmt.Errorf("failed to get traffic stats: %v", err)
	}

	return &stats, nil
}

func New(ip string) *Client {
	return &Client{
		baseURL: fmt.Sprintf("http://%s", ip),
	}
}

package source

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
)

// Github describes github repo of gmdata
type Github struct {
	repoURL string
}

// NewGithub creates a new Github
func NewGithub() *Github {
	if os.Getenv("GITHUB_URL") == "" {
		logger.Sugar.Fatal(`Environment "GITHUB_URL" must be set`)
	}
	return &Github{
		repoURL: os.Getenv("GITHUB_URL"),
	}
}

// Fetch is to get details from github
func (g *Github) fetch(fileName string) ([]byte, error) {
	fileURL := fmt.Sprintf("%s/%s", g.repoURL, fileName)
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, fileURL, nil)
	if err != nil {
		logger.Sugar.Errorf("failed to fetch %s: %v", fileName, err)
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.VERSION.raw")

	resp, err := client.Do(req)
	if err != nil {
		logger.Sugar.Errorf("failed to fetch file[%s]: %v", fileName, err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Sugar.Errorf("failed to read %s: %v", fileName, err)
		return nil, err
	}
	defer resp.Body.Close()

	if len(body) == 0 {
		logger.Sugar.Errorf("file %s is empty", fileName)
		return nil, module.ErrInvalidConfigJSON
	}

	if !json.Valid(body) {
		return nil, module.ErrInvalidConfigJSON
	}
	return body, nil
}

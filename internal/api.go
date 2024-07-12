package downloader

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

// API is the shutterstock-downloader API.
type API struct {
	logger Logger
	config *Config
	client *resty.Client
}

func (a *API) GetLicenses(page string) *LicenseResponse {
	licenseResponse := LicenseResponse{}
	resp, err := a.client.R().
		SetQueryParams(map[string]string{
			"per_page": "200",
			"page":     page,
		}).
		SetResult(&licenseResponse).
		Get(a.config.apiURL + "images/licenses")
	if err != nil {
		a.logger.Errorf("error: %v", err)
		os.Exit(1)
	}
	if resp.StatusCode() != http.StatusOK {
		a.logger.Errorf("invalid status code: %d", resp.StatusCode())
		os.Exit(1)
	}

	return &licenseResponse
}

// GetDownloadLink returns the download link.
func (a *API) GetDownloadLink(id string) (*DownloadLink, error) {
	downloadLink := DownloadLink{}

	resp, err := a.client.R().
		SetResult(&downloadLink).
		Post(a.config.apiURL + "images/licenses/" + id + "/downloads")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		a.logger.Errorf("invalid status code: %d", resp.StatusCode())

		return nil, errors.New("invalid status code")
	}

	return &downloadLink, nil
}

// DownloadImage downloads the image.
func (a *API) DownloadImage(id, url string) error {
	resp, err := a.client.R().SetOutput(a.config.imagesDir + id + ".jpg").Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New("invalid status code")
	}

	return nil
}

// NewAPI returns the API object.
func NewAPI(logger Logger, config *Config) *API {
	client := resty.New()
	client.SetRetryCount(config.retryCount).AddRetryCondition(
		func(r *resty.Response, _ error) bool {
			if r.StatusCode() == http.StatusTooManyRequests {
				dur := time.Minute * time.Duration(config.retryTimeout)
				logger.Errorf("Start sleeping for %v", dur)
				time.Sleep(dur)
				logger.Infof("Stop sleeping for %v", dur)

				return true
			}

			return r.StatusCode() != http.StatusOK
		},
	).SetDebug(config.isDebug).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetAuthToken(config.token)

	return &API{
		client: client,
		config: config,
		logger: logger,
	}
}

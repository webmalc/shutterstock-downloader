package downloader

import (
	"strconv"
	"time"
)

// Downloader downloads images.
type Downloader struct {
	logger Logger
	config *Config
	api    *API
	csv    *CSV
}

// Run downloads images.
func (d *Downloader) Run() {
	licenses := d.getLicenseIDs()
	c := make(chan DownloadedImage)
	result := make([]DownloadedImage, 0)
	for _, l := range licenses {
		time.Sleep(time.Second * 90)
		go d.download(l, c)
	}
	for range licenses {
		result = append(result, <-c)
	}
	err := d.csv.Write(result)
	if err != nil {
		d.logger.Errorf("error: %v", err)
	}
}

func (d *Downloader) getLicenseIDs() []License {
	licenses := make([]License, 0)
	page := 1
	for {
		r := d.api.GetLicenses(strconv.Itoa(page))
		if len(r.Data) == 0 {
			break
		}
		licenses = append(licenses, r.Data...)

		page++
	}

	return licenses
}

func (d *Downloader) download(license License, c chan DownloadedImage) {
	isDownloadable := false
	link, err := d.api.GetDownloadLink(license.ID)
	if err == nil {
		err = d.api.DownloadImage(license.Image.ID, link.URL)
		if err == nil {
			isDownloadable = true
		}
	}
	c <- DownloadedImage{
		ID:             license.ID,
		ImageID:        license.Image.ID,
		IsDownloadable: isDownloadable,
	}
}

// NewDownloader returns a new downloader object.
func NewDownloader(logger Logger) *Downloader {
	c := NewConfig()
	d := &Downloader{
		logger: logger,
		config: c,
		api:    NewAPI(logger, c),
		csv:    NewCSV(c),
	}

	return d
}

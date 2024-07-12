package downloader

type Image struct {
	ID string
}

// License is the struct for the license.
type License struct {
	IsDownloadable bool `json:"is_downloadable"`
	ID             string
	Image          Image
}

// LicenseResponse is the struct for the license response.
type LicenseResponse struct {
	TotalCount int `json:"total_count"`
	Page       int
	Data       []License
}

// DownloadLink is the struct for the download link.
type DownloadLink struct {
	URL string
}

// DownloadedImage is the struct for the downloaded image.
type DownloadedImage struct {
	ID             string
	ImageID        string
	IsDownloadable bool
}

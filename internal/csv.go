package downloader

import (
	"encoding/csv"
	"os"
	"strconv"
)

// CSV is the struct for the CSV.
type CSV struct {
	config *Config
}

// Write writes the CSV.
func (c *CSV) Write(images []DownloadedImage) error {
	data := [][]string{}
	for _, image := range images {
		data = append(data, []string{
			image.ID, image.ImageID, strconv.FormatBool(image.IsDownloadable),
		})
	}
	file, err := os.Create(c.config.csvFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.WriteAll(data)

	if err != nil {
		return err
	}

	return nil
}

// NewCSV returns a new CSV object.
func NewCSV(config *Config) *CSV {
	return &CSV{config: config}
}

package downloader

import "github.com/spf13/viper"

// Config is the cmd configuration struct.
type Config struct {
	apiURL       string
	isDebug      bool
	retryCount   int
	retryTimeout int
	token        string
	csvFilename  string
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	config := &Config{
		apiURL:       viper.GetString("api_url"),
		isDebug:      !viper.GetBool("is_prod"),
		retryCount:   viper.GetInt("retry_count"),
		retryTimeout: viper.GetInt("retry_timeout"),
		token:        viper.GetString("token"),
		csvFilename:  viper.GetString("csv_filename"),
	}

	return config
}

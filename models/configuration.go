package models

// Configuration model
type Configuration struct {
	HTTP struct {
		Address string `mapstructure:"address" validate:"required,tcp_addr"`
	} `mapstructure:"http" validate:"required"`
	Encryption struct {
		EncryptResponseKey string `mapstructure:"encrypt_response_key" validate:"required"`
		EncryptCPMKey      string `mapstructure:"encrypt_cpm_key" validate:"required"`
	} `mapstructure:"encryption" validate:"required"`
	MODE     string `mapstructure:"mode" validate:"required"`
	Tracking struct {
		Filename   string `mapstructure:"filename" validate:"required"`
		MaxSize    int    `mapstructure:"max_size" validate:"required"`
		MaxBackups int    `mapstructure:"max_backups" validate:"required"`
	} `mapstructure:"tracking" validate:"required"`
	ThirdPartyTrackingCallback struct {
		Filename   string `mapstructure:"filename" validate:"required"`
		MaxSize    int    `mapstructure:"max_size" validate:"required"`
		MaxBackups int    `mapstructure:"max_backups" validate:"required"`
	} `mapstructure:"third_party_tracking_callback" validate:"required"`
	Logging struct {
		Level string `mapstructure:"level" validate:"required,eq=debug|eq=info|eq=warn|eq=warning|eq=error|eq=fatal|eq=panic"`
	} `mapstructure:"logging" validate:"required"`
	DB struct {
		DataSourceName string `mapstructure:"dsn" validate:"required,dsn"`
		MaxOpenConns   int    `mapstructure:"max_open_conns" validate:"required,min=1"`
		MaxIdleConns   int    `mapstructure:"max_idle_conns" validate:"required,min=1,ltefield=MaxOpenConns"`
	} `mapstructure:"db" validate:"required"`
	Redis struct {
		URL      string `mapstructure:"url" validate:"required,redis_url"`
		PoolSize int    `mapstructure:"pool_size" validate:"required,min=1"`
	} `mapstructure:"redis"`
	CVRPrediction struct {
		ABTestID      string `mapstructure:"abtest_id"`
		ABTestRatio   int64  `mapstructure:"abtest_ratio"`
		GitCommit     string `mapstructure:"git_commit" validate:"omitempty,len=7"`
		RedisURL      string `mapstructure:"redis_url" validate:"omitempty,redis_url"`
		RedisPoolSize int    `mapstructure:"redis_pool_size" validate:"omitempty,min=1"`
	} `mapstructure:"cvr_prediction"`
	Locator struct {
		IPIPFile        string `mapstructure:"ipip_file"`
		CountryCodeFile string `mapstructure:"country_code_file"`
	} `mapstructure:"locator"`
	Iacip struct {
		AreaFile   string `mapstructure:"area_file"`
		IPDataFile string `mapstructure:"ip_data_file"`
	} `mapstructure:"iacip"`
	Statistics struct {
		Endpoint string `mapstructure:"endpoint" validate:"required,url"`
		Overview string `mapstructure:"overview" validate:"required,url"`
	} `mapstructure:"statistics" validate:"required"`
	RPC struct {
		Address               string `mapstructure:"address" validate:"required,url"`
		MaxIdleConnsPerHost   int    `mapstructure:"max_idle_conns_per_host" validate:"required,min=1"`
		MaxConcurrentRequests int    `mapstructure:"max_concurrent_requests" validate:"required,min=1"`
	} `mapstructure:"rpc" validate:"required"`
	Prophet struct {
		ABTestID string `mapstructure:"abtest_id"`
		URL      string `mapstructure:"url"`
	}
	Report struct {
		URL string `mapstructure:"url"`
	}
	RequestInterval map[string]int `mapstructure:"request_interval"`
}

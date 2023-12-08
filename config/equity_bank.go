package config

type EquityBank struct {
	ConsumerSecret string `json:"consumerSecret"`
	ApiKey         string `json:"apiKey"`
	MerchantCode   string `json:"merchantCode"`
	BaseApi        string `json:"baseApi"`
	PrivateKey     string `json:"privateKey"`
}

package config

type MpesaConfig struct {
	ConsumerKey    string `json:"mpesaConsumerKey"`
	ConsumerSecret string `json:"mpesaConsumerSecret"`
	BaseApi        string `json:"baseApi"`
}

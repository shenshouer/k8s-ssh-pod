package config

type Config struct {
	KubeConfig string `json:"kubeconfig"`
	ServeAddr  string `json:"addr"`
	PrivateKey string `json:"key"`
}

func GetConfig() *Config {
	return &Config{
		KubeConfig: `/Users/shenshouer/.kube/config`,
		ServeAddr:  ":2222",
		PrivateKey: "./host_key",
	}
}

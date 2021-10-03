package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

//Config ...
type Config struct {
	App struct {
		Name        string `yaml:"name"`
		Port        int    `yaml:"port"`
		Host        string `yaml:"host"`
		NgrokTunnel string `yaml:"ngrok_tunnel"`
		DBPath      string `yaml:"db_path"`
	} `yaml:"app"`

	OctyAPIURIs struct {
		CreateProfile          string `yaml:"create_profile"`
		UpdateProfile          string `yaml:"update_profile"`
		GetProfiles            string `yaml:"get_profiles"`
		IdentifyProfiles       string `yaml:"identify_profiles"`
		CreateEvent            string `yaml:"create_event"`
		CreateItem             string `yaml:"create_item"`
		UpdateItem             string `yaml:"update_item"`
		DeleteItem             string `yaml:"delete_item"`
		GetTemplates           string `yaml:"get_templates"`
		GenerateContent        string `yaml:"generate_content"`
		PredictRecommendations string `yaml:"predict_recommendations"`
	} `yaml:"octy_api_uris"`

	OctyCreds struct {
		PublicKey string `yaml:"public_key"`
		SecretKey string `yaml:"secret_key"`
	} `yaml:"octy_creds"`

	Shopify struct {
		StoreRootURL   string `yaml:"store_root_url"`
		GetCustomerURI string `yaml:"get_customer_uri"`
		GetProductsURI string `yaml:"get_products_uri"`
		APIKey         string `yaml:"api_key"`
		APISecret      string `yaml:"api_secret"`
		WebhookSecret  string `yaml:"webhook_secret"`
	} `yaml:"shopify"`
}

//NewConfig ...
func NewConfig(configFile string) (*Config, error) {
	abs, err := filepath.Abs(configFile) // get absolute path of config file
	if err != nil {
		return nil, err
	}
	file, err := os.Open(abs)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	cfg := &Config{}
	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}

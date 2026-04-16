package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config holds all application configuration loaded from environment variables.
type Config struct {
	// Server
	Port string

	// Database
	DatabaseURL string

	// Auth
	JWTSecret       string
	SIWEDomain      string
	AdminWallet     string
	AdminPrivateKey string

	// Blockchain
	AlchemyBaseSepolia string
	AlchemyBaseMainnet string
	ChainID            string

	// Contract addresses
	AdminMultisigAddress      string
	TimelockAddress           string
	ReputationRegistryAddress string
	DonationNFTAddress        string
	CampaignFactoryAddress    string

	// IPFS / Pinata
	PinataJWT       string
	PinataAPIKey    string
	PinataAPISecret string

	// The Graph
	GraphURL         string
	GraphAccessToken string

	// Email / Resend
	ResendAPIKey string
	EmailFrom    string

	// Logging
	LogLevel string
}

// Load reads environment variables (real env first, then .env file fallback).
func Load() (*Config, error) {
	// Load .env file if it exists — godotenv correctly handles comments & spaces.
	// We look for .env in the working directory, then one level up (e.g. when
	// running `go run ./cmd/server` from within the project root).
	envFiles := []string{".env", "../.env"}
	for _, f := range envFiles {
		if _, err := os.Stat(f); err == nil {
			// Don't overwrite vars that are already set in the real environment.
			if err := godotenv.Load(f); err != nil {
				return nil, fmt.Errorf("config: loading %s: %w", f, err)
			}
			break
		}
	}

	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Defaults
	v.SetDefault("PORT", "8080")
	v.SetDefault("LOG_LEVEL", "info")
	v.SetDefault("SIWE_DOMAIN", "localhost")
	v.SetDefault("CHAIN_ID", "84532")

	cfg := &Config{
		Port:        v.GetString("PORT"),
		DatabaseURL: v.GetString("DATABASE_URL"),

		JWTSecret:       v.GetString("JWT_SECRET"),
		SIWEDomain:      v.GetString("SIWE_DOMAIN"),
		AdminWallet:     strings.ToLower(v.GetString("ADMIN_WALLET")),
		AdminPrivateKey: v.GetString("ADMIN_PRIVATE_KEY"),

		AlchemyBaseSepolia: v.GetString("ALCHEMY_BASE_SEPOLIA_URL"),
		AlchemyBaseMainnet: v.GetString("ALCHEMY_BASE_MAINNET_URL"),
		ChainID:            v.GetString("CHAIN_ID"),

		AdminMultisigAddress:      v.GetString("ADMIN_MULTISIG_ADDRESS"),
		TimelockAddress:           v.GetString("TIMELOCK_ADDRESS"),
		ReputationRegistryAddress: v.GetString("REPUTATION_REGISTRY_ADDRESS"),
		DonationNFTAddress:        v.GetString("DONATION_NFT_ADDRESS"),
		CampaignFactoryAddress:    v.GetString("CAMPAIGN_FACTORY_ADDRESS"),

		PinataJWT:       v.GetString("PINATA_JWT"),
		PinataAPIKey:    v.GetString("PINATA_API_KEY"),
		PinataAPISecret: v.GetString("PINATA_API_SECRET"),

		GraphURL:         v.GetString("GRAPH_URL"),
		GraphAccessToken: v.GetString("GRAPH_ACCESS_TOKEN"),

		ResendAPIKey: v.GetString("RESEND_API_KEY"),
		EmailFrom:    v.GetString("EMAIL_FROM"),

		LogLevel: v.GetString("LOG_LEVEL"),
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// validate checks that all required fields are present.
func (c *Config) validate() error {
	required := map[string]string{
		"DATABASE_URL": c.DatabaseURL,
		"JWT_SECRET":   c.JWTSecret,
	}
	for key, val := range required {
		if val == "" {
			return fmt.Errorf("config: required env variable %s is not set", key)
		}
	}
	return nil
}

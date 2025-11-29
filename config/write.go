package config

import (
	"fmt"
	"path/filepath"

	cmtconfig "github.com/cometbft/cometbft/config"
)

// WriteConfig writes the cometbft config.toml and sonata.toml files to the config directory
func WriteConfig(configDir string, cfg *Config) error {
	// write cometbft config.toml
	cmtconfig.WriteConfigFile(filepath.Join(configDir, "config.toml"), cfg.CometBFT)

	// write sonata.toml
	if err := cfg.Sonata.SaveAs(filepath.Join(configDir, "sonata.toml")); err != nil {
		return fmt.Errorf("writing sonata.toml: %w", err)
	}
	return nil
}

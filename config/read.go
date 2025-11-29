package config

import (
	"fmt"
	"path/filepath"

	cmtconfig "github.com/cometbft/cometbft/config"
	"github.com/spf13/viper"
)

func ReadConfig(homeDir string) (*Config, error) {
	if homeDir == "" {
		homeDir = DefaultHomeDirPath()
	}

	// Read Sonata config
	sonataConfig, err := ReadSonataConfig(homeDir)
	if err != nil {
		return nil, fmt.Errorf("reading sonata config: %w", err)
	}

	// Read CometBFT config
	cmtConfig, err := ReadCometBFTConfig(homeDir)
	if err != nil {
		return nil, fmt.Errorf("reading cometbft config: %w", err)
	}

	return &Config{
		Sonata:   sonataConfig,
		CometBFT: cmtConfig,
	}, nil
}

func ReadSonataConfig(homeDir string) (*SonataConfig, error) {
	sonataConfig := DefaultSonataConfig()
	sonataConfig.SetRoot(homeDir)

	v := viper.New()
	v.SetConfigFile(filepath.Join(homeDir, "config", "sonata.toml"))

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}
	if err := v.Unmarshal(sonataConfig); err != nil {
		return nil, fmt.Errorf("unmarshaling config: %w", err)
	}
	if err := sonataConfig.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("validating config: %w", err)
	}

	return sonataConfig, nil
}

func ReadCometBFTConfig(homeDir string) (*cmtconfig.Config, error) {
	cmtConfig := cmtconfig.DefaultConfig()
	cmtConfig.SetRoot(homeDir)

	v := viper.New()
	v.SetConfigFile(filepath.Join(homeDir, "config", "config.toml"))

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}
	if err := v.Unmarshal(cmtConfig); err != nil {
		return nil, fmt.Errorf("unmarshaling config: %w", err)
	}
	if err := cmtConfig.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("validating config: %w", err)
	}

	return cmtConfig, nil
}

package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Username                 string `yaml:"username,omitempty"`
	Password                 string `yaml:"password,omitempty"`
	DownloadsDirectory       string `yaml:"downloads_directory,omitempty"`
	Quality                  string `yaml:"quality,omitempty"`
	SortByContext            bool   `yaml:"sort_by_context,omitempty"`
	CoverSize                string `yaml:"cover_size,omitempty"`
	KeepCover                bool   `yaml:"keep_cover,omitempty"`
	FixTags                  bool   `yaml:"fix_tags,omitempty"`
	TrackFileTemplate        string `yaml:"track_file_template,omitempty"`
	ReleaseDirectoryTemplate string `yaml:"release_directory_template,omitempty"`
	WhitespaceCharacter      string `yaml:"whitespace_character,omitempty"`
	Proxy                    string `yaml:"proxy,omitempty"`
}

const (
	DefaultCoverSize = "1400x1400"
)

func Parse(filePath string) (*AppConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	config := AppConfig{
		Quality:                  "lossless",
		CoverSize:                DefaultCoverSize,
		TrackFileTemplate:        "{number}. {artists} - {name} ({mix_name})",
		ReleaseDirectoryTemplate: "[{catalog_number}] {artists} - {name}",
		FixTags:                  true,
	}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	if config.Username == "" || config.Password == "" {
		return nil, fmt.Errorf("username or password is not provided")
	}

	if config.DownloadsDirectory == "" {
		return nil, fmt.Errorf("no downloads directory provided")
	}

	return &config, nil
}

func (c *AppConfig) Save(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer file.Close()
	encoder := yaml.NewEncoder(file)
	if err := encoder.Encode(&c); err != nil {
		return fmt.Errorf("encode config: %w", err)
	}
	return nil
}

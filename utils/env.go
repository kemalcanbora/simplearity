package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type EnvSettings struct {
	AppDir       string
	IgnoreFile   string
	YamlFile     string
	Dockerfile   string
	AppImageName string
	MaxAppImages int
	Singularity  bool
	JobName      string
	CpusPerTask  int
	Mem          string
	Partition    string
}

func LoadEnv(filePath string) error {
	err := godotenv.Load(filePath)
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

func LoadConfig() (*EnvSettings, error) {
	config := &EnvSettings{
		AppDir:       os.Getenv("APP_DIR"),
		IgnoreFile:   os.Getenv("IGNORE_FILE"),
		YamlFile:     os.Getenv("YAML_FILE"),
		Dockerfile:   os.Getenv("DOCKERFILE"),
		AppImageName: os.Getenv("APP_IMAGE_NAME"),
		JobName:      os.Getenv("JOB_NAME"),
		Mem:          os.Getenv("MEM"),
		Partition:    os.Getenv("PARTITION"),
	}

	maxAppImages, err := strconv.Atoi(os.Getenv("MAX_APP_IMAGES"))
	if err != nil {
		return nil, fmt.Errorf("invalid MAX_APP_IMAGES: %w", err)
	}
	config.MaxAppImages = maxAppImages

	singularity, err := strconv.ParseBool(os.Getenv("SINGULARITY"))
	if err != nil {
		return nil, fmt.Errorf("invalid SINGULARITY: %w", err)
	}
	config.Singularity = singularity

	cpusPerTask, err := strconv.Atoi(os.Getenv("CPUS_PER_TASK"))
	if err != nil {
		return nil, fmt.Errorf("invalid CPUS_PER_TASK: %w", err)
	}
	config.CpusPerTask = cpusPerTask

	return config, nil
}

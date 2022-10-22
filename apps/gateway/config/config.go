package config

import (
	"errors"
	"os"
)

type Config struct {
	GatewayURL string
	AuthURL    string
	CourseURL  string
	ExamURL    string
	ThreadURL  string
	Deployment string
}

func LoadConfig() (*Config, error) {
	GatewayURL := os.Getenv("GATEWAY_URL")
	if GatewayURL == "" {
		return nil, errors.New("environment variable not set")
	}
	AuthURL := os.Getenv("AUTH_URL")
	if AuthURL == "" {
		return nil, errors.New("environment variable not set")
	}
	CourseURL := os.Getenv("COURSE_URL")
	if CourseURL == "" {
		return nil, errors.New("environment variable not set")
	}
	ExamURL := os.Getenv("EXAM_URL")
	if ExamURL == "" {
		return nil, errors.New("environment variable not set")
	}
	ThreadURL := os.Getenv("EXAM_URL")
	if ThreadURL == "" {
		return nil, errors.New("environment variable not set")
	}
	Deployment := os.Getenv("DEPLOYMENT")
	if Deployment == "" {
		return nil, errors.New("environment variable not set")
	}

	return &Config{
		GatewayURL, AuthURL, CourseURL, ExamURL, ThreadURL, Deployment,
	}, nil
}

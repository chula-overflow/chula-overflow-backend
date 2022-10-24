package config

import (
	"errors"
	"os"
)

type Config struct {
	APIBaseURL string
	GatewayURL string
	AuthURL    string
	CourseURL  string
	Deployment string
}

func LoadConfig() (*Config, error) {
	APIBaseURL := os.Getenv("API_BASE_URL")
	if APIBaseURL == "" {
		return nil, errors.New("enrionment variable not set")
	}
	GatewayHost := os.Getenv("GATEWAY_HOST")
	if GatewayHost == "" {
		return nil, errors.New("enrionment variable not set")
	}
	GatewayPort := os.Getenv("GATEWAY_PORT")
	if GatewayPort == "" {
		return nil, errors.New("enrionment variable not set")
	}
	AuthHost := os.Getenv("AUTH_HOST")
	if AuthHost == "" {
		return nil, errors.New("enrionment variable not set")
	}
	AuthPort := os.Getenv("AUTH_PORT")
	if AuthPort == "" {
		return nil, errors.New("enrionment variable not set")
	}
	CourseHost := os.Getenv("COURSE_HOST")
	if CourseHost == "" {
		return nil, errors.New("enrionment variable not set")
	}
	CoursePort := os.Getenv("COURSE_PORT")
	if CoursePort == "" {
		return nil, errors.New("enrionment variable not set")
	}
	Deployment := os.Getenv("DEPLOYMENT")
	if Deployment == "" {
		return nil, errors.New("enrionment variable not set")
	}

	GatewayURL := GatewayHost + ":" + GatewayPort
	AuthURL := AuthHost + ":" + AuthPort
	CourseURL := CourseHost + ":" + CoursePort

	return &Config{
		APIBaseURL, GatewayURL, AuthURL, CourseURL, Deployment,
	}, nil
}

package models

import (
	"net/http"
	"regexp"
)

// Config struct
type Config struct {
	Port          int               `yaml:"port"`
	Storage       *Storage          `yaml:"storage"`
	Sessions      map[string]string `yaml:"sessions"`
	MultiUserMode bool              `yaml:"multiUser"`
	RootPath      string            `yaml:"rootPath"`
}

// Storage struct
type Storage struct {
	Driver     string `yaml:"driver"`
	Connection string `yaml:"connection"`
	Name       string `yaml:"name"`
}

type contextKey int

const (
	//ContextLoggerKey key
	ContextLoggerKey contextKey = iota
	//ContextReqIDKey key
	ContextReqIDKey contextKey = iota
	//ContextAuthMockKey key
	ContextAuthMockKey contextKey = iota
	//XRequestID key
	XRequestID = "X-Request-Id"
)

const (
	// CtxAPIVersion key
	CtxAPIVersion contextKey = iota
	// CtxValueOwner key
	CtxValueOwner
	// CtxValueEnvID key
	CtxValueEnvID
	// CtxValueAuth key
	CtxValueAuth
)

// OwnerFromContext returns context value for project owner
func OwnerFromContext(r *http.Request) string {
	owner := r.Context().Value(CtxValueOwner)
	return owner.(string)
}

// IsCodeValid returns true if code has valid value
func IsCodeValid(code string) bool {
	res, _ := regexp.Match(`^([a-z0-9\-_]+)$`, []byte(code))
	return res
}

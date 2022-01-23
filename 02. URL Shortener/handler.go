package main

import (
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	pathsToUrls = nil
	fallback = nil
	return nil
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	yml = nil
	fallback = nil
	return nil, nil
}

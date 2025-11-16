package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (cfg apiConfig) ensureAssetsDir() error {
	if _, err := os.Stat(cfg.assetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.assetsRoot, 0755)
	}
	return nil
}

func extensionType(mediaType string) string {
	parts := strings.Split(mediaType, "/")
	if len(parts) != 2 {
		return ".bin"
	}
	return "." + parts[1]
}

func assetPath(mediaType string) string {
	assetName := make([]byte, 32)
	_, err := rand.Read(assetName)
	if err != nil {
		panic("failed to generate random bytes")
	}
	assetNameString := base64.RawURLEncoding.EncodeToString(assetName)
	ext := extensionType(mediaType)
	return fmt.Sprintf("%s%s", assetNameString, ext)
}

func (cfg apiConfig) getObjectUrl(key string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", cfg.s3Bucket, cfg.s3Region, key)
}

func (cfg apiConfig) assetDiskPath(assetPath string) string {
	return filepath.Join(cfg.assetsRoot, assetPath)
}

func (cfg *apiConfig) assetUrl(assetPath string) string {
	return fmt.Sprintf("http://localhost:%s/assets/%s", cfg.port, assetPath)
}

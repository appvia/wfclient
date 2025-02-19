/**
 * Copyright 2025 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/appvia/wfclient/pkg/authtypes"
	"github.com/appvia/wfclient/pkg/common"
)

const (
	EnvWayfinderServer    = "WAYFINDER_SERVER"
	EnvWayfinderToken     = "WAYFINDER_TOKEN"
	EnvWayfinderWorkspace = "WAYFINDER_WORKSPACE"
)

func IsEphemeralConfig() bool {
	return os.Getenv(EnvWayfinderServer) != "" && os.Getenv(EnvWayfinderToken) != ""
}

// GetConfig returns either the ephemeral configuration from environment variables if provided, or
// the current configured file - creating it if it does not exist.
func GetConfig() (*Config, error) {
	if IsEphemeralConfig() {
		return CreateEphemeralConfiguration(), nil
	}

	return GetOrCreateClientConfiguration()
}

// CreateEphemeralConfiguration creates a fake configuration from the environments
// variables - largely used for CI
func CreateEphemeralConfiguration() *Config {
	name := "default"

	cfg := &Config{}
	cfg.CurrentProfile = name

	server := os.Getenv(EnvWayfinderServer)
	token := os.Getenv(EnvWayfinderToken)

	// @step: determine the token type and place into the right section
	identity := &Identity{}

	switch IsExchangeToken([]byte(token)) {
	case true:
		identity.RefreshToken = token

	default:
		identity.Token = token
	}

	cfg.CreateProfile(name, server)
	cfg.AddAuthInfo(name, &AuthInfo{Identity: identity})
	cfg.AddServer(name, &Server{Endpoint: server})

	workspace := os.Getenv(EnvWayfinderWorkspace)
	if workspace != "" {
		p := cfg.GetProfile(name)
		p.Workspace = workspace
	}

	return cfg
}

// IsExchangeToken checks if the token is an exchange token
func IsExchangeToken(token []byte) bool {
	exch, err := authtypes.IsExchangeToken(token)
	return exch && err == nil
}

// IsAccessToken checks if the token is an access token
func IsAccessToken(token []byte) bool {
	access, err := authtypes.IsAccessToken(token)
	return access && err == nil
}

// GetClientConfigurationPath returns the path to the client config
func GetClientConfigurationPath() string {
	// @step: retrieve the configuration path from env of default path
	path := os.ExpandEnv(os.Getenv(DefaultWayfinderConfigPathEnv))
	if path == "" {
		path = os.ExpandEnv(DefaultWayfinderConfigPath)
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return path
	}

	return abs
}

// GetClientPath returns the base of the client configuration
func GetClientPath() string {
	return filepath.Dir(GetClientConfigurationPath())
}

// GetOrCreateClientConfiguration is responsible for retrieving the client configuration
var GetOrCreateClientConfiguration = func() (*Config, error) {
	path := GetClientConfigurationPath()
	common.LogWithoutContext().WithField("path", path).Debug("using wayfinder configration file")

	// @step: check the file exists else create it
	if found, err := fileExists(path); err != nil {
		return nil, err
	} else if !found {
		// @step: we need to write an empty file for now
		if err := UpdateConfig(NewEmpty(), path); err != nil {
			return nil, err
		}

		return NewEmpty(), nil
	}

	// @step: open the configuration file for reading
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return New(file)
}

// UpdateConfig is responsible for writing the configuration to disk
var UpdateConfig = func(config *Config, path string) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), os.FileMode(0750)); err != nil {
		return err
	}

	return os.WriteFile(path, data, os.FileMode(0640))
}

// fileExists checks if a file exists
func fileExists(filename string) (bool, error) {
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return !info.IsDir(), nil
}

/*
	Copyright 2023 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package config

import (
	"errors"
	"github.com/loopholelabs/acme"
	"github.com/spf13/pflag"
)

var (
	ErrEmailRequired     = errors.New("email is required")
	ErrDirectoryRequired = errors.New("directory is required")
)

const (
	DefaultDisabled = false
)

type Config struct {
	Disabled  bool   `yaml:"disabled"`
	Email     string `yaml:"email"`
	Directory string `yaml:"directory"`
	KID       string `yaml:"kid"`
	HMAC      string `yaml:"hmac"`
}

func New() *Config {
	return &Config{
		Disabled: DefaultDisabled,
	}
}

func (c *Config) Validate() error {
	if !c.Disabled {
		if c.Email == "" {
			return ErrEmailRequired
		}
		if c.Directory == "" {
			return ErrDirectoryRequired
		}
	}

	return nil
}

func (c *Config) RootPersistentFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&c.Disabled, "acme-disabled", DefaultDisabled, "Disable acme")
	flags.StringVar(&c.Email, "acme-email", "", "Email address to use for ACME registration")
	flags.StringVar(&c.Directory, "acme-directory", "", "Directory URL for ACME registration")
	flags.StringVar(&c.KID, "acme-kid", "", "KID to use for ACME registration")
	flags.StringVar(&c.HMAC, "acme-hmac", "", "HMAC to use for ACME registration")
}

func (c *Config) GenerateOptions(logName string) *acme.Options {
	return &acme.Options{
		LogName:   logName,
		Disabled:  c.Disabled,
		Email:     c.Email,
		Directory: c.Directory,
		KID:       c.KID,
		HMAC:      c.HMAC,
	}
}

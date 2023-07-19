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

package acme

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge/resolver"
	"github.com/go-acme/lego/v4/lego"
	legoLog "github.com/go-acme/lego/v4/log"
	"github.com/go-acme/lego/v4/registration"
	acmeLog "github.com/loopholelabs/acme/pkg/logger"
	"github.com/rs/zerolog"
)

var (
	ErrDisabled = errors.New("acme is disabled")
)

type Options struct {
	LogName   string
	Disabled  bool
	Email     string
	Directory string
	KID       string
	HMAC      string
}

// ACME is a wrapper for the acme client
type ACME struct {
	logger  *zerolog.Logger
	options *Options

	client *lego.Client
}

func New(options *Options, logger *zerolog.Logger) (*ACME, error) {
	l := logger.With().Str(options.LogName, "ACME").Logger()
	if options.Disabled {
		l.Warn().Msg("disabled")
		return nil, ErrDisabled
	}

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	legoLog.Logger = (*acmeLog.Logger)(logger)
	user := &User{
		Email: options.Email,
		Key:   privateKey,
	}

	legoConfig := lego.NewConfig(user)
	legoConfig.CADirURL = options.Directory
	legoConfig.Certificate.KeyType = certcrypto.EC256

	client, err := lego.NewClient(legoConfig)
	if err != nil {
		return nil, err
	}

	if options.KID != "" && options.HMAC != "" {
		user.Registration, err = client.Registration.RegisterWithExternalAccountBinding(registration.RegisterEABOptions{
			TermsOfServiceAgreed: true,
			Kid:                  options.KID,
			HmacEncoded:          options.HMAC,
		})
	} else {
		user.Registration, err = client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	}

	return &ACME{
		logger:  &l,
		options: options,
		client:  client,
	}, nil
}

func (e *ACME) Challenge() *resolver.SolverManager {
	return e.client.Challenge
}

func (e *ACME) Certificate() *certificate.Certifier {
	return e.client.Certificate
}

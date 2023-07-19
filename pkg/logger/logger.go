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

package logger

import (
	"github.com/go-acme/lego/v4/log"
	"github.com/rs/zerolog"
)

var _ log.StdLogger = (*Logger)(nil)

type Logger zerolog.Logger

func (l *Logger) Fatal(args ...interface{}) {
	(*zerolog.Logger)(l).Fatal().Msgf("%v", args...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	(*zerolog.Logger)(l).Fatal().Msgf("%v", args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	(*zerolog.Logger)(l).Fatal().Msgf(format, args...)
}

func (l *Logger) Print(args ...interface{}) {
	(*zerolog.Logger)(l).Info().Msgf("%v", args...)
}

func (l *Logger) Println(args ...interface{}) {
	(*zerolog.Logger)(l).Info().Msgf("%v", args...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	(*zerolog.Logger)(l).Info().Msgf(format, args...)
}

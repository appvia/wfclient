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

package common

import (
	"context"

	"github.com/sirupsen/logrus"
)

// Logger represents Wayfinder's standard logging interface
type Logger interface {
	logrus.Ext1FieldLogger
}

// Default to logrus
var defaultLogger Logger = logrus.StandardLogger()

var logProvider = func() Logger {
	return defaultLogger
}

// SetLogProvider allows the current logging provider to be set
func SetLogProvider(lf func() Logger) {
	logProvider = lf
}

type logContexts int

const (
	logContextUser logContexts = iota
)

// WithUser add the username to the context key values, returning a new context
func WithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, logContextUser, user)
}

// LogWithoutContext should ONLY be used where there is no relevant context - you should prefer
// calling Log(ctx) and add contexts in where they are missing to using this.
func LogWithoutContext() Logger {
	return logProvider()
}

// Log gets a logger to use, initialised with any relevant values from the context
func Log(ctx context.Context) Logger {
	log := logProvider()
	if ctx != nil {
		if user, ok := ctx.Value(logContextUser).(string); ok {
			log = log.WithField("user", user)
		}
	}
	return log
}

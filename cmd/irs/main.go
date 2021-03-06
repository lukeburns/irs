// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/moov-io/identity/pkg/logging"
	"github.com/moov-io/irs/pkg/service"
)

func main() {
	env := &service.Environment{
		Logger: logging.NewDefaultLogger().WithKeyValue("app", "irs"),
	}

	env, err := service.NewEnvironment(env)
	if err != nil {
		env.Logger.Fatal().LogError("Error loading up environment.", err)
		os.Exit(1)
	}
	defer env.Shutdown()

	env.Logger.Info().Log("Starting services")
	shutdown := env.RunServers(true)
	defer shutdown()
}

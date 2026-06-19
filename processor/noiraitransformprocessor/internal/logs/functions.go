// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/no1rstack/noirai-otel-collector/processor/noiraitransformprocessor/internal/logs"

import (
	"fmt"

	noiraiFuncs "github.com/no1rstack/noirai-otel-collector/processor/noiraitransformprocessor/ottlfunctions"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"
)

func NoirAILogFunctions() map[string]ottl.Factory[*ottllog.TransformContext] {
	factoryMap := map[string]ottl.Factory[*ottllog.TransformContext]{}
	for _, f := range []ottl.Factory[*ottllog.TransformContext]{
		noiraiFuncs.NewExprFactory(),
		noiraiFuncs.NewGrokParseFactory[*ottllog.TransformContext](),
		noiraiFuncs.NewHexToIntFactory[*ottllog.TransformContext](),
	} {
		factoryMap[f.Name()] = f
	}
	return factoryMap
}

func LogFunctions() map[string]ottl.Factory[*ottllog.TransformContext] {
	logFunctions := ottlfuncs.StandardFuncs[*ottllog.TransformContext]()

	for name, factory := range NoirAILogFunctions() {
		_, exists := logFunctions[name]
		if exists {
			panic(fmt.Sprintf("ottl func %s already exists", name))
		}
		logFunctions[name] = factory
	}

	return logFunctions
}

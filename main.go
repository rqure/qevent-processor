package main

import (
	qmq "github.com/rqure/qmq/src"
)

func main() {
	engine := qmq.NewDefaultEngine(qmq.DefaultEngineConfig{
		NameProvider:               &NameProvider{},
		TransformerProviderFactory: &TransformerProviderFactory{},
		EngineProcessor:            &EngineProcessor{},
	})
	engine.Run()
}

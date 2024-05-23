package main

import qmq "github.com/rqure/qmq/src"

type TransformerProviderFactory struct {
}

func (t *TransformerProviderFactory) Create(components qmq.EngineComponentProvider) qmq.TransformerProvider {
	transformerProvider := qmq.NewDefaultTransformerProvider()
	// transformerProvider.Set("consumer:"+t.AppNameProvider.Get()+":logs", []qmq.Transformer{
	// 	qmq.NewMessageToAnyTransformer(components.WithLogger()),
	// 	NewAnyToLogTransformer(components.WithLogger()),
	// })
	return transformerProvider
}

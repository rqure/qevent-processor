package main

type NameProvider struct{}

func (np *NameProvider) Get() string {
	return "qevent-processor"
}

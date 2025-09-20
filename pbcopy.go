package main

type Copy interface {
	ReadAll() (string, error)
	Write(string) (int, error)
}

package main

import "os/exec"

type DarwinPbCopy struct {
}

func NewDarwingPbCopy() Copy {
	return &DarwinPbCopy{}
}

func (pb *DarwinPbCopy) ReadAll() (string, error) {
	out := exec.Command("pbpaste")
	clip, err := out.Output()
	if err != nil {
		return "", err
	}

	return string(clip), nil
}

func (pb *DarwinPbCopy) Write(data string) (int, error) {
	return 0, nil
}

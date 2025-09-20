package main

type LinuxPbCopy struct {
}

func NewLinuxPbCopy() Copy {
	return &LinuxPbCopy{}
}

func (pb *LinuxPbCopy) ReadAll() (string, error) {
	return "", nil
}

func (pb *LinuxPbCopy) Write(data string) (int, error) {
	return 0, nil
}

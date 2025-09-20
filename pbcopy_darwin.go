//go:build darwin

package pbcopy

import "os/exec"

func readAll() (string, error) {
	out := exec.Command("pbpaste")
	clip, err := out.Output()
	if err != nil {
		return "", err
	}

	return string(clip), nil
}

func write(data []byte) (int, error) {
	return 0, nil
}

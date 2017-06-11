package foobar

import (
	"log"
	"os/exec"
	"testing"
)

func TestAdd(t *testing.T) {
	s := Add(1, 2)
	if s != 3 {
		t.Errorf("exptect 3, got %v\n", s)
	}
}

func TestAdd2(t *testing.T) {
	cmd := exec.Command("socat",
		"pty,raw,echo=0,link=/dev/shm/foo",
		"pty,raw,echo=0,link=/dev/shm/bar")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer cmd.Wait()
	defer cmd.Process.Kill()

	s := Add(3, 4)
	if s != 7 {
		t.Errorf("exptect 7, got %v\n", s)
	}
}

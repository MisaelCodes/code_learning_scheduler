package logic

import (
	"fmt"
	"os/exec"
)

type Action interface {
	Activate() error
	Cancel() error
	IsOn() bool
	Run() error
}

type WebBrowsing struct {
	launcher string
	url      string
	on       bool
}

func (wb *WebBrowsing) Run() error {
	if !wb.on {
		return fmt.Errorf("can't launch %s because is the action is deactivated", wb.launcher)
	}
	cmd := exec.Command(wb.launcher, wb.url)
	return cmd.Run()
}

func (wb *WebBrowsing) IsOn() bool {
	return wb.on
}

func (wb *WebBrowsing) Cancel() error {
	if wb.on {
		wb.on = false
		return nil
	}
	return fmt.Errorf("the %s laucher is already deactivated", wb.launcher)
}

func (wb *WebBrowsing) Activate() error {
	if !wb.on {
		wb.on = true
		return nil
	}
	return fmt.Errorf("the %s laucher is already activated", wb.launcher)
}

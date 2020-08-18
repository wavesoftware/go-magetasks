package tasks

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/hashicorp/go-multierror"
	"github.com/wavesoftware/go-ensure"
	"github.com/wavesoftware/go-magetasks/internal"
)

type Task struct {
	icon      string
	action    string
	multiline bool
}

func Start(icon, action string) *Task {
	t := &Task{
		icon:      icon,
		action:    action,
		multiline: false,
	}
	t.start()
	return t
}

func StartMultiline(icon, action string) *Task {
	t := &Task{
		icon:      icon,
		action:    action,
		multiline: true,
	}
	t.start()
	return t
}

func (t *Task) start() {
	if t.multiline {
		fmt.Printf("%s %s %s\n", internal.MageTag, t.icon, t.action)
	} else {
		fmt.Printf("%s %s %s... ", internal.MageTag, t.icon, t.action)
	}
}

func (t *Task) End(errs ...error) {
	var msg string
	merr := multierror.Append(nil, errs...)
	err := merr.ErrorOrNil()
	green := color.New(color.FgHiGreen).Add(color.Bold).SprintFunc()
	red := color.New(color.FgHiRed).Add(color.Bold).SprintFunc()
	if err != nil {
		if t.multiline {
			msg = internal.MageTag + red(fmt.Sprintf(" %s have failed!\n", t.action))
		} else {
			msg = red(fmt.Sprintln("failed!"))
		}
	} else {
		if t.multiline {
			msg = internal.MageTag + green(fmt.Sprintf(" %s was successful.\n", t.action))
		} else {
			msg = green(fmt.Sprintln("done."))
		}
	}

	fmt.Print(msg)
	ensure.NoError(err)
}

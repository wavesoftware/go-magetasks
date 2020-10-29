package tasks

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/hashicorp/go-multierror"
	"github.com/wavesoftware/go-ensure"
)

// Task represents a mage task enriched with icon and description that might
// be multiline.
type Task struct {
	icon      string
	action    string
	multiline bool
}

// Start will start a single line task.
func Start(icon, action string) *Task {
	t := &Task{
		icon:      icon,
		action:    action,
		multiline: false,
	}
	t.start()
	return t
}

// StartMultiline will start a multi line task.
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
		fmt.Printf("%s %s %s\n", mageTag(), t.icon, t.action)
	} else {
		fmt.Printf("%s %s %s... ", mageTag(), t.icon, t.action)
	}
}

// End will report task completion, either successful or failures.
func (t *Task) End(errs ...error) {
	var msg string
	merr := multierror.Append(nil, errs...)
	err := merr.ErrorOrNil()
	if err != nil {
		msg = erroneousMsg(t)
	} else {
		msg = successfulMsg(t)
	}

	fmt.Print(msg)
	ensure.NoError(err)
}

func erroneousMsg(t *Task) string {
	red := color.New(color.FgHiRed).Add(color.Bold).SprintFunc()
	if t.multiline {
		return mageTag() + red(fmt.Sprintf(" %s have failed!\n", t.action))
	}
	return red(fmt.Sprintln("failed!"))
}

func successfulMsg(t *Task) string {
	green := color.New(color.FgHiGreen).Add(color.Bold).SprintFunc()
	if t.multiline {
		return mageTag() + green(fmt.Sprintf(" %s was successful.\n", t.action))
	}
	return green(fmt.Sprintln("done."))
}

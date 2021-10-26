package git

import (
	"regexp"
	"strings"

	"github.com/magefile/mage/sh"
)

// ref.: https://regex101.com/r/ppnq02/1
var remoteTagPattern = regexp.MustCompile(`^[0-9a-f]{7,}\s+refs/tags/([^^]+)(:?\^{})$`)

type shellInfo struct {
	Remote
}

func (s shellInfo) Description() (string, error) {
	return sh.Output("git", "describe", "--always", "--tags", "--dirty")
}

func (s shellInfo) Tags() ([]string, error) {
	output, err := sh.Output("git", "ls-remote", "--tags", s.remote())
	if err != nil {
		return nil, err
	}
	tags := parseLsRemoteTagsOutput(output)
	return tags, nil
}

func parseLsRemoteTagsOutput(output string) []string {
	lines := strings.Split(output, "\n")
	tagsMap := make(map[string]bool)
	for _, line := range lines {
		match := remoteTagPattern.FindSubmatch([]byte(line))
		if match == nil {
			continue
		}
		tag := string(match[1])
		tagsMap[tag] = true
	}
	tags := make([]string, 0, len(tagsMap))
	for tag := range tagsMap {
		tags = append(tags, tag)
	}
	return tags
}

func (s shellInfo) remote() string {
	if s.Remote.URL != "" {
		return s.Remote.URL
	}
	return s.Remote.Name
}

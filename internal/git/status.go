package git

import (
	"os/exec"
	"strings"
)

type FileStatus struct {
	Path   string
	Staged bool
	State  string
}

func GetStatus() ([]FileStatus, error) {
	out, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimRight(string(out), "\n"), "\n")
	var statuses []FileStatus
	for _, line := range lines {
		if len(line) < 3 {
			continue
		}
		indexStatus := line[0]
		workStatus := line[1]
		path := strings.TrimSpace(line[3:])

		fs := FileStatus{Path: path}

		switch {
		case indexStatus == '?' && workStatus == '?':
			fs.State = "untracked"

		case indexStatus == 'A':
			fs.State = "added"
			fs.Staged = true

		case indexStatus == 'M':
			fs.State = "modified"
			fs.Staged = true

		case indexStatus == 'D':
			fs.State = "deleted"
			fs.Staged = true

		case workStatus == 'M':
			fs.State = "modified"

		case workStatus == 'D':
			fs.State = "deleted"

		default:
			fs.State = "changed"
		}
		statuses = append(statuses, fs)
	}
	return statuses, nil
}

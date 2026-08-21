package git

import (
	"os/exec"
	"strconv"
	"strings"
)

type Commit struct {
	Hash    string
	Author  string
	Date    string
	Message string
}

func GetLog(lglimit int) ([]Commit, error) {
	// I asked for this format from claude
	format := "--pretty=format:%h\x1f%an\x1f%ar\x1f%s"
	out, err := exec.Command("git", "log", format, "-n", strconv.Itoa(lglimit)).Output()
	if err != nil {
		return nil, err
	}

	// seperate them
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	commits := make([]Commit, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "\x1f")
		if len(parts) != 4 {
			continue
		}
		commits = append(commits, Commit{
			Hash:    parts[0],
			Author:  parts[1],
			Date:    parts[2],
			Message: parts[3],
		})

	}
	return commits, nil
}

package git

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetGraph(limit int) ([]string, error) {
	format := "\x02%h\x1f%an\x1f%d\x1f%ar"
	out, err := exec.Command("git", "log", "--all", "--graph", "--pretty=format:"+format, "-n", strconv.Itoa(limit)).Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(string(out), "\n"), nil
}

func SpreadGraph(lines []string) []string {
	var out []string
	for _, line := range lines {
		out = append(out, line)

		// A spaccer typa thing
		graphPart := line
		if idx := strings.Index(line, "\x02"); idx != -1 {
			graphPart = line[:idx] // strip it
		}
		spacer := strings.ReplaceAll(graphPart, "*", "|")
		out = append(out, spacer)
	}
	return out
}

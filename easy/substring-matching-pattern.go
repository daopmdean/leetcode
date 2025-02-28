package easy

import "strings"

func hasMatch(s string, p string) bool {
	if len(p) == 1 {
		return true
	}

	asteriskIdx := strings.Index(p, "*")
	if asteriskIdx == -1 {
		return false
	}

	if asteriskIdx == 0 {
		return strings.Contains(s, p[1:])
	}

	if asteriskIdx == len(p)-1 {
		return strings.Contains(s, p[:len(p)-1])
	}

	firstPart := p[:asteriskIdx]
	firstPartIdx := strings.Index(s, firstPart)
	if firstPartIdx == -1 {
		return false
	}

	secondPart := p[asteriskIdx+1:]
	secondPartIdx := strings.Index(s[firstPartIdx+len(firstPart):], secondPart)
	if secondPartIdx == -1 {
		return false
	}

	return true
}

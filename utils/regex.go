package utils

import "regexp"

func MatchRouteWithRegex(route string, pattern string) (bool, []string) {
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(route)

	if len(matches) > 0 {
		return true, matches
	}
	return false, nil
}

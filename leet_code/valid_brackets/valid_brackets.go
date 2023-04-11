package valid_brackets

import (
	"fmt"
	"strings"
)

type TrackString struct {
	Track string
}

func (p *TrackString) TrackBracket(bracket string) bool {
	switch bracket {
	case "}":
		if strings.HasSuffix(p.Track, "{") {
			p.Track = p.Track[:len(p.Track)-1]
			return true
		} else {
			return false
		}
	case ")":
		if strings.HasSuffix(p.Track, "(") {
			p.Track = p.Track[:len(p.Track)-1]
			return true
		} else {
			return false
		}
	case "]":
		if strings.HasSuffix(p.Track, "[") {
			p.Track = p.Track[:len(p.Track)-1]
			return true
		} else {
			return false
		}
	default:
		p.Track = p.Track + bracket
		return true
	}
}

func isValid(s string) bool {
	tracker := &TrackString{}
	for _, bracket := range s {
		if !tracker.TrackBracket(fmt.Sprintf("%c", bracket)) {
			return false
		}
	}
	return tracker.Track == ""
}

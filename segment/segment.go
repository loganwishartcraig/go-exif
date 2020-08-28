package segment

import (
	"fmt"
)

type Segment struct {
	Body []byte
}

func (s *Segment) String() string {

	end := 50

	if len(s.Body) < 50 {
		end = len(s.Body)
	}

	return fmt.Sprintf("Segment Preview: [ % x ]", s.Body[:end])

}

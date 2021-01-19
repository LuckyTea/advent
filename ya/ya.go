package ya

import (
	"fmt"
	"sort"
)

func transform(in []int) string {
	if len(in) == 0 {
		return ""
	}

	sort.Ints(in)

	var (
		first = in[0]
		last  int
		comma bool
		x     string
		out   string
	)

	for i := range in {
		if i == 0 {
			last = in[i]
			continue
		}

		if last != in[i]-1 {
			if first != last {
				out = fmt.Sprintf("%s%s%d-%d", out, x, first, last)
			} else {
				out = fmt.Sprintf("%s%s%d", out, x, first)
			}

			if !comma {
				x = ","
				comma = true
			}

			if i == len(in) {
				out = fmt.Sprintf("%s,%d", out, in[i])

				break
			}

			first = in[i]
		}

		last = in[i]
	}

	if first != in[0] {
		out = fmt.Sprintf("%s,%d", out, in[len(in)-1])
	} else {
		out = fmt.Sprintf("%d-%d", first, in[len(in)-1])
	}

	return out
}

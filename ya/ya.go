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
		out   string
	)

	for i := range in {
		if i == 0 {
			last = in[i]
			continue
		}

		if last != in[i]-1 {
			if len(out) != 0 {
				out += ","
			}

			if first != last {
				out = fmt.Sprintf("%s%d-%d", out, first, last)
			} else {
				out = fmt.Sprintf("%s%d", out, first)
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

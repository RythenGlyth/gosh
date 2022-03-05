package util

import (
	"fmt"
	"strings"
)

type Args struct {
	Prog string

	// UParams are the unnamed parameters.
	UParams []string

	// NParams are the named parameters ( = parameters with value),
	// short and long options merged.
	NParams map[string][]string

	// SOps are the short options.
	SOps []rune

	// LOps are the long options, specified with two dashes.
	LOps []string
}

// ParseLine parses a command line with the following assumptions:
// The first word is the program name.
// Short options are one-character each, prefixed by a single dash.
// Short options can be grouped, e.g. -aS is equal to -a -S.
// Long options are entire strings and prefixed by two dashes.
//
// Both short and long options can have a value: a number of words
// that will be read if the option is specified in sparams or
// lparams (e.g. -s signal).
// If no value is found, it is added to the SOps instead.
//
// Everything else is assumed to be an unnamed parameter.
// Everything after two dashes -- without an option is also
// assumed to be an unnamed parameter, regardless of prefixes.
//
// It is possible to specify multiple named short options, e.g.
// -hsq 5 3 is equal to --help --signal 5 --queue 3
// assuming that s(ignal) and q(ueue) are specified as named
// short options.
func ParseLine(line string, sparams map[rune]int, lparams map[string]int) (*Args, error) {
	fields, err := Sep(line)
	if err != nil {
		return nil, err
	}

	args := Args{}
	args.Prog = fields[0]
	copy(fields, fields[1:])
	fields = fields[:len(fields)-1]
	args.NParams = make(map[string][]string)

	for i := 0; i < len(fields); i++ {
		field := fields[i]
		if field == "--" {
			// read rest as parameters
			i++
			for i < len(fields) {
				args.UParams = append(args.UParams, fields[i])
				i++
			}
		} else if strings.HasPrefix(field, "--") {
			op := field[2:]
			n := containsS(op, lparams)
			if n > 0 && i+n < len(fields) {
				// parse values after op
				args.NParams[op] = fields[i+1 : i+1+n]
				i += n
			} else {
				args.LOps = append(args.LOps, op)
			}
		} else if strings.HasPrefix(field, "-") {
			op := field[1:]
			for _, r := range op {
				n := containsR(r, sparams)
				if n > 0 && i+n < len(fields) {
					// parse values after op
					args.NParams[string(r)] = fields[i+1 : i+1+n]
					i += n
				} else {
					args.SOps = append(args.SOps, r)
				}
			}
		} else {
			args.UParams = append(args.UParams, field)
		}
	}

	return &args, nil
}

func (a *Args) String() string {
	var out strings.Builder
	fmt.Fprintf(&out, "Prog:  %s\n", a.Prog)

	fmt.Fprint(&out, "NPs:\n")
	for opt, vals := range a.NParams {
		fmt.Fprintf(&out, " - %s: %v\n", opt, vals)
	}

	fmt.Fprintf(&out, "SOps: %v\n", a.SOps)
	fmt.Fprintf(&out, "LOps: %v\n", a.LOps)
	fmt.Fprintf(&out, "UParams: %v\n", a.UParams)

	return out.String()
}

func containsR(x rune, xs map[rune]int) int {
	for i, n := range xs {
		if i == x {
			return n
		}
	}

	return 0
}

func containsS(x string, xs map[string]int) int {
	for i, n := range xs {
		if i == x {
			return n
		}
	}

	return 0
}

// Sep separates the line into arguments by splitting at spaces.
// Spaces inside quotes ("" and '') are kept,
// as well as those escaped by a backslash \.
// The quotes and the backslash are consumed.
func Sep(line string) ([]string, error) {
	r := []rune(line)
	args := []string{}
	j := 0
	args = append(args, "")

	for i := 0; i < len(r); i++ {
		if r[i] == '"' || r[i] == '\'' { // parse quoted string
			rq := r[i] // right quote = left quote for now
			lp := i    // left pos
			i++

			for r[i] != rq { // search for right quote
				i++
				if i >= len(r) {
					return nil, &ErrMissingClosingQuotes{rq, i, line}
				}
			}

			args[j] = string(r[lp+1 : i]) // append entire qoted string consuming quotes
			j++
			args = append(args, "")
		} else if r[i] == '\\' {
			i++             // consume backslash
			if i < len(r) { // ignore trailing backslash
				args[j] += string(r[i]) // append next character as-is
			}
		} else if r[i] == '\r' {
			// ignore
		} else if r[i] == ' ' || r[i] == '\t' || r[i] == '\n' {
			if args[j] != "" { // don't create empty arguments
				j++
				args = append(args, "")
			}
		} else {
			args[j] += string(r[i]) // append other characters as-is
		}
	}

	if args[j] == "" {
		args = args[:j]
	}

	return args, nil
}

type ErrMissingClosingQuotes struct {
	LeftQuote rune
	LeftPos   int
	Line      string
}

func (e *ErrMissingClosingQuotes) Error() string {
	return fmt.Sprintf("missing closing quote for opening %q at %d: %s",
		e.LeftQuote, e.LeftPos, e.Line)
}

package bkill

import (
	"fmt"
	"gosh/src/shared"
	"gosh/src/util"
	"strconv"
	"strings"
)

// Kill implements a built-in command to kill other processes.
// This can be useful if the system has run out of PIDs.
type Kill struct{}

var (
	kill_sparams = map[rune]int{'s': 1, 'q': 1, 'l': 1}
	kill_lparams = map[string]int{"signal": 1, "timeout": 2, "list": 1}
)

func (b *Kill) Match(line string) bool {
	return strings.HasPrefix(line, "kill")
}

func (b *Kill) Eval(g shared.IGosh, line string) error {
	args, err := util.ParseLine(line, kill_sparams, kill_lparams)
	if err != nil {
		return err
	}

	what, ok := args.NParams["l"]
	if !ok {
		what, ok = args.NParams["list"]
	}
	if ok {
		return b.convertSignal(g, what[0])
	}

	if util.ContainsR('l', args.SOps) >= 0 || util.ContainsS("list", args.LOps) >= 0 {
		return b.listSignals(g)
	}

	if util.ContainsR('L', args.SOps) >= 0 || util.ContainsS("table", args.LOps) >= 0 {
		return b.tableSignals(g)
	}

	// TODO

	return nil
}

func (b *Kill) convertSignal(g shared.IGosh, sig string) error {
	n, err := strconv.Atoi(sig)
	if err == nil {
		// convert number -> string
		mysig := bynumber(n)
		if mysig == "" {
			return &ErrUnknownSignal{sig}
		}

		g.WriteString(mysig)
		g.WriteString("\n")
	} else {
		// convert string -> string
		mysig := byname("SIG" + sig)
		if int(mysig) == 0 {
			return &ErrUnknownSignal{sig}
		}

		//g.WriteString(mysig.String())
		g.WriteString(sig)
		g.WriteString("\n")
	}
	return nil
}

func (b *Kill) listSignals(g shared.IGosh) error {
	var out strings.Builder
	llen := 0

	for _, t := range allsigs {
		out.WriteString(t.name)

		llen += len(t.name) + 1
		if llen >= 68 {
			out.WriteString("\n")
			llen = 0
		} else {
			out.WriteString(" ")
		}
	}

	if llen != 0 {
		out.WriteString("\n")
	}

	g.WriteString(out.String())

	return nil
}

func (b *Kill) tableSignals(g shared.IGosh) error {
	var cells []string

	for _, t := range allsigs {
		cells = append(cells, fmt.Sprintf("%2d %s", t.sig, t.name))
	}

	w, _ := g.Size()
	util.PrintTable(g, cells, w)

	return nil
}

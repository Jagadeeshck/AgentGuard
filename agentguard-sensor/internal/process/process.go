package process

import (
	"bufio"
	"os/exec"
	"strconv"
	"strings"
)

type Proc struct {
	Name                    string
	PID, PPID               int
	Executable, CommandLine string
}

func Scan() []Proc {
	cmd := exec.Command("ps", "-eo", "pid,ppid,comm,args")
	out, err := cmd.Output()
	if err != nil {
		return nil
	}
	return ParsePSOutput(string(out))
}

func ParsePSOutput(out string) []Proc {
	s := bufio.NewScanner(strings.NewReader(out))
	res := []Proc{}
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if strings.HasPrefix(l, "PID") || l == "" {
			continue
		}
		f := strings.Fields(l)
		if len(f) < 4 {
			continue
		}
		pid, _ := strconv.Atoi(f[0])
		ppid, _ := strconv.Atoi(f[1])
		res = append(res, Proc{Name: f[2], PID: pid, PPID: ppid, Executable: f[2], CommandLine: strings.Join(f[3:], " ")})
	}
	return res
}

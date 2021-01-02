package agentUtils

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetExeDir() string { // where exe file in
	var dir string
	arg1 := strings.ToLower(os.Args[0])

	name := filepath.Base(arg1)
	if strings.Index(name, "tester") == 0 && strings.Index(arg1, "go-build") < 0 {
		p, _ := exec.LookPath(os.Args[0])
		if strings.Index(p, string(os.PathSeparator)) > -1 {
			dir = p[:strings.LastIndex(p, string(os.PathSeparator))]
		}
	} else { // debug
		dir, _ = os.Getwd()
	}

	dir, _ = filepath.Abs(dir)
	dir = AddSepIfNeeded(dir)

	//fmt.Printf("Debug: Launch %s in %s \n", arg1, dir)
	return dir
}

func AddSepIfNeeded(pth string) string {
	sepa := string(os.PathSeparator)

	if strings.LastIndex(pth, sepa) < len(pth)-1 {
		pth += sepa
	}
	return pth
}

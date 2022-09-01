package updater

import (
	"bytes"
	"github.com/metafates/mangal/constant"
	"github.com/samber/lo"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type installationMethod int

const (
	unknown installationMethod = iota
	homebrew
	scoop
	termux
	script
)

// detectInstallationMethod detects the installation method.
func detectInstallationMethod() installationMethod {
	for _, t := range []lo.Tuple2[installationMethod, func() bool]{
		{scoop, isUnderScoop},
		{homebrew, isUnderHomebrew},
		{termux, isUnderTermux},
	} {
		if t.B() {
			return t.A
		}
	}

	if lo.Contains([]string{"darwin", "linux"}, runtime.GOOS) {
		path, err := os.Executable()
		if err != nil {
			return unknown
		}

		if path == "/usr/local/bin/"+constant.Mangal {
			return script
		}
	}

	return unknown
}

// isUnderTermux returns true if mangal is running under termux.
func isUnderTermux() (ok bool) {
	return has("termux-setup-storage")
}

// isUnderHomebrew returns true if mangal is running under homebrew.
func isUnderHomebrew() (ok bool) {
	if !has("brew") {
		return
	}

	out, err := execute("brew", "list", "--formula")
	if err != nil {
		return false
	}

	return strings.Contains(out, constant.Mangal)
}

// isUnderScoop returns true if mangal is running under scoop.
func isUnderScoop() (ok bool) {
	if !has("scoop") {
		return false
	}

	path, err := os.Executable()
	if err != nil {
		return false
	}

	return strings.Contains(path, filepath.Join("scoop", "shims"))
}

// has returns true if the command exists.
func has(command string) bool {
	ok, err := exec.LookPath(command)
	return err == nil && ok != ""
}

// execute the command and returns the output as string.
func execute(command string, arguments ...string) (output string, err error) {
	stdout := bytes.NewBufferString("")

	cmd := exec.Command(command, arguments...)
	cmd.Stdout = stdout
	err = cmd.Run()

	return stdout.String(), err
}

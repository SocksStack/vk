package generate

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/SocksStack/vk/utils"
	"log"
	"os/exec"
	"path"
	"runtime"
)

func Create(moduleName, curPath string) error {
	err := utils.Mkdir(curPath)

	mod := path.Join(curPath, "go.mod")
	// 如果存在则返回
	if utils.IsExist(mod) {
		return nil
	}

	if err != nil {
		panic(err)
	}
	var stdout, stderr bytes.Buffer
	var cmd *exec.Cmd
	goos := runtime.GOOS
	switch goos {
	case "windows":
		cmd = exec.Command("cmd", "/c", fmt.Sprintf("cd %s && go mod init %s && cd ..", moduleName, moduleName))
		break
	case "linux":
		cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && go mod init %s && cd ..", moduleName, moduleName))
		break
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return errors.New(stderr.String())
	}
	out := stdout.String()
	if out != "" {
		log.Printf(out)
	}
	return nil
}
package log

import (
	"github.com/kpango/glg"
	"os"
	"path/filepath"
)

func InitLogger(logsPath string, lmode glg.MODE, filename string) *os.File {
	mode := int(0777)
	log := glg.FileWriter(filepath.Join(logsPath, filename+".log"), os.FileMode(mode))
	//rotate := NewRotateWriter(infolog, time.Second*10, bytes.NewBuffer(make([]byte, 0, 4096)))
	glg.Get().SetMode(lmode).AddLevelWriter(glg.INFO, log).AddLevelWriter(glg.ERR, log).AddLevelWriter(glg.LOG, log).AddLevelWriter(glg.FATAL, log).AddLevelWriter(glg.WARN, log)
	return log
}

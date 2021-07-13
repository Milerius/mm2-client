package mm2_tools_server

import (
	"github.com/kpango/glg"
	"os"
	"path/filepath"
)

func InitLogger(logsPath string) (*os.File, *os.File) {
	mode := int(0777)
	infolog := glg.FileWriter(filepath.Join(logsPath, "mm2.tools.server.info.log"), os.FileMode(mode))
	errlog := glg.FileWriter(filepath.Join(logsPath, "/mm2.tools.server.error.log"), os.FileMode(mode))
	glg.Get().SetMode(glg.BOTH).AddLevelWriter(glg.INFO, infolog).AddLevelWriter(glg.ERR, errlog).AddLevelWriter(glg.LOG, infolog).AddLevelWriter(glg.FATAL, errlog)
	return infolog, errlog
}

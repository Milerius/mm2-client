package constants

import "mm2_client/helpers"

var (
	GMM2Dir       = helpers.GetWorkingDir() + "/mm2"
	GMM2BinPath   = GMM2Dir + "/mm2"
	GMM2ConfPath  = GMM2Dir + "/MM2.json"
	GMM2CoinsPath = GMM2Dir + "/coins.json"
)

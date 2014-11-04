package goenv

func (goenv *Goenv) GetSyncDbFilename() (dbname string) {
	dbname = goenv.RequireFailToEnv("EBOH_DB_NAME")
	return
}

func (goenv *Goenv) GetListBatchSize() (listBatchSize int) {
	listBatchSize = goenv.GetIntFailToEnv("EBOH_LIST_BATCH_SIZE", 20)
	return
}

func (goenv *Goenv) GetUploadWorkers() (listBatchSize int) {
	listBatchSize = goenv.GetIntFailToEnv("EBOH_UPLOAD_WORKERS", 5)
	return
}

func (goenv *Goenv) GetDownloadWorkers() (listBatchSize int) {
	listBatchSize = goenv.GetIntFailToEnv("EBOH_DOWNLOAD_WORKERS", 5)
	return
}

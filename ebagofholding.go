package goenv

func (goenv *Goenv) GetSyncDbFilename() (dbname string) {
	dbname = goenv.RequireFailToEnv("EBOH_DB_NAME")
	return
}

func (goenv *Goenv) GetListBatchSize() (listBatchSize int) {
	listBatchSize = goenv.GetIntFailToEnv("EBOH_LIST_BATCH_SIZE", 20)
	return
}

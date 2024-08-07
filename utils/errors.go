package utils

func HandleError(err error, msg string) error {
	if err != nil {
		LOG.Error(err)
		LOG.Error(msg)
		return err
	}

	return nil
}

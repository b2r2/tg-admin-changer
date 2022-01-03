package repositories

func (r *repo) Close() error {
	if err := r.cache.Close(); err != nil {
		return err
	}

	return r.db.Close()
}

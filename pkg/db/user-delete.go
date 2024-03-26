package db

func (p *postgres) DeleteUserById(idUser int) (err error) {
	_, err = p.db.Exec("DELETE FROM users WHERE id = $1", idUser)
	if err != nil {
		return
	}

	return
}

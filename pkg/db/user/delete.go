package user

func (p *postgres) Delete(idUser int) (err error) {
	_, err = p.db.Exec("DELETE FROM users WHERE id = $1", idUser)
	if err != nil {
		return
	}

	return
}

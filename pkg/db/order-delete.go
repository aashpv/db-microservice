package db

func (p *postgres) DeleteOrderById(idOrder int) (err error) {
	_, err = p.db.Exec("DELETE FROM orders WHERE id = $1", idOrder)
	if err != nil {
		return
	}

	return
}

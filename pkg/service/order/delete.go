package order

func (s *service) DeleteOrderById(id int) (err error) {
	err = s.pgs.DeleteOrderById(id)
	if err != nil { // if err use log for writing in file
		return
	}
	// если ошибок нет, то тоже надо логировать

	return
}

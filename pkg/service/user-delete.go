package service

func (s *service) DeleteUserById(id int) (err error) {
	err = s.pgs.DeleteUserById(id)
	if err != nil { // if err use log for writing in file
		return
	}
	// если ошибок нет, то тоже надо логировать

	return
}

package service

func (s *Service) Delete(id int64) error {
	return s.repo.Delete(id)
}

package mocks

type ScreenPresenter struct {
	OnPresent                    func() error
	OnClear                      func() error
	CallCntPresent, CallCntClear int
}

func (s *ScreenPresenter) Present() error {
	s.CallCntPresent++
	if s.OnPresent != nil {
		return s.OnPresent()
	}
	return nil
}

func (s *ScreenPresenter) Clear() error {
	s.CallCntClear++
	if s.OnClear != nil {
		return s.OnClear()
	}
	return nil
}

package service

import "tsc-p9-part3/mocks"

type Service struct {
	sender mocks.EmailSender
}

func NewService(sender mocks.EmailSender) *Service {
	return &Service{sender: sender}
}

func (s *Service) Notify(to, subject, body string) error {
	return s.sender.Send(to, subject, body)
}

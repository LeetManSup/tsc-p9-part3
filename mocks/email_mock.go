package mocks

import "fmt"

type MockEmailSender struct {
	Calls []string
}

func (m *MockEmailSender) Send(to, subject, body string) error {
	// Логируем вызов
	m.Calls = append(m.Calls, fmt.Sprintf("%s|%s|%s", to, subject, body))
	return nil
}

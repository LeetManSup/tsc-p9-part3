package mocks

// EmailSender описывает поведение отправщика писем
type EmailSender interface {
	Send(to, subject, body string) error
}

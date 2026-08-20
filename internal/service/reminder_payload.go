package service

import "todolist/internal/payload"

type PayloadService struct {
	producer *payload.Producer
	consumer *payload.Consumer
}

func NewPayloadService(producer *payload.Producer, consumer *payload.Consumer) *PayloadService {
	return &PayloadService{producer: producer, consumer: consumer}
}

func (s *PayloadService) Queue(id, message string) []byte {
	queued := s.producer.Submit(id, message)
	return append([]byte(nil), queued...)
}

func (s *PayloadService) Read(id string) string {
	return string(append([]byte(nil), s.consumer.Read(id)...))
}

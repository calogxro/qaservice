package service

import (
	"github.com/calogxro/qaservice/db"
	"github.com/calogxro/qaservice/domain"
)

type QAService struct {
	eventStore db.IEventStore
}

func NewQAService(es db.IEventStore) *QAService {
	return &QAService{
		eventStore: es,
	}
}

func (s *QAService) answerExists(key string) bool {
	answer, _ := RecreateAnswerState(s.eventStore, key)
	return answer != nil
}

func (s *QAService) CreateAnswer(answer domain.Answer) (*domain.Event, error) {
	if s.answerExists(answer.Key) {
		return nil, &domain.KeyExists{}
	}
	event, _ := domain.NewAnswerCreatedEvent(answer)
	s.eventStore.AddEvent(event)
	return event, nil
}

func (s *QAService) UpdateAnswer(answer domain.Answer) (*domain.Event, error) {
	if !s.answerExists(answer.Key) {
		return nil, &domain.KeyNotFound{}
	}
	event, _ := domain.NewAnswerUpdatedEvent(answer)
	s.eventStore.AddEvent(event)
	return event, nil
}

func (s *QAService) DeleteAnswer(key string) (*domain.Event, error) {
	if !s.answerExists(key) {
		return nil, &domain.KeyNotFound{}
	}
	event, _ := domain.NewAnswerDeletedEvent(domain.Answer{Key: key})
	s.eventStore.AddEvent(event)
	return event, nil
}

func (s *QAService) GetHistory(key string) ([]*domain.Event, error) {
	events, _ := s.eventStore.GetHistory(key)
	return events, nil
}

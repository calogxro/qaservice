package eventstore

import (
	"encoding/json"

	"github.com/calogxro/qaservice/domain"
)

func RecreateAnswerState(store EventStore, key string) (*domain.Answer, error) {
	var answer *domain.Answer
	events, _ := store.GetHistory(key)
	if len(events) > 0 {
		lastEvent := events[len(events)-1]
		if lastEvent != nil && lastEvent.Type != domain.ANSWER_DELETED_EVENT {
			err := json.Unmarshal([]byte(lastEvent.Data), &answer)
			if err != nil {
				return nil, err
			}
		}
	}
	return answer, nil
}

// AnswerExists returns false if the key was deleted.
func AnswerExists(eventStore EventStore, key string) bool {
	answer, _ := RecreateAnswerState(eventStore, key)
	return answer != nil
}

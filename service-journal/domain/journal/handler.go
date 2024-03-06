package journal

import (
	"service-journal/domain/journal/feature"
)

type JournalHandler interface {
}

type journalHandler struct {
	journalFeature feature.JournalFeature
}

func NewJournalHandler(journalFeature feature.JournalFeature) JournalHandler {
	return &journalHandler{
		journalFeature: journalFeature,
	}
}

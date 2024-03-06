package server

import (
	"service-journal/delivery/container"
	"service-journal/domain/journal"
)

type handler struct {
	journalHandler journal.JournalHandler
}

func SetupHandler(container container.Container) handler {
	return *&handler{
		journalHandler: journal.NewJournalHandler(container.JournalFeature),
	}
}

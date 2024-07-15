package docupd

import (
	"sync"
)

type Processor interface {
	Process(d *Document) (*Document, error)
}

type DocumentProcessor struct {
	mu       sync.Mutex
	stateMap map[string]*DocumentState // мапа, в которой хранятся состояния документов c ключем в виде Url (по-хорошему надо сделать через бд)
}

// структура, которая будет хранить самое ранее и самое позднее состояния документа
type DocumentState struct {
	EarliestFetchDoc *Document
	LatestFetchDoc   *Document
}

func NewDocumentProcessor() *DocumentProcessor {
	return &DocumentProcessor{
		stateMap: make(map[string]*DocumentState),
	}
}

func (dp *DocumentProcessor) Process(d *Document) (*Document, error) {
	dp.mu.Lock()
	defer dp.mu.Unlock()

	state, exists := dp.stateMap[d.Url]
	switch {
	case !exists:
		d.FirstFetchTime = d.FetchTime
		state = &DocumentState{
			EarliestFetchDoc: d,
			LatestFetchDoc:   d,
		}
		dp.stateMap[d.Url] = state
		return d, nil

	case d.FetchTime > state.LatestFetchDoc.FetchTime:
		state.LatestFetchDoc = d

	case d.FetchTime < state.EarliestFetchDoc.FetchTime:
		d.FirstFetchTime = d.FetchTime
		state.EarliestFetchDoc = d
	}

	result := &Document{
		Url:            d.Url,
		PubDate:        state.EarliestFetchDoc.PubDate,
		FetchTime:      state.LatestFetchDoc.FetchTime,
		Text:           state.LatestFetchDoc.Text,
		FirstFetchTime: state.EarliestFetchDoc.FirstFetchTime,
	}

	return result, nil
}

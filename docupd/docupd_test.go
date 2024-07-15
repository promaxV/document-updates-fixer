package docupd

import (
	"testing"
)

// Тест на соответствие базовым правилам
func TestBasicRules(t *testing.T) {
	processor := NewDocumentProcessor()

	inq := []*Document{
		{Url: "http://tests.docs.com", PubDate: 11, FetchTime: 2, Text: "Second version text"},
		{Url: "http://tests.docs.com", PubDate: 12, FetchTime: 3, Text: "Third version text"},
		{Url: "http://tests.docs.com", PubDate: 13, FetchTime: 0, Text: "Initial version text"},
		{Url: "http://tests.docs.com", PubDate: 14, FetchTime: 1, Text: "First version text"},
		{Url: "http://tests.docs.com", PubDate: 15, FetchTime: 4, Text: "Fourth version text"},
	}

	expectedPubDates := []uint64{11, 11, 13, 13, 13}
	expectedFetchTimes := []uint64{2, 3, 3, 3, 4}
	expectedTexts := []string{"Second version text", "Third version text", "Third version text", "Third version text", "Fourth version text"}
	expectedFirstFetchTimes := []uint64{2, 2, 0, 0, 0}

	outq := []*Document{}

	for len(inq) > 0 {
		inDoc := inq[0]
		processedDoc, err := processor.Process(inDoc)
		if err == nil {
			outq = append(outq, processedDoc)
		}
		inq = inq[1:]
	}

	for i, upd := range outq {
		if upd.FetchTime != expectedFetchTimes[i] {
			t.Errorf("FetchTime error: expected %v got %v in doc %d", expectedFetchTimes[i], upd.FetchTime, i)
		}
		if upd.Text != expectedTexts[i] {
			t.Errorf("Text error: expected %v got %v in doc %d", expectedTexts[i], upd.Text, i)
		}
		if upd.PubDate != expectedPubDates[i] {
			t.Errorf("PubDate error: expected %v got %v in doc %d", expectedPubDates[i], upd.PubDate, i)
		}
		if upd.FirstFetchTime != expectedFirstFetchTimes[i] {
			t.Errorf("FirstFetchTime error: expected %v got %v in doc %d", expectedFirstFetchTimes[i], upd.FirstFetchTime, i)
		}
	}
}

// Тест с дублированными сообщениями
func TestDoubles(t *testing.T) {
	processor := NewDocumentProcessor()

	inq := []*Document{
		{Url: "http://tests.docs.com", PubDate: 11, FetchTime: 2, Text: "Second version text"},
		{Url: "http://tests.docs.com", PubDate: 12, FetchTime: 3, Text: "Third version text"},
		{Url: "http://tests.docs.com", PubDate: 13, FetchTime: 0, Text: "Initial version text"},
		{Url: "http://tests.docs.com", PubDate: 14, FetchTime: 1, Text: "First version text"},
		{Url: "http://tests.docs.com", PubDate: 15, FetchTime: 4, Text: "Fourth version text"},
		{Url: "http://tests.docs.com", PubDate: 13, FetchTime: 0, Text: "Initial version text"},
		{Url: "http://tests.docs.com", PubDate: 12, FetchTime: 3, Text: "Third version text"},
		{Url: "http://tests.docs.com", PubDate: 13, FetchTime: 0, Text: "Initial version text"},
		{Url: "http://tests.docs.com", PubDate: 14, FetchTime: 1, Text: "First version text"},
	}

	expectedPubDates := []uint64{11, 11, 13, 13, 13, 13, 13, 13, 13}
	expectedFetchTimes := []uint64{2, 3, 3, 3, 4, 4, 4, 4, 4}
	expectedTexts := []string{"Second version text", "Third version text", "Third version text", "Third version text", "Fourth version text", "Fourth version text", "Fourth version text", "Fourth version text", "Fourth version text"}
	expectedFirstFetchTimes := []uint64{2, 2, 0, 0, 0, 0, 0, 0, 0}

	outq := []*Document{}

	for len(inq) > 0 {
		inDoc := inq[0]
		processedDoc, err := processor.Process(inDoc)
		if err == nil {
			outq = append(outq, processedDoc)
		}
		inq = inq[1:]
	}

	for i, upd := range outq {
		if upd.FetchTime != expectedFetchTimes[i] {
			t.Errorf("FetchTime error: expected %v got %v in doc %d", expectedFetchTimes[i], upd.FetchTime, i)
		}
		if upd.Text != expectedTexts[i] {
			t.Errorf("Text error: expected %v got %v in doc %d", expectedTexts[i], upd.Text, i)
		}
		if upd.PubDate != expectedPubDates[i] {
			t.Errorf("PubDate error: expected %v got %v in doc %d", expectedPubDates[i], upd.PubDate, i)
		}
		if upd.FirstFetchTime != expectedFirstFetchTimes[i] {
			t.Errorf("FirstFetchTime error: expected %v got %v in doc %d", expectedFirstFetchTimes[i], upd.FirstFetchTime, i)
		}
	}
}

// Тест с разными документами
func TestTwoDocs(t *testing.T) {
	processor := NewDocumentProcessor()

	inq := []*Document{
		{Url: "http://1.tests.docs.com", PubDate: 11, FetchTime: 2, Text: "Second version text"},
		{Url: "http://1.tests.docs.com", PubDate: 12, FetchTime: 3, Text: "Third version text"},
		{Url: "http://2.tests.docs.com", PubDate: 13, FetchTime: 0, Text: "Initial version text 2"},
		{Url: "http://1.tests.docs.com", PubDate: 14, FetchTime: 1, Text: "First version text"},
		{Url: "http://2.tests.docs.com", PubDate: 15, FetchTime: 4, Text: "Fourth version text 2"},
		{Url: "http://1.tests.docs.com", PubDate: 16, FetchTime: 0, Text: "Initial version text"},
		{Url: "http://2.tests.docs.com", PubDate: 17, FetchTime: 3, Text: "Third version text 2"},
		{Url: "http://2.tests.docs.com", PubDate: 13, FetchTime: 0, Text: "Initial version text 2"},
		{Url: "http://2.tests.docs.com", PubDate: 19, FetchTime: 1, Text: "First version text 2"},
	}

	expectedPubDates := []uint64{11, 11, 13, 14, 13, 16, 13, 13, 13}
	expectedFetchTimes := []uint64{2, 3, 0, 3, 4, 3, 4, 4, 4}
	expectedTexts := []string{"Second version text", "Third version text", "Initial version text 2", "Third version text", "Fourth version text 2", "Third version text", "Fourth version text 2", "Fourth version text 2", "Fourth version text 2"}
	expectedFirstFetchTimes := []uint64{2, 2, 0, 1, 0, 0, 0, 0, 0}

	outq := []*Document{}

	for len(inq) > 0 {
		inDoc := inq[0]
		processedDoc, err := processor.Process(inDoc)
		if err == nil {
			outq = append(outq, processedDoc)
		}
		inq = inq[1:]
	}

	for i, upd := range outq {
		if upd.FetchTime != expectedFetchTimes[i] {
			t.Errorf("FetchTime error: expected %v got %v in doc %d", expectedFetchTimes[i], upd.FetchTime, i)
		}
		if upd.Text != expectedTexts[i] {
			t.Errorf("Text error: expected %v got %v in doc %d", expectedTexts[i], upd.Text, i)
		}
		if upd.PubDate != expectedPubDates[i] {
			t.Errorf("PubDate error: expected %v got %v in doc %d", expectedPubDates[i], upd.PubDate, i)
		}

		if upd.FirstFetchTime != expectedFirstFetchTimes[i] {
			t.Errorf("FirstFetchTime error: expected %v got %v in doc %d", expectedFirstFetchTimes[i], upd.FirstFetchTime, i)
		}
	}

}

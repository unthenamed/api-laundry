package model

import "time"

var (
	layoutDate   = "02-01-2006"
	defaultDate  = "2006-01-02"
	databaseDate = time.RFC3339
)

func (r *Response) FormatDate() {
	r.BillDate = format(parse(r.BillDate, databaseDate), layoutDate)
	r.EntryDate = format(parse(r.EntryDate, databaseDate), layoutDate)
	r.FinishDate = format(parse(r.FinishDate, databaseDate), layoutDate)
}

func parse(s string, l string) time.Time {
	r, _ := time.Parse(l, s)
	return r
}

func format(t time.Time, l string) string {
	return t.Format(l)

}

func (t *Transaction) QueryDate() {
	if t.Query.StartDate != "" {
		t.Query.StartDate = format(parse(t.Query.StartDate, layoutDate), defaultDate)
	}
	if t.Query.EndDate != "" {
		t.Query.EndDate = format(parse(t.Query.EndDate, layoutDate), defaultDate)
	}
}

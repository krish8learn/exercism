package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func commatize(i int, thou string) string {
	// convert eg 1000 to "1,000", where the ',' is from the thou arg
	if i < 1000 {
		return fmt.Sprintf("%d", i)
	}
	return fmt.Sprintf("%s%s%03d", commatize(i/1000, thou), thou, i%1000)
}

type cs struct {
	i int
	s string
	e error
}

func process_entry(co chan cs, i int, entry Entry, locale, currency string) {
	var res cs
	if len(entry.Date) != 10 || entry.Date[4] != '-' || entry.Date[7] != '-' {
		res = cs{e: errors.New("bad date format")}
	} else {
		date, desc, cents := entry.Date, entry.Description, entry.Change
		year, month, day, head, tail := date[0:4], date[5:7], date[8:10], "", " "
		var dt string // decimal point and thousands separator
		switch locale {
		case "nl-NL":
			date, dt = day+"-"+month+"-"+year, ",."
		case "en-US":
			date, dt = month+"/"+day+"/"+year, ".,"
		}
		if len(desc) > 25 {
			desc = desc[:22] + "..."
		}
		negative := cents < 0
		if negative {
			cents *= -1
		}
		switch locale {
		case "nl-NL":
			if negative {
				tail = "-"
			}
			currency += " "
		case "en-US":
			if negative {
				head, tail = "(", ")"
			}
		}
		pounds := commatize(cents/100, string(dt[1]))
		pence := fmt.Sprintf("%s%02d", string(dt[0]), cents%100)
		amount := head + currency + pounds + pence + tail
		line := fmt.Sprintf("%-10s | %-25s | %13s\n", date, desc, amount)
		res = cs{i: i, s: line}
	}
	co <- res
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	switch currency {
	case "EUR":
		currency = "€"
	case "USD":
		currency = "$"
	default:
		return "", errors.New("unknown currency")
	}
	var date, desc, change string
	switch locale {
	case "nl-NL":
		date, desc, change = "Datum", "Omschrijving", "Verandering"
	case "en-US":
		date, desc, change = "Date", "Description", "Change"
	default:
		return "", errors.New("unknown locale")
	}
	col_hdr := fmt.Sprintf("%-10s | %-25s | %s\n", date, desc, change)
	entries_copy := make([]Entry, len(entries))
	copy(entries_copy, entries)
	compare_func := func(i, j int) bool {
		ei, ej := entries_copy[i], entries_copy[j]
		ei_Desc, ej_Desc := ei.Description, ej.Description
		if ei.Date != ej.Date {
			return ei.Date < ej.Date
		}
		if ei_Desc != ej_Desc {
			return ei_Desc < ej_Desc
		}
		return ei.Change < ej.Change
	}
	sort.Slice(entries_copy, compare_func)
	// Parallelism, always a great idea
	co := make(chan cs)
	for i, entry := range entries_copy {
		go process_entry(co, i, entry, locale, currency)
	}
	ss := make([]string, len(entries_copy))
	for range entries_copy {
		res := <-co
		if res.e != nil {
			return "", res.e
		}
		ss[res.i] = res.s
	}
	res := col_hdr + strings.Join(ss, "")
	return res, nil
}

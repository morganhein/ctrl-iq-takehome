package control_iq_takehome

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Operation string

const (
	UPLOAD   Operation = "upload"
	DOWNLOAD Operation = "download"
)

type Entry struct {
	Date      time.Time
	Username  string
	Operation Operation
	Size      int //kB
}

// Report analyzes a file and spits out the results
func Report(fileName string) error {
	raw, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	entries, err := Parse(raw)
	if err != nil {
		return err
	}
	fmt.Printf("Number of users: %v\n", countUsers(entries))
	rand.Seed(time.Now().Unix())
	sizeReq := rand.Intn(85-1) + 1
	fmt.Printf("Number of uploads above %vKB: %v\n", sizeReq, countUploadsAboveSize(entries, sizeReq))
	user := "jeff22"
	day := time.Date(2020, 4, 15, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Number of uploads by %v on %v: %v\n", user, day.Format("2006-01-02"), countUserActionsOnDay(entries, user, UPLOAD, day))
	return nil
}

// Parse expects a csv input with the following fields:
// timestamp (date), username (string), operation (string/enum), size (int)
func Parse(data []byte) ([]Entry, error) {
	r := bytes.NewReader(data)
	csvr := csv.NewReader(r)
	records, err := csvr.ReadAll()
	if err != nil {
		return nil, err
	}
	//the first record is the header, so we have to skip it, and set indexes appropriately
	results := make([]Entry, len(records)-1)
	for i := 1; i < len(records); i++ {
		rec := records[i]
		//parse time
		t, err := time.Parse(time.UnixDate, rec[0])
		if err != nil {
			return nil, err
		}
		//parse size
		s, err := strconv.Atoi(rec[3])
		if err != nil {
			return nil, err
		}
		results[i-1] = Entry{
			Date:      t,
			Username:  rec[1],
			Operation: Operation(rec[2]),
			Size:      s,
		}
	}
	return results, nil
}

// countUsers returns the count of unique users in a list of entries
func countUsers(entries []Entry) int {
	uniqueUsers := make(map[string]interface{})
	for _, v := range entries {
		uniqueUsers[v.Username] = nil
	}
	return len(uniqueUsers)
}

// countUploadsAboveSize performs a > comparison, NOT >=
func countUploadsAboveSize(entries []Entry, size int) int {
	count := 0
	for _, v := range entries {
		if v.Size > size {
			count++
		}
	}
	return count
}

func countUserActionsOnDay(entries []Entry, user string, op Operation, t time.Time) int {
	count := 0
	for _, v := range entries {
		if strings.ToLower(v.Username) != strings.ToLower(user) {
			continue
		}
		if v.Operation != op {
			continue
		}
		if !isSameDay(t, v.Date) {
			continue
		}
		count++
	}
	return count
}

func isSameDay(t1 time.Time, t2 time.Time) bool {
	return t1.Truncate(24 * time.Hour).Equal(t2.Truncate(24 * time.Hour))
}

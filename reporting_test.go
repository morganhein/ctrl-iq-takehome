package control_iq_takehome

import (
	"os"
	"testing"
	"time"
)

//These tests should probably be done on a smaller subset of the csv, but for expediencies sake...

func TestParse(t *testing.T) {
	data, err := os.ReadFile("server_log.csv")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	entries, err := Parse(data)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	if len(entries) != 657 {
		t.Logf("Expected to to find 657 entries, but found %v instead", len(entries))
		t.FailNow()
	}
	//these tests could be data driven, or whatever. For now just a sanity check
	if entries[0].Username != "sarah94" {
		t.Logf("Expected to to find sarah94 as the first entry, but found %v instead", entries[0].Username)
		t.FailNow()
	}
	if entries[len(entries)-1].Username != "jeff22" {
		t.Logf("Expected to to find jeff22 as the first entry, but found %v instead", entries[0].Username)
		t.FailNow()
	}
}

func TestCountUsers(t *testing.T) {
	data, err := os.ReadFile("server_log.csv")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	entries, err := Parse(data)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	userCount := countUsers(entries)
	if userCount != 6 { //known magic number, which should instead be determined from some testdata in this function
		t.Logf("Expected to to find 6 users, but found %v instead", userCount)
		t.FailNow()
	}
}

func TestCountUsersAboveSize(t *testing.T) {
	data, err := os.ReadFile("server_log.csv")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	entries, err := Parse(data)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	uploadSizeCount := countUploadsAboveSize(entries, 50)
	if uploadSizeCount != 277 { //known magic number, which should instead be determined from some testdata in this function
		t.Logf("Expected to to find 277 uploads, but found %v instead", uploadSizeCount)
		t.FailNow()
	}
}

func TestCountUserActionsOnDay(t *testing.T) {
	data, err := os.ReadFile("server_log.csv")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	entries, err := Parse(data)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	interactionCount := countUserActionsOnDay(entries, "jeff22", UPLOAD, time.Date(2020, 4, 15, 0, 0, 0, 0, time.UTC))
	if interactionCount != 3 { //known magic number, which should instead be determined from some testdata in this function
		t.Logf("Expected to to find 3 interactions, but found %v instead", interactionCount)
		t.FailNow()
	}
}

func TestIsSameDay(t *testing.T) {
	t1 := time.Unix(123, 0)
	t2 := time.Unix(124, 2)
	equal := isSameDay(t1, t2)
	if !equal {
		t.FailNow()
	}

	t1 = time.Date(2022, 5, 4, 3, 2, 1, 1, time.UTC)
	t2 = time.Date(2022, 5, 4, 20, 0, 2, 4, time.UTC)

	equal = isSameDay(t1, t2)
	if !equal {
		t.FailNow()
	}
}

package process

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

type testJSON struct {
	Test JSONTime `json:"test"`
}

var (
	jsonDate = []byte("{\"test\":\"2018-08-18T10:31:17+02:00\"}")
	rawDate  = "2018-08-18T10:31:17+02:00"
)

func getTestTime(t *testing.T) JSONTime {
	date, err := time.Parse(time.RFC3339, rawDate)
	if err != nil {
		t.Fatal(err)
	}

	return JSONTime{Time: date}
}

func TestJSONTime_MarshalJSON(t *testing.T) {
	test := &testJSON{
		Test: JSONTime(getTestTime(t)),
	}

	b, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(b, jsonDate) {
		t.Fatalf("Expected to get %q, got %q", string(jsonDate), string(b))
	}
}

func TestJSONTime_UnmarshalJSON(t *testing.T) {
	test := &testJSON{}
	testTime := getTestTime(t)

	if err := json.Unmarshal(jsonDate, test); err != nil {
		t.Fatal(err)
	}

	if !test.Test.Equal(testTime.Time) {
		t.Fatalf("Dates are not equal: %v, %v", test.Test, testTime)
	}
}

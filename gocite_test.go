package gocite_test

import (
	"testing"

	"github.com/ThomasK81/gocite"
)

type testpair struct {
	input                  string
	outputSplit            gocite.CTSURN
	outputRange, outputCTS bool
}

var tests = []testpair{
	{input: "urn:cts:collection:workgroup.work:1-27", outputSplit: gocite.CTSURN{ID: "urn:cts:collection:workgroup.work:1-27", Base: "urn", Protocol: "cts", Namespace: "collection", Work: "workgroup.work", Passage: "1-27"}, outputRange: true, outputCTS: true},
	{input: "urn:cts:collection:workgroup.work:27.3", outputSplit: gocite.CTSURN{ID: "urn:cts:collection:workgroup.work:27.3", Base: "urn", Protocol: "cts", Namespace: "collection", Work: "workgroup.work", Passage: "27.3"}, outputRange: false, outputCTS: true},
	{input: "not:cts:collection:workgroup.work:27.3", outputSplit: gocite.CTSURN{ID: "not:cts:collection:workgroup.work:27.3", InValid: true}, outputRange: false, outputCTS: false}}

var firstPassage = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:1",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is the first node.",
	},
	Index: 0,
	First: gocite.PassLoc{PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Last:  gocite.PassLoc{PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
	Prev:  gocite.PassLoc{PassageID: "", Index: nil},
	Next:  gocite.PassLoc{PassageID: "urn:cts:collection:workgroup.work:2-4", Index: 1},
}

var testcorpus = gocite.Work{
	WorkID:   "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{},
	Ordered:  true}

func TestSplitCTS(t *testing.T) {
	for _, pair := range tests {
		v := gocite.SplitCTS(pair.input)
		if v != pair.outputSplit {
			t.Error(
				"For", pair.input,
				"expected", pair.outputSplit,
				"got", v,
			)
		}
	}
}

func TestIsRange(t *testing.T) {
	for _, pair := range tests {
		v := gocite.IsRange(pair.input)
		if v != pair.outputRange {
			t.Error(
				"For", pair.input,
				"expected", pair.outputRange,
				"got", v,
			)
		}
	}
}

func TestIsCTSURN(t *testing.T) {
	for _, pair := range tests {
		v := gocite.IsCTSURN(pair.input)
		if v != pair.outputCTS {
			t.Error(
				"For", pair.input,
				"expected", pair.outputCTS,
				"got", v,
			)
		}
	}
}

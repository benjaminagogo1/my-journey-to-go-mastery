package main

import (
	"testing"
)

func TestArticle(t *testing.T) {
	input := fixArticle("There it was. A amazing rock!")
	expected := "There it was. An amazing rock!"
	if input != expected {
		t.Errorf("fixArticle failed: got %s, expected %s", input, expected)
	}
}

func TestCase(t *testing.T) {
	input := caseTransFromation("I should stop SHOUTING (low)")
	expected := "I should stop shouting"
	if input != expected {
		t.Errorf("case failed: got %s, expected %s", input, expected)
	}
}

func TestConversion(t *testing.T) {
	input := conversion("1E (hex) files were added")
	expected := "30 files were added"
	if input != expected {
		t.Errorf("conversion failed: got %s, expected %s", input, expected)
	}
}

func TestSingleQuote(t *testing.T) {
	input := fixSingleQuote("As Elton John said: ' I am the most well-known homosexual in the world '")
	expected := "As Elton John said: 'I am the most well-known homosexual in the world'"
	if input != expected {
		t.Errorf("fixsinglequote failed: got %s, expected %s", input, expected)
	}
}

func TestFixPunctuation(t *testing.T) {
	input := fixPunctuation("I was thinking ... You were right")
	expected := "I was thinking... You were right"
	if input != expected {
		t.Errorf("fixPunctuation failed: got %s, expected %s", input, expected)
	}
}

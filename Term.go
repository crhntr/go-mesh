package gomesh

import "encoding/xml"

type Term struct {
	XMLName xml.Name `xml:"Term"`
	UI      UI       `xml:"TermUI"`

	ConceptPreferred YN       `xml:"ConceptPreferredTermYN,attr"`
	IsPermuted       YN       `xml:"IsPermutedTermYN,attr"`
	RecordPreferred  YN       `xml:"RecordPreferredTermYN,attr"`
	Created          Date     `xml:"DateCreated"`
	ThesaurusIDs     []string `xml:"ThesaurusIDlist>ThesaurusID"`

	String string `xml:"String"`
}

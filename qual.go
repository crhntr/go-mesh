package gomesh

import (
	"encoding/xml"
	"io"
)

type QualifierRecord struct {
	UI          UI     `xml:"QualifierUI"`
	Name        string `xml:"QualifierName>String"`
	Created     Date   `xml:"DateCreated"`
	Revised     Date   `xml:"DateRevised"`
	Established Date   `xml:"DateEstablished"`
	Annotation  string
	HistoryNote string
	TreeNumbers []string `xml:"TreeNumberList>TreeNumber"`
}

func ParseQualifierRecordSet(r io.Reader) (<-chan QualifierRecord, chan error) {
	dec := xml.NewDecoder(r)
	qrc := make(chan QualifierRecord)
	errc := make(chan error)
	go func() {
		for {
			tok, err := dec.Token()
			if err != nil {
				errc <- err
				close(qrc)
				return
			}

			switch t := tok.(type) {
			case xml.StartElement:
				if t.Name.Local == "QualifierRecord" {
					qr := QualifierRecord{}

					if err := dec.DecodeElement(&qr, &t); err != nil {
						errc <- err
						close(qrc)
						return
					}

					qrc <- qr
				}
			}
		}
	}()
	return qrc, errc
}

/*

<QualifierRecord>
  <QualifierUI>Q000000981</QualifierUI>
  <QualifierName>
   <String>diagnostic imaging</String>
  </QualifierName>
  <DateCreated>
   <Year>2016</Year>
   <Month>06</Month>
   <Day>29</Day>
  </DateCreated>
  <DateRevised>
   <Year>2016</Year>
   <Month>06</Month>
   <Day>08</Day>
  </DateRevised>
  <DateEstablished>
   <Year>2017</Year>
   <Month>01</Month>
   <Day>01</Day>
  </DateEstablished>
  <Annotation>subheading only; coordinate with specific  imaging technique if pertinent
  </Annotation>
  <HistoryNote>2017(1967)
  </HistoryNote>
  <TreeNumberList>
   <TreeNumber>Y04.010</TreeNumber>
  </TreeNumberList>
  <ConceptList>
   <Concept PreferredConceptYN="Y">
    <ConceptUI>M000614856</ConceptUI>
    <ConceptName>
     <String>diagnostic imaging</String>
    </ConceptName>
    <ScopeNote>Used for the visualization of an anatomical structure or for the diagnosis of disease.  Commonly used imaging techniques include radiography, radionuclide imaging, thermography, tomography, and ultrasonography
    </ScopeNote>
    <ConceptRelationList>
     <ConceptRelation RelationName="NRW">
     <Concept1UI>M000614856</Concept1UI>
     <Concept2UI>M0030904</Concept2UI>
     </ConceptRelation>
     <ConceptRelation RelationName="NRW">
     <Concept1UI>M000614856</Concept1UI>
     <Concept2UI>M0030903</Concept2UI>
     </ConceptRelation>
     <ConceptRelation RelationName="NRW">
     <Concept1UI>M000614856</Concept1UI>
     <Concept2UI>M0030734</Concept2UI>
     </ConceptRelation>
     <ConceptRelation RelationName="NRW">
     <Concept1UI>M000614856</Concept1UI>
     <Concept2UI>M0030733</Concept2UI>
     </ConceptRelation>
    </ConceptRelationList>
    <TermList>
     <Term  ConceptPreferredTermYN="Y"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="Y">
      <TermUI>T000895609</TermUI>
      <String>diagnostic imaging</String>
      <DateCreated>
       <Year>2016</Year>
       <Month>02</Month>
       <Day>19</Day>
      </DateCreated>
      <Abbreviation>DG</Abbreviation>
      <EntryVersion>DIAG IMAGE</EntryVersion>
     </Term>
    </TermList>
   </Concept>
   <Concept PreferredConceptYN="N">
    <ConceptUI>M0030904</ConceptUI>
    <ConceptName>
     <String>ultrasound</String>
    </ConceptName>
    <ConceptRelationList>
     <ConceptRelation RelationName="NRW">
     <Concept1UI>M000614856</Concept1UI>
     <Concept2UI>M0030904</Concept2UI>
     </ConceptRelation>
    </ConceptRelationList>
    <TermList>
     <Term  ConceptPreferredTermYN="Y"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="N">
      <TermUI>T061379</TermUI>
      <String>ultrasound</String>
     </Term>
    </TermList>
   </Concept>
   <Concept PreferredConceptYN="N">
    <ConceptUI>M0030903</ConceptUI>
    <ConceptName>
     <String>ultrasonography</String>
    </ConceptName>
    <ScopeNote>Used with organs and regions for ultrasonic imaging and with diseases for ultrasonic diagnosis. Does not include ultrasonic therapy.
    </ScopeNote>
    <ConceptRelationList>
     <ConceptRelation RelationName="NRW">
     <Concept1UI>M000614856</Concept1UI>
     <Concept2UI>M0030903</Concept2UI>
     </ConceptRelation>
    </ConceptRelationList>
    <TermList>
     <Term  ConceptPreferredTermYN="Y"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="N">
      <TermUI>T061375</TermUI>
      <String>ultrasonography</String>
     </Term>
     <Term  ConceptPreferredTermYN="N"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="N">
      <TermUI>T061377</TermUI>
      <String>ultrasonic diagnosis</String>
     </Term>
     <Term  ConceptPreferredTermYN="N"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="N">
      <TermUI>T061378</TermUI>
      <String>echotomography</String>
     </Term>
     <Term  ConceptPreferredTermYN="N"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="N">
      <TermUI>T061376</TermUI>
      <String>echography</String>
     </Term>
    </TermList>
   </Concept>
   <Concept PreferredConceptYN="N">
    <ConceptUI>M0030734</ConceptUI>
    <ConceptName>
     <String>radionuclide imaging</String>
    </ConceptName>
    <ScopeNote>Used for radionuclide imaging of any anatomical structure, or for the diagnosis of disease.
    </ScopeNote>
    <ConceptRelationList>
     <ConceptRelation RelationName="NRW">
     <Concept1UI>M000614856</Concept1UI>
     <Concept2UI>M0030734</Concept2UI>
     </ConceptRelation>
    </ConceptRelationList>
    <TermList>
     <Term  ConceptPreferredTermYN="Y"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="N">
      <TermUI>T061163</TermUI>
      <String>radionuclide imaging</String>
     </Term>
     <Term  ConceptPreferredTermYN="N"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="N">
      <TermUI>T061164</TermUI>
      <String>radioisotope scanning</String>
     </Term>
    </TermList>
   </Concept>
   <Concept PreferredConceptYN="N">
    <ConceptUI>M0030733</ConceptUI>
    <ConceptName>
     <String>radiography</String>
    </ConceptName>
    <ScopeNote>Used with organs, regions, and diseases for x-ray examinations.
    </ScopeNote>
    <ConceptRelationList>
     <ConceptRelation RelationName="NRW">
     <Concept1UI>M000614856</Concept1UI>
     <Concept2UI>M0030733</Concept2UI>
     </ConceptRelation>
    </ConceptRelationList>
    <TermList>
     <Term  ConceptPreferredTermYN="Y"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="N">
      <TermUI>T061158</TermUI>
      <String>radiography</String>
     </Term>
     <Term  ConceptPreferredTermYN="N"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="N">
      <TermUI>T061162</TermUI>
      <String>X-ray image</String>
     </Term>
     <Term  ConceptPreferredTermYN="N"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="N">
      <TermUI>T061161</TermUI>
      <String>X-ray diagnosis</String>
     </Term>
     <Term  ConceptPreferredTermYN="N"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="N">
      <TermUI>T061159</TermUI>
      <String>roentgenography</String>
     </Term>
     <Term  ConceptPreferredTermYN="N"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="N">
      <TermUI>T061160</TermUI>
      <String>X-ray</String>
     </Term>
    </TermList>
   </Concept>
  </ConceptList>
 </QualifierRecord>

*/

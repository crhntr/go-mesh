package gomesh

type Concept struct {
	UI                     UI     `xml:"ConceptUI"`
	Name                   string `xml:"ConceptName>String"`
	RegistryNumber         string
	CASN1Name              string
	ConceptRelations       []ConceptRelation `xml:"ConceptRelationList>ConceptRelation"`
	Prefered               YN                `xml:"PreferredConceptYN,attr"`
	ScopeNote              string
	RelatedRegistryNumbers []string `xml:"RelatedRegistryNumberList>RelatedRegistryNumber"`
}

type ConceptRelation struct {
	Name string `xml:"RelationName,attr"`
	C1   UI
	C2   UI
}

/*
<Concept PreferredConceptYN="N">
    <ConceptUI>M0030204</ConceptUI>
    <ConceptName>
     <String>agenesis</String>
    </ConceptName>
    <ConceptRelationList>
     <ConceptRelation RelationName="NRW">
     <Concept1UI>M0030203</Concept1UI>
     <Concept2UI>M0030204</Concept2UI>
     </ConceptRelation>
    </ConceptRelationList>
    <TermList>
     <Term  ConceptPreferredTermYN="Y"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="N">
      <TermUI>T060546</TermUI>
      <String>agenesis</String>
     </Term>
    </TermList>
   </Concept>


   <Concept PreferredConceptYN="Y">
    <ConceptUI>M0000004</ConceptUI>
    <ConceptName>
     <String>Abbreviations as Topic</String>
    </ConceptName>
    <ScopeNote>Shortened forms of written words or phrases used for brevity.
    </ScopeNote>
    <ConceptRelationList>
     <ConceptRelation RelationName="NRW">
     <Concept1UI>M0000004</Concept1UI>
     <Concept2UI>M0511063</Concept2UI>
     </ConceptRelation>
    </ConceptRelationList>
    <TermList>
     <Term  ConceptPreferredTermYN="Y"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="Y">
      <TermUI>T698652</TermUI>
      <String>Abbreviations as Topic</String>
      <DateCreated>
       <Year>2007</Year>
       <Month>05</Month>
       <Day>31</Day>
      </DateCreated>
      <ThesaurusIDlist>
       <ThesaurusID>NLM (2008)</ThesaurusID>
      </ThesaurusIDlist>
     </Term>
    </TermList>
   </Concept>
   <Concept PreferredConceptYN="N">
    <ConceptUI>M0511063</ConceptUI>
    <ConceptName>
     <String>Acronyms as Topic</String>
    </ConceptName>
    <ConceptRelationList>
     <ConceptRelation RelationName="NRW">
     <Concept1UI>M0000004</Concept1UI>
     <Concept2UI>M0511063</Concept2UI>
     </ConceptRelation>
    </ConceptRelationList>
    <TermList>
     <Term  ConceptPreferredTermYN="Y"  IsPermutedTermYN="N"  LexicalTag="NON"  RecordPreferredTermYN="N">
      <TermUI>T701041</TermUI>
      <String>Acronyms as Topic</String>
      <DateCreated>
       <Year>2007</Year>
       <Month>06</Month>
       <Day>29</Day>
      </DateCreated>
      <ThesaurusIDlist>
       <ThesaurusID>NLM (2008)</ThesaurusID>
      </ThesaurusIDlist>
     </Term>
    </TermList>
   </Concept>
*/

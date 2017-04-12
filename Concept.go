package mesh

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

package gomesh

type PharmacologicalAction struct {
	UI   UI     `xml:"DescriptorReferredTo>DescriptorUI"`
	Name string `xml:"DescriptorReferredTo>DescriptorName>String"`

	Substances []Substance `xml:"PharmacologicalActionSubstanceList>Substance"`
}

type Substance struct {
	UI   UI     `xml:"RecordUI"`
	Name string `xml:"RecordName>String"`
}

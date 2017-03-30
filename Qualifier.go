package gomesh

type AllowableQualifier struct {
	UI   UI     `xml:"QualifierReferredTo>QualifierUI"`
	Name string `xml:"QualifierReferredTo>QualifierName>String"`

	Abbreviation string
}

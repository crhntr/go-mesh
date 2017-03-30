package gomesh

type UI string

type UIType int

const (
	UnknownUI UIType = iota
	ConceptUI
	QualifierUI
	TermUI
	DescriptorUI
	// RecordUI // C or D
)

func (ui UI) Type() UIType {
	if uit, exists := map[byte]UIType{
		'Q': QualifierUI,
		'M': ConceptUI,
		'T': TermUI,
		'D': DescriptorUI,
	}[(string(ui))[0]]; !exists {
		return UnknownUI
	} else {
		return uit
	}
}

package gomesh

type UI string

type UIType int

const (
	UnknownUI    UIType = iota
	DescriptorUI        // 'D'
	QualifierUI         // 'Q'
	TermUI              // 'T'
	ConceptUI           // 'M'
	// RecordUI // C or D
)

func (ui UI) Type() UIType {
	if len(ui) < 2 {
		return UnknownUI
	}
	str := string(ui)

	switch {
	case str[0] == 'D', str[0] == '*' && str[1] == 'D':
		return DescriptorUI
	case str[0] == 'Q', str[0] == '*' && str[1] == 'Q':
		return QualifierUI
	case str[0] == 'T':
		return TermUI
	case str[0] == 'M':
		return ConceptUI
	default:
		return UnknownUI
	}
}

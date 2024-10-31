package penny

type Penny struct {
	Texts []string
}

func (p *Penny) AddText(text string) {
	p.Texts = append(p.Texts, text)
}

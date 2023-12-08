package pipe

type IPipe interface {
	Write(id, value string) bool
	Read(id string) (string, bool)
}

type Pipe struct {
	m map[string]string
}

func NewPipe() *Pipe {
	return &Pipe{
		m: make(map[string]string),
	}
}

func (p *Pipe) Write(key, value string) bool {
	if _, ok := p.m[key]; ok {
		return false
	}
	p.m[key] = value
	return true
}

func (p *Pipe) Read(key string) (string, bool) {
	if value, ok := p.m[key]; ok {
		delete(p.m, key)
		return value, true
	}
	return "", false
}

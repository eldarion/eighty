package engine

type Engine struct {
	Vhosts map[string]*Vhost
}

func New() *Engine {
	e := &Engine{}
	e.Vhosts = make(map[string]*Vhost)
	e.Vhosts["_eighty"] = &Vhost{Mode: "management"}
	return e
}

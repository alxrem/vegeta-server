package vegeta

import (
	vegeta "github.com/tsenart/vegeta/lib"
	"github.com/tsenart/vegeta/lib/plot"
	"io"
)

type Plot struct {
	p *plot.Plot
}

func NewPlot() *Plot {
	return &Plot{
		plot.New(),
	}
}

func (p *Plot) Add(r *vegeta.Result) {
	_ = p.p.Add(r)
}

func (p *Plot) NewReporter() vegeta.Reporter {
	return func(w io.Writer) (err error) {
		_, err = p.p.WriteTo(w)
		return err
	}
}

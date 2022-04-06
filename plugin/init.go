package plugin

import (
	"github.com/allentom/harukap"
	"github.com/davidbyttow/govips/v2/vips"
)

type InitPlugin struct {
}

func (p *InitPlugin) OnInit(e *harukap.HarukaAppEngine) error {
	vips.Startup(nil)
	return nil
}

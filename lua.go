package host

import (
	"fmt"
	"github.com/vela-ssoc/vela-kit/vela"
	"github.com/vela-ssoc/vela-kit/lua"
)

var xEnv vela.Environment

func (h *Host) String() string                         { return fmt.Sprintf("%p", h) }
func (h *Host) Type() lua.LValueType                   { return lua.LTObject }
func (h *Host) AssertFloat64() (float64, bool)         { return 0, false }
func (h *Host) AssertString() (string, bool)           { return "", false }
func (h *Host) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (h *Host) Peek() lua.LValue                       { return h }

func WithEnv(env vela.Environment) {
	xEnv = env
	env.Set("host", &Host{})
}

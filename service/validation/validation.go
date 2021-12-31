package validation

import (
	"github.com/seed95/shortening/build/messages"
	"github.com/seed95/shortening/pkg/derrors"
	"github.com/seed95/shortening/pkg/log"
)

func (h *handler) Alias(alias string) error {
	if len(alias) < h.cfg.AliasMinLength {
		h.logger.Error(&log.Field{
			Section:  "service.validation",
			Function: "Alias",
			Params:   map[string]interface{}{"alias": alias},
			Message:  h.translator.Translate(messages.InvalidAliasLength),
		})

		return derrors.New(derrors.Invalid, messages.InvalidAliasLength)
	}
	return nil
}

package grip

import (
	"github.com/mschoenlaub/grip/level"
	"github.com/mschoenlaub/grip/send"
	"github.com/mschoenlaub/grip/slogger"
)

// NewJournalerFromSlogger takes a slogger logging instance and
// returns a functionally equivalent Jouranler instance.
func NewJournalerFromSlogger(logger *slogger.Logger) (Journaler, error) {
	l := send.LevelInfo{Default: level.Debug, Threshold: level.Debug}

	j := NewJournaler(logger.Name)
	if err := j.GetSender().SetLevel(l); err != nil {
		return nil, err
	}

	sender, err := send.NewMultiSender(logger.Name, l, logger.Appenders)
	if err != nil {
		return nil, err
	}

	if err := j.SetSender(sender); err != nil {
		return nil, err
	}

	return j, nil
}

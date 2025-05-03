package logger

import "testing"

func TestInitLogger(t *testing.T) {
	InitLogger()
	Log.Info().Msg("HERE")
	Log.Info().Str("HERE AAS ASAS", "AA").Msg("ASA")
	Log.Info().Msgf("HERE %s", "AA ASASAS")
}

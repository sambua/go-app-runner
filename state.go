package application

import config "config"

type appState struct {
	state    string
	isDebug  bool
	logLevel string
}

func newAppState() (*appState, error) {
	state, err := config.NewState()
	if err != nil {
		return nil, err
	}

	return &appState{
		state:    state.Env,
		isDebug:  state.Debug,
		logLevel: state.DebugLevel,
	}, err
}

func (a *appState) IsDebug() bool {
	return a.isDebug
}

func (a *appState) LogLevel() string {
	return a.logLevel
}

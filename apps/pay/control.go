package pay

import "github.com/mavryk-network/mavbake/ami"

func (app *Mavpay) Start(args ...string) (int, error) {
	return ami.StartApp(app.GetPath(), args...)
}

func (app *Mavpay) Stop(args ...string) (int, error) {
	return ami.StopApp(app.GetPath(), args...)
}

func (app *Mavpay) Remove(all bool, args ...string) (int, error) {
	return ami.RemoveApp(app.GetPath(), all, args...)
}

func (app *Mavpay) Execute(args ...string) (int, error) {
	return ami.Execute(app.GetPath(), args...)
}

func (app *Mavpay) ExecuteGetOutput(args ...string) (string, int, error) {
	return ami.ExecuteGetOutput(app.GetPath(), args...)
}

func (app *Mavpay) ExecuteWithOutputChannel(outputChannel chan<- string, args ...string) (int, error) {
	return ami.ExecuteWithOutputChannel(app.GetPath(), outputChannel, args...)
}

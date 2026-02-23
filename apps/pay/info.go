package pay

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/mavryk-network/mavbake/ami"
	"github.com/mavryk-network/mavbake/apps/base"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type Info struct {
	base.InfoBase
	Services map[string]base.AmiServiceInfo `json:"services"`
	Type     string                         `json:"type"`
	Version  string                         `json:"version"`
}

func (i *Info) UnmarshalJSON(data []byte) error {
	type Alias Info
	aux := &struct {
		Services json.RawMessage `json:"services"`
		*Alias
	}{
		Alias: (*Alias)(i),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if err := base.UnmarshalIfNotEmptyArray(aux.Services, &i.Services); err != nil {
		return err
	}

	return nil
}

type InfoCollectionOptions struct {
	Timeout  int
	Services bool
}

func (infoCollectionOptions *InfoCollectionOptions) toAmiArgs() []string {
	args := make([]string, 0)
	return args
}

func (nico *InfoCollectionOptions) All() bool {
	return true
}

func (app *Mavpay) getInfoCollectionOptions(optionsJson []byte) *InfoCollectionOptions {
	result := &InfoCollectionOptions{}
	json.Unmarshal(optionsJson, result)
	return result
}

func (app *Mavpay) GetAvailableInfoCollectionOptions() []base.AmiInfoCollectionOption {
	result := make([]base.AmiInfoCollectionOption, 0)
	options := InfoCollectionOptions{}
	val := reflect.ValueOf(options)

	for i := 0; i < val.NumField(); i++ {
		result = append(result, base.AmiInfoCollectionOption{
			Name: strings.ToLower(val.Type().Field(i).Name),
			Type: strings.ToLower(val.Type().Field(i).Type.Name()),
		})
	}
	return result
}

func (app *Mavpay) GetInfoFromOptions(options *InfoCollectionOptions) (Info, error) {
	args := options.toAmiArgs()
	infoBytes, _, err := ami.ExecuteInfo(app.GetPath(), args...)
	if err != nil {
		failedInfo := Info{
			InfoBase: base.GenerateFailedInfo(string(infoBytes), err),
		}
		return failedInfo, fmt.Errorf("failed to collect app info (%s)", err.Error())
	}

	info, err := base.ParseInfoOutput[Info](infoBytes)
	if err != nil {
		return Info{InfoBase: base.GenerateFailedInfo(string(infoBytes), err)}, err
	}
	return info, nil
}

func (app *Mavpay) GetInfo(optionsJson []byte) (any, error) {
	return app.GetInfoFromOptions(app.getInfoCollectionOptions(optionsJson))
}

func (app *Mavpay) GetServiceInfo() (map[string]base.AmiServiceInfo, error) {
	result := map[string]base.AmiServiceInfo{}

	info, err := app.GetInfoFromOptions(&InfoCollectionOptions{Services: true})
	if err != nil {
		return result, err
	}

	return info.Services, err
}

func (app *Mavpay) IsServiceStatus(id string, status string) (bool, error) {
	return base.IsServiceStatus(app, id, status)
}

func (app *Mavpay) IsAnyServiceStatus(status string) (bool, error) {
	return base.IsAnyServiceStatus(app, status)
}

func (app *Mavpay) PrintInfo(optionsJson []byte) error {
	mavpayInfoRaw, err := app.GetInfo(optionsJson)
	if err != nil {
		return err
	}
	mavpayInfo, ok := mavpayInfoRaw.(Info)
	if !ok {
		return fmt.Errorf("invalid mavpay info type")
	}

	mavpayTable := table.NewWriter()
	mavpayTable.SetStyle(table.StyleLight)
	mavpayTable.SetColumnConfigs([]table.ColumnConfig{{Number: 1, Align: text.AlignLeft}, {Number: 2, Align: text.AlignLeft}})
	mavpayTable.SetOutputMirror(os.Stdout)
	mavpayTable.AppendHeader(table.Row{app.GetLabel(), app.GetLabel()}, table.RowConfig{AutoMerge: true})

	mavpayTable.AppendRow(table.Row{"Status", mavpayInfo.Status})
	mavpayTable.AppendRow(table.Row{"Status Level", mavpayInfo.Level})

	mavpayTable.AppendSeparator()
	mavpayTable.AppendRow(table.Row{"Services", "Services"}, table.RowConfig{AutoMerge: true})
	mavpayTable.AppendSeparator()
	mavpayTable.AppendRow(table.Row{"Name", "Status (Started)"})
	mavpayTable.AppendSeparator()

	for k, v := range mavpayInfo.Services {
		mavpayTable.AppendRow(table.Row{k, fmt.Sprintf("%v (%v)", v.Status, v.Started)})
	}

	mavpayTable.Render()
	return nil
}

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/mavryk-network/mavbake/ami"
	"github.com/mavryk-network/mavbake/cli"
	"github.com/mavryk-network/mavbake/constants"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type bbTextFormatter struct {
	log.TextFormatter
}

func (f *bbTextFormatter) Format(entry *log.Entry) ([]byte, error) {
	result := entry.Time.Format("15:04:05")
	result = result + " [" + strings.ToUpper(string(entry.Level.String())) + "] (mavbake) " + entry.Message + "\n"
	for k, v := range entry.Data {
		result = result + k + "=" + fmt.Sprint(v) + "\n"
	}
	return []byte(result), nil
}

type bbJsonFormatter struct {
	log.JSONFormatter
}

func (f *bbJsonFormatter) Format(entry *log.Entry) ([]byte, error) {
	//strconv.FormatInt(entry.Time.Unix(), 10)
	l, err := f.JSONFormatter.Format(entry)
	if err != nil {
		return []byte{}, err
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(l, &result)
	if err != nil {
		return []byte{}, err
	}
	delete(result, "time")
	result["timestamp"] = strconv.FormatInt(entry.Time.Unix(), 10)
	result["module"] = "mavbake"
	resultLog, err := json.Marshal(result)
	resultLog = append(resultLog, byte('\n'))
	return resultLog, err
}

var (
	RootCmd = &cobra.Command{
		Use:   "mavbake",
		Short: "mavbake CLI",
		Long: `mavbake CLI
Copyright © 2024 Mavryk Dynamics
`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {

			if cmd.Flags().Changed("path") {
				cli.BBdir, _ = cmd.Flags().GetString("path")
			} else {
				bbDir := os.Getenv("BB_DIR")
				if bbDir != "" {
					cli.BBdir = bbDir
				}
			}

			switch level, _ := cmd.Flags().GetString("log-level"); level {
			case "trace":
				log.SetLevel(log.TraceLevel)
				cli.LogLevel = "trace"
			case "debug":
				log.SetLevel(log.DebugLevel)
				cli.LogLevel = "debug"
			case "warn":
				log.SetLevel(log.WarnLevel)
				cli.LogLevel = "warn"
			case "error":
				log.SetLevel(log.ErrorLevel)
				cli.LogLevel = "error"
			default:
				log.SetLevel(log.InfoLevel)
			}
			log.Trace("Log level set to '" + cli.LogLevel + "'")

			cli.IsRemoteInstance, _ = cmd.Flags().GetBool("remote-instance")
			if cli.IsRemoteInstance {
				remoteVars, _ := cmd.Flags().GetString("remote-instance-vars")
				if remoteVars != "" {
					vars := strings.Split(remoteVars, ";")
					for _, _var := range vars {
						kvPair := strings.Split(_var, "=")
						if len(kvPair) != 2 {
							continue
						}
						ami.REMOTE_VARS[kvPair[0]] = kvPair[1]
					}
				}
			}

			format, _ := cmd.Flags().GetString("output-format")
			// if cli.IsRemoteInstance { // override for remote instance
			// 	format = "json"
			// }

			switch format {
			case "json":
				cli.JsonLogFormat = true
				log.SetFormatter(&bbJsonFormatter{})
				log.Trace("Output format set to 'json'")
			case "text":
				log.SetFormatter(&bbTextFormatter{})
				log.Trace("Output format set to 'text'")
			default:
				if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) == 0 {
					log.SetFormatter(&bbJsonFormatter{})
					log.Trace("Output format automatically set to 'json'")
				} else {
					log.SetFormatter(&bbTextFormatter{})
					log.Trace("Output format automatically set to 'text'")
				}
			}

			// init ami options
			ami.SetOptions(ami.Options{
				LogLevel:             cli.LogLevel,
				JsonLogFormat:        cli.JsonLogFormat,
				DoNotCheckForLocator: cli.IsRemoteInstance,
			})

		},
	}
)

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.PersistentFlags().StringP("path", "p", constants.DefaultBBDirectory, "Path to mavpay instance")
	RootCmd.PersistentFlags().StringP("output-format", "o", "auto", "Sets output log format (json/text/auto)")
	RootCmd.PersistentFlags().StringP("log-level", "l", "info", "Sets output log format (json/text/auto)")
	RootCmd.PersistentFlags().Bool("remote-instance", false, "Tells mavbake to operate in remote-instance mode")
	RootCmd.PersistentFlags().MarkHidden("remote-instance")
	RootCmd.PersistentFlags().String("remote-instance-vars", "", "Tells mavbake to which remote vars to set (available only with remote-instance)")
	RootCmd.PersistentFlags().MarkHidden("remote-instance-vars")
	RootCmd.PersistentFlags().SetInterspersed(false)
}

func ExecuteTest(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimSpace(buf.String()), err
}

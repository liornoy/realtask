/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"github.com/spf13/cobra"
	"github.com/liornoy/realtask/scheduler"
)


const defaultSchedulerListenPort = "9009"
// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start starts the scheduler",
	Long:  fmt.Sprintf(`start starts the scheduler. set port with -port, defaults to %s\n`, defaultSchedulerListenPort),
	Run: func(cmd *cobra.Command, args []string) {
		p := cmd.Flag("port").Value.String()
		err := validatePort(p)
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}

		err=startScheduler(p)
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("started scheduler on port: ", p)
		os.Setenv("REALTASK_SCHEDULER_PORT", p)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("port", "p",defaultSchedulerListenPort,"set the listening port for the scheduler")
}

// checkPort 
func validatePort(p string) error{
	_,err := strconv.Atoi(p)
	if err!=nil{
		return fmt.Errorf("invalid port value")
	}
	
	return nil
}
func startScheduler(p string) error{
	return scheduler.New(p)
}


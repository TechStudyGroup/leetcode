package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tosone/logging"
	"github.com/unknwon/com"

	"6leetcode/cmd/crawler"
	"6leetcode/cmd/csvgen"
	"6leetcode/cmd/gen"
	"6leetcode/cmd/version"
	"6leetcode/common/table"
)

// RootCmd represents the base command when called without any sub commands
var RootCmd = &cobra.Command{
	Short: "Leetcode tool.",
	Long:  `Leetcode tool.`,
}

const DefaultConfig = "/etc/6leetcode/config.yml"

func init() {
	var config string
	RootCmd.PersistentFlags().StringVarP(&config, "config", "c", "./config.yml", "config file")

	var crawlerCmd = &cobra.Command{
		Use:   "crawler",
		Short: "Travel all of the leetcode problems info.",
		Long:  `Travel all of the leetcode problems info.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(_ *cobra.Command, _ []string) {
			var err error
			if err = initConfig(config); err != nil {
				logging.Errorf("Init config with error: %+v", err)
				return
			}
			if err = initTable(Questions); err != nil {
				logging.Errorf("Init table with error: %+v", err)
				return
			}
			if err = crawler.Initialize(); err != nil {
				logging.Errorf("Init crawler with error: %+v", err)
				return
			}
		},
	}
	RootCmd.AddCommand(crawlerCmd) // crawler commander

	var csvCmd = &cobra.Command{
		Use:   "csv",
		Short: "Travel all of the leetcode problems info.",
		Long:  `Travel all of the leetcode problems info.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(_ *cobra.Command, _ []string) {
			var err error
			if err = initConfig(config); err != nil {
				logging.Errorf("Init config with error: %+v", err)
				return
			}
			if err = initTable(Questions); err != nil {
				logging.Errorf("Init table with error: %+v", err)
				return
			}
			if err = csvgen.Initialize(); err != nil {
				logging.Errorf("Init crawler with error: %+v", err)
				return
			}
		},
	}
	RootCmd.AddCommand(csvCmd) // crawler commander

	var genCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generate readme.",
		Long:  `Generate readme.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(_ *cobra.Command, _ []string) {
			var err error
			if err = initConfig(config); err != nil {
				logging.Error(err)
			}
			if err = gen.Initialize(); err != nil {
				logging.Errorf("Got error: %+v", err)
				return
			}
		},
	}
	RootCmd.AddCommand(genCmd)

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Get version.",
		Long:  `The version that build detail information.`,
		Run: func(_ *cobra.Command, _ []string) {
			version.Initialize()
		},
	}
	RootCmd.AddCommand(versionCmd) // version commander

	RootCmd.Use = viper.GetString("AppName")
}

func initTable(questions embed.FS) (err error) {
	if err = table.Initialize(); err != nil {
		logging.Errorf("Got error: %+v", err)
		return
	}
	return
}

func initConfig(config string) (err error) {
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	if com.IsFile(config) {
		viper.SetConfigFile(config)
	} else if config == "./config.yml" && com.IsFile(DefaultConfig) {
		viper.SetConfigFile(DefaultConfig)
	}
	if err := viper.ReadInConfig(); err != nil {
		logging.Error("Cannot find the specified config file.")
	}

	if os.Getenv("DATABASE_ENGINE") != "" {
		viper.Set("Database.Engine", "sqlite3")
		viper.Set("Database.Path", "6leetcode.db")
	}
	if os.Getenv("LOGIN_NAME") != "" {
		viper.Set("LOGIN_NAME", os.Getenv("LOGIN_NAME"))
	}
	if os.Getenv("LOGIN_PASSWORD") != "" {
		viper.Set("LOGIN_PASSWORD", os.Getenv("LOGIN_PASSWORD"))
	}
	return
}

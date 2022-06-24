package cmd

import (
	"fmt"
	"log"

	"github.com/jlvihv/dbtorust/controller"
	"github.com/jlvihv/dbtorust/utils"
	"github.com/spf13/cobra"
)

var (
	db      string
	table   string
	clip    bool
	file    string
)

var rootCmd = &cobra.Command{
	Use:   "dbtogo",
	Short: "将数据库中的表转换为 go 结构体",
	Run: func(_ *cobra.Command, _ []string) {
		run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&db, "db", "", "", "数据库名")
	rootCmd.Flags().StringVarP(&table, "table", "", "", "表名")

	rootCmd.Flags().BoolVarP(&clip, "clip", "", false, "输出到系统剪贴板")
	rootCmd.Flags().StringVarP(&file, "file", "", "", "输出到文件")

	rootCmd.Flags().StringVarP(utils.ConfigPath(), "config", "", "", "指定配置文件所在位置")
}

func run() {
	if len(db) == 0 || len(table) == 0 {
		fmt.Printf("db: %s, table: %s\n", db, table)
		fmt.Println("数据库名与表名不得为空")
		return
	}
	c := controller.NewController()

	c.GetColumns(db, table).ConvertToStructColumns()

	c.ToUpperCamelCase().Generate()

	if !clip && len(file) == 0 {
		c.Stdout()
	}
	if clip {
		c.Clipboard()
	}
	if len(file) != 0 {
		c.File(file)
	}
}

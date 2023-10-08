package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"oneindex-img-uploader/cmd/uploader"
	"os"
)

var (
	cfgFile string
	url     string
)

func init() {
	uploader.InitConfig(rootCmd)
}

var rootCmd = &cobra.Command{
	Use:   "uploader [options] <local_path> [[local_path2] [local_path3] ...]",
	Short: "oneindex img uploader is a tool to upload local image to oneindex images",
	Long: `oneindex img uploader is a tool to upload local image to oneindex images, 
                if the local path is a dir, then uploader will walk through sub dir, upload all possible image files
                if the local path is a file, then uploader will upload directly.
                This program is mainly designed for Typora custom image function`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := uploader.InitUploader(); err != nil {
			log.Printf("init uploader failed: %s", err.Error())
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}
		vp := uploader.CollectValidatePaths(args)
		pairs := uploader.Upload(vp)
		uploader.PrintResultByOrder(vp, pairs)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("err: %s", err.Error())
		os.Exit(1)
	}
}

package main

import (
	"context"

	_ "github.com/going/porter/backend/all" // import all backends
	"github.com/rclone/rclone/cmd"
	"github.com/rclone/rclone/fs/operations"
	"github.com/rclone/rclone/fs/sync"
	_ "github.com/rclone/rclone/lib/plugin" // import plugins
)

func main() {
	fsrc, srcFileName, fdst := cmd.NewFsSrcFileDst(args)
	if len(fsrc.Root()) > 7 && "isFile:" == fsrc.Root()[0:7] {
		srcFileName = fsrc.Root()[7:]
	}
	cmd.Run(true, true, command, func() error {
		if srcFileName == "" {
			return sync.CopyDir(context.Background(), fdst, fsrc, createEmptySrcDirs)
		}
		return operations.CopyFile(context.Background(), fdst, fsrc, srcFileName, srcFileName)
	})
}

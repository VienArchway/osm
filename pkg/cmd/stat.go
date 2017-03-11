package cmd

import (
	"fmt"

	"github.com/appscode/go-term"
	otx "github.com/appscode/osm/pkg/context"
	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

type itemStatRequest struct {
	context   string
	container string
	itemID    string
}

func NewCmdStat() *cobra.Command {
	req := &itemStatRequest{}
	cmd := &cobra.Command{
		Use:     "stat <id>",
		Short:   "Stat item from container",
		Example: "osm stat -c mybucket f1.txt",
		Run: func(cmd *cobra.Command, args []string) {
			req.itemID = args[0]
			statItem(req)
		},
	}

	cmd.Flags().StringVar(&req.context, "context", "", "Name of osmconfig context to use")
	cmd.Flags().StringVarP(&req.container, "container", "c", "", "Name of container")
	return cmd
}

func statItem(req *itemStatRequest) {
	cfg, err := otx.LoadConfig()
	term.ExitOnError(err)

	loc, err := cfg.Dial(req.context)
	term.ExitOnError(err)

	c, err := loc.Container(req.container)
	term.ExitOnError(err)

	item, err := c.Item(req.itemID)
	term.ExitOnError(err)

	term.Successln("Found item " + req.itemID)
	term.Infoln(fmt.Sprintf("URL = %v", item.URL().String()))
	if sz, err := item.Size(); err == nil {
		term.Infoln(fmt.Sprintf("Size = %v", humanize.Bytes(uint64(sz))))
	}
	if mtime, err := item.LastMod(); err == nil {
		term.Infoln(fmt.Sprintf("Last modified = %v", mtime.String()))
	}
	if etag, err := item.ETag(); err == nil {
		term.Infoln(fmt.Sprintf("Etag = %v", etag))
	}
	if md, err := item.Metadata(); err == nil {
		for k, v := range md {
			term.Infoln(fmt.Sprintf("%v=%v", k, v))
		}
	}
}

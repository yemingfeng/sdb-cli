package cli

import (
	"context"
	"fmt"
	"github.com/abiosoft/ishell/v2"
	pb "github.com/yemingfeng/sdb-cli/pkg/protobuf"
	"strconv"
)

func RegisterClusterCmd(shell *ishell.Shell) {
	shell.AddCmd(newCInfoCmd())
}

func newCInfoCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name: "cinfo",
		Help: "cinfo",
		Func: func(c *ishell.Context) {
			response, err := client.CInfo(context.Background(), &pb.CInfoRequest{})
			if err != nil {
				c.Println(err.Error())
			} else {
				for i := range response.Nodes {
					c.Println(fmt.Sprintf("%d", response.Nodes[i].Id) + "\t" + response.Nodes[i].Address + "\t" + strconv.FormatBool(response.Nodes[i].Leader))
				}
			}
		},
	}
}

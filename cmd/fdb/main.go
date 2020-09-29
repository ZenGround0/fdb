package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/urfave/cli/v2"
)

var decodeCmd = &cli.Command{
	Name:        "decode",
	Description: "decode raw blockchain data",
	Subcommands: []*cli.Command{
		{
			Name:  "bf",
			Usage: "decode raw bitfield bytes",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "b64"},
			},
			Action: runBitfieldDecode,
		},
		{
			Name:  "int",
			Usage: "decode raw big int bytes",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "b64"},
			},
			Action: runIntDecode,
		},
	},
}

func main() {
	app := &cli.App{
		Name:        "fdb",
		Usage:       "filecoin blockchain debug utilities",
		Description: "filecoin blockchain debug utilities",
		Commands: []*cli.Command{
			decodeCmd,
		},
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	for _, c := range app.Commands {
		sort.Sort(cli.FlagsByName(c.Flags))
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func runBitfieldDecode(c *cli.Context) error {
	str := c.Args().First()
	var b []byte
	var err error
	if c.Bool("b64") { // base 64
		b, err = base64.StdEncoding.DecodeString(str)
		if err != nil {
			return err
		}
	} else { // hex string
		b, err = hex.DecodeString(str)
		if err != nil {
			return err
		}
	}

	bf, err := bitfield.NewFromBytes(b)
	if err != nil {
		return err
	}

	var bitset []uint64
	bf.ForEach(func(u uint64) error {
		bitset = append(bitset, u)
		return nil
	})
	outBs, err := json.Marshal(bitset)
	if err != nil {
		return err
	}
	fmt.Println(string(outBs))

	return nil
}

func runIntDecode(c *cli.Context) error {
	str := c.Args().First()
	var b []byte
	var err error
	if c.Bool("b64") { // base 64
		b, err = base64.StdEncoding.DecodeString(str)
		if err != nil {
			return err
		}
	} else { // hex string
		b, err = hex.DecodeString(str)
		if err != nil {
			return err
		}
	}

	i, err := big.FromBytes(b)
	if err != nil {
		return err
	}

	fmt.Println(i)

	return nil
}

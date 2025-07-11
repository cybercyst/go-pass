package commands

import (
	"github.com/urfave/cli/v3"
)

var Show = &cli.Command{
	Name: "show",
	Usage: `Show existing password and optionally put it on the clipboard.
If put on the clipboard, it will be cleared in 45 seconds.
`,
	Description: `
Decrypt  and  print  a password named pass-name. If --clip or -c is specified, do not print the password but instead copy the first (or otherwise specified) line to the clipboard using xclip(1) or wl-
clipboard(1) and then restore the clipboard after 45 (or PASSWORD_STORE_CLIP_TIME) seconds. If --qrcode or -q is specified, do not print the password but instead display a QR  code  using  qrencode(1)
either to the terminal or graphically if supported.
`,
	Action: stubAction,
}

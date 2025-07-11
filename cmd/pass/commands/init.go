package commands

import (
	"github.com/urfave/cli/v3"
)

var Init = &cli.Command{
	Name:  "init",
	Usage: "Initialize a new password storage and use gpg-id for encryption. Selectively reencrypt existing passwords using a new gpg-id.",
	Description: `Initialize new password storage and use gpg-id for encryption. Multiple gpg-ids may be specified, in order to encrypt each password with multiple ids. This command must be run first before a  password
store  can  be  used.  If the specified gpg-id is different from the key used in any existing files, these files will be reencrypted to use the new id.  Note that use of gpg-agent(1) is recommended so
that the batch decryption does not require as much user intervention. If --path or -p is specified, along with an argument, a specific gpg-id or set of gpg-ids is assigned for that specific sub folder
of the password store. If only one gpg-id is given, and it is an empty string, then the current .gpg-id file for the specified sub-folder (or root if unspecified) is removed.
`,
	Action: stubAction,
}

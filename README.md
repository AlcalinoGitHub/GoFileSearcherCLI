# GoFileSearcherCLI

This is a CLI tool that uses a search algorithm to find a file in a certain folder written in golang  
The code is currently singlethreaded, but it is **blazingly fast 🔥** (hehe) thanks to Go performance

The CLI can be launched with several flags

+ `-p` : it assigns the path to search. Defaults to `"."`

+ `-s` : it assigns the name to search (if not provided, everything will be returned). Defaults to `""`

+ `-r` : it assigns the recursive attribute. If true, all subfolders will be searched recursively as well. Defaults to `false`

+ `-m` : it assigns the sorting mode. name, date or size. Defaults to `none`

The CLI will output the files and directories found to the terminal in a JSON format, or will panic if any errors occur

*Note: Running as **SUDO** is recommended when searching on the `~` or `/` directory.*

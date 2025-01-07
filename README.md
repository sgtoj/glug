# glug

`glug` is a cli tool to download binaries from github (and potentially other
sources) in a quick, scriptable, and easily repeatable manner. It was developed
as a bored side project to explore using lua from go. It may be useful to
others, but the future of this project is not guaranteed.

## features

- fetches binaries from github releases (other hosts planned)
- simple configuration of tool registry via lua
- quick usage: `glug get [tool_name]`

## usage

1. build or download `glug`
   ```bash
   go install github.com/sgtoj/glug@latest
   ```
2. run `glug get` for a tool specified in the registry
   ```bash
   glug get fzf
   ```
3. binary will be placed in the configured directory (`./tmp/bin` by default,
   or per your own config)

## how it works

1. parses a lua registry file which declares each tool’s owner, repo, version
   strategy, and url template
2. resolves the latest version from the corresponding release page
3. downloads the release artifact
4. optionally unpacks it if needed (zip or tar.gz)
5. moves the final binary to the specified output location

## project status

This started off as a bored side project to scratch a personal itch. It may or
may not receive ongoing support or active development. Issues and pull requests
are welcome, but please understand that responses may be slow.

## contributing

Contributions, suggestions, or bug reports are always welcome! Feel free to open
an issue or pull request if you’d like to help. Note that major changes may
require some discussion beforehand.


# envset

A small Linux utility to set environment variables across different shells with a single command. envset can apply variables to the current session and also persist them permanently across Bash, Zsh, Fish and other common shells.

Why use envset?
- Single command to add, update, remove and list environment variables.
- Ensures variables are added to the correct shell configuration files so they persist across sessions.
- Minimal, predictable behavior suitable for scripts and interactive use.

## Features
- Set environment variables for the current shell session.
- Persist environment variables across shells (Bash, Zsh, Fish, and POSIX shells).
- Remove persistent environment variables cleanly.
- List configured environment variables managed by envset.
- Safe: backs up modified files before writing.

## Installation

Option A — Install from releases (recommended)
1. Download the latest release binary for your architecture from the project's Releases page.
2. Make it executable and move it to a directory on your PATH, for example:
```bash
chmod +x envset-linux-amd64
sudo mv envset-linux-amd64 /usr/local/bin/envset
```

Option B — Install via curl (quick installer)
Note: Review any install script before executing it.
```bash
curl -sSL https://example.com/envset/install.sh | sudo bash
```

Option C — Build from source
```bash
git clone https://github.com/your-org/envset.git
cd envset
make build        # or the build command the project uses
sudo cp ./envset /usr/local/bin/
```

After installation, verify:
```bash
envset --version
```

## Usage

Important: The exact flags and behavior depend on the installed version. The examples below show common patterns envset supports.

Basic set (current session only)
```bash
envset set MY_VAR "some value"
# or
export MY_VAR="some value" && envset sync
```

Set and persist across shells
```bash
# Persist MY_VAR so it is available in future sessions for supported shells
envset set --global MY_VAR "some value"

# or explicitly target persistence
envset set --persist MY_VAR "some value"
```

Get the value of a variable
```bash
envset get MY_VAR
# prints "some value" or a message when not set
```

Remove a variable (from persisted config and/or current session)
```bash
# Remove from persisted config
envset remove --global MY_VAR

# Remove from current session only
envset remove MY_VAR
```

List variables managed by envset
```bash
envset list
```

Examples that show how envset persists variables:
- For Bash/Zsh it may append or update lines in ~/.bashrc, ~/.bash_profile or ~/.zshrc:
  export MY_VAR="some value"
- For Fish it may write appropriate set commands to ~/.config/fish/config.fish:
  set -x MY_VAR "some value"

envset attempts to detect your shell and update the appropriate file(s). It also keeps a backup of files it modifies (e.g. ~/.bashrc.envset.bak-YYYYMMDD-HHMMSS).

## Notes and best practices
- Always inspect changes: envset will create backups but it's good practice to open and verify the shell config files after persistent changes.
- When using system-managed shells (e.g., in some Linux distributions where profile files are maintained by distribution packaging), prefer adding user-level variables or consult system admin guidelines.
- For scripts running in automated environments, prefer setting variables in the process environment or passing them explicitly instead of relying on persistent changes.

## Uninstall
To remove persistent entries created by envset, run the appropriate remove commands:
```bash
envset remove --global MY_VAR
```
To remove the binary:
```bash
sudo rm /usr/local/bin/envset
```

## Contributing
Contributions, bug reports and feature requests are welcome. Please open issues or pull requests on the repository. If you submit a change that modifies how shell files are edited, include tests and manual steps you used to validate correctness.

## License

MIT License

Copyright (c) 2025 nandisomnath

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

---

If you want, I can:
- add a LICENSE file to the repository with the same MIT text,
- tailor the README to match the exact CLI flags and behavior of your implementation (if you paste the current help output of envset, I will update the Usage section to match),
- generate an example install script or package manifest for common distros.

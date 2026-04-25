# ddg-cli

Get DuckDuckGo search results quickly from your terminal.

`ddg-cli` is a lightweight DuckDuckGo CLI client written in Go. It uses the DuckDuckGo Lite endpoint and keeps dependencies minimal.

## Features

- Search DuckDuckGo with query syntax similar to the website.
- Limit result count with `-limit`.
- Toggle concise output with `-minimal-output` (`-m`).
- Filter by region with `-region` (`-kl`).
- Control safe search with `-safe-search` (`-kp`).
- Output results as JSON with `-json` (`-j`).
- Keep dependencies minimal (Go standard library + `golang.org/x/net/html`).

## Requirements

- Go 1.26.2 or later

## Installation

### Option A: Download from releases

Download the latest archive for your platform from [releases](https://github.com/kr4phy/ddg-cli/releases/latest), then extract and run it.

### Option B: Install using go

Use `go install` to install the binary:

```bash
go install github.com/kr4phy/ddg-cli@latest
```

### Option C: Build from source

Clone this repository:

```bash
git clone https://github.com/kr4phy/ddg-cli
```

Build the binary:
```bash
go build .
```

Run the built binary:
```bash
./ddg-cli
```

Run directly:
```bash
go run .
```

Install globally (optional):
```bash
go install .
```

## Usage

Basic:

```bash
./ddg-cli <query>
```

With options:

```bash
./ddg-cli [options] <query>
```
Note: Put the query after options. Any remaining tokens are joined into a single query string.

Examples:

```bash
./ddg-cli github
./ddg-cli -limit 5 golang cli
./ddg-cli -m linux terminal search
./ddg-cli -json open source licenses
./ddg-cli -region us-en -safe-search 1 github actions
```

## Options

| Option | Alias | Description | Default |
| --- | --- | --- | --- |
| `-limit` | - | Limit the number of results | `10` |
| `-minimal-output` | `-m` | Show only title and URL | `false` |
| `-region` | `-kl` | Set DuckDuckGo region | `wt-wt` |
| `-safe-search` | `-kp` | Safe search level (`1` on, `-1` moderate, `-2` off) | `-1` |
| `-json` | `-j` | Output search results as JSON | `false` |

Run `./ddg-cli -h` to see the latest built-in help output.

If no query is provided, the program prints usage and exits.
If no results are found, it prints `No results found.`

## Output Format

Default output:
- Numbered title
- URL
- Description (omitted with `-minimal-output`)

JSON output (`-json`) returns an array of objects with:
- `Index`
- `Title`
- `URL`
- `Description`

## Optimizations

Compared with `ddgr`, this Go implementation shows lower user/system CPU time and lower peak memory in this test, while some I/O-related metrics may vary depending on the run.
The table below is a raw comparison measured with `/usr/bin/time -v`.

| Metric | ddgr (v2.2) | ddg-cli | Improvement / Difference |
| --- | --- | --- | --- |
| Command being timed | "../../ddgr-2.2/ddgr --np github" | "./ddg-cli github" | - |
| User time (seconds) | 0.24 | 0.03 | ddg-cli is ~8.0x faster |
| System time (seconds) | 0.04 | 0.02 | ddg-cli uses 50% less system time |
| Percent of CPU this job got | 20% | 5% | ddg-cli uses 75% less CPU |
| Elapsed (wall clock) time (h:mm:ss or m:ss) | 0:01.43 | 0:01.02 | ddg-cli is 0.41s faster |
| Average shared text size (kbytes) | 0 | 0 | Identical |
| Average unshared data size (kbytes) | 0 | 0 | Identical |
| Average stack size (kbytes) | 0 | 0 | Identical |
| Average total size (kbytes) | 0 | 0 | Identical |
| Maximum resident set size (kbytes) | 32524 | 13076 | ddg-cli uses ~60% less RAM |
| Average resident set size (kbytes) | 0 | 0 | Identical |
| Major (requiring I/O) page faults | 39 | 89 | ddgr requires less I/O |
| Minor (reclaiming a frame) page faults | 5418 | 1819 | ddg-cli has ~66% fewer memory reclaims |
| Voluntary context switches | 141 | 730 | ddgr switches less frequently |
| Involuntary context switches | 126 | 40 | ddg-cli has ~68% fewer forced switches |
| Swaps | 0 | 0 | None |
| File system inputs | 10820 | 15064 | ddg-cli reads ~39% more data |
| File system outputs | 0 | 0 | None |
| Socket messages sent | 0 | 0 | None |
| Socket messages received | 0 | 0 | None |
| Signals delivered | 0 | 0 | None |
| Page size (bytes) | 4096 | 4096 | Identical |
| Exit status | 0 | 0 | Both Success |

`ddg-cli` uses the DuckDuckGo Lite endpoint instead of the full HTML endpoint.
The project also keeps dependencies small by using only the standard library plus `golang.org/x/net/html`.

Benchmark note: These are single-run measurements and can vary by network, cache state, machine load, and DuckDuckGo response time.

## License

This project is licensed under the GNU General Public License v3.0.
See `LICENSE` for details.

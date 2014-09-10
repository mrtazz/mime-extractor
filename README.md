# mime-extractor

## Overview
Simple tool to extract the subject of an email on STDIN and print the parsed
MIME body as text. It prints the subject on the first line and body on
subsequent lines to be compatible with
[disrupt](https://github.com/mrtazz/disrupt).

## Usage
On the command line:
```
cat foo.email | mime-extractor
```
or with procmail and disrupt for cheap push notifications:
```
:0
* ^(Cc|To).*foo@example.com
| $HOME/bin/mime-extractor | $HOME/bin/cli-notifier
```

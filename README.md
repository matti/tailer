# tailer

tails multiple files and outputs with prefixes wheater they exist or not.

why: `tail -f` / `tail -F` multi file output is horrible and it doesn't handle all cases when files come and go

## example

    $ tailer test1 test2
      test1: hello
      test2: world
      test1: again
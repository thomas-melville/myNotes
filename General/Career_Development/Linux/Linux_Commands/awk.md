# awk

useful command for extracting specific text from formatted strings

echo "..." | awk '{print $3}'

print the third piece of text per line after each line has been split

## flags

-F <text>   separator to use to split test, default is space
            multiple separators can be specified using |
            this can also be a regular expression
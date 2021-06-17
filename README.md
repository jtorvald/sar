# sar

Simple Search and Replace that supports regex with tabs and newlines from the commandline.

Usage: `cat file.txt | sar "some (text):\n\t(foo)" "other \$1:\n\t\${2}bar"`

Purpose: easy string replace in CI/CD when `sed` lets you down.

Second purpose: play with github workflow.

# LICENSE

MIT

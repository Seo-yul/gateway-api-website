#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."

INDEX="content/en/geps/_index.md"

# Save a copy of the current file.
cp "$INDEX" "$INDEX.bak"

# Regenerate.
hack/update-geps.sh > /dev/null

# Compare.
if ! diff -q "$INDEX" "$INDEX.bak" > /dev/null 2>&1; then
    echo "ERROR: $INDEX is out of date. Run 'make update-geps' to regenerate."
    diff -u "$INDEX.bak" "$INDEX" || true
    mv "$INDEX.bak" "$INDEX"
    exit 1
fi

rm -f "$INDEX.bak"
echo "OK: $INDEX is up to date."

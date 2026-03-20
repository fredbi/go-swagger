#!/usr/bin/env bash
# SPDX-License-Identifier: Apache-2.0
#
# Fetch all open issues from go-swagger/go-swagger into a JSON file.
# Usage: ./fetch-issues.sh [output_file]
#
# Requires: gh (GitHub CLI) authenticated with GH_TOKEN.

set -euo pipefail

REPO="go-swagger/go-swagger"
OUTPUT="${1:-.claude/agents/issues.json}"

echo "Fetching all open issues from ${REPO}..."

gh issue list \
  --repo "${REPO}" \
  --state open \
  --limit 9999 \
  --json number,title,labels,createdAt,updatedAt,author,url,body,milestone,isPinned \
  | python3 -c "
import json, sys
issues = json.load(sys.stdin)
# Sort by issue number
issues.sort(key=lambda x: x['number'])
json.dump(issues, sys.stdout, indent=2)
" > "${OUTPUT}"

COUNT=$(python3 -c "import json; print(len(json.load(open('${OUTPUT}'))))")
echo "Done. ${COUNT} open issues saved to ${OUTPUT}"

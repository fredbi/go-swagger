#!/usr/bin/env bash
# SPDX-License-Identifier: Apache-2.0
#
# Attach or detach sub-issues to/from a parent issue using the GitHub GraphQL API.
#
# Usage:
#   ./sub-issues.sh attach <parent_issue_number> <child_issue_number> [child_issue_number...]
#   ./sub-issues.sh detach <parent_issue_number> <child_issue_number> [child_issue_number...]
#   ./sub-issues.sh list   <parent_issue_number>
#
# Requires: gh (GitHub CLI) authenticated with GH_TOKEN.
#
# NOTE: GitHub sub-issues API constraints:
#   - A parent issue may have at most 100 direct children
#   - Up to 8 nested levels

set -euo pipefail

REPO="go-swagger/go-swagger"
OWNER="go-swagger"
REPO_NAME="go-swagger"

usage() {
  echo "Usage:"
  echo "  $0 attach <parent_issue#> <child_issue#> [child_issue#...]"
  echo "  $0 detach <parent_issue#> <child_issue#> [child_issue#...]"
  echo "  $0 list   <parent_issue#>"
  exit 1
}

# Resolve an issue number to its GraphQL node ID
issue_node_id() {
  local issue_number="$1"
  gh api graphql -f query='
    query($owner: String!, $repo: String!, $number: Int!) {
      repository(owner: $owner, name: $repo) {
        issue(number: $number) {
          id
        }
      }
    }
  ' -F owner="${OWNER}" -F repo="${REPO_NAME}" -F number="${issue_number}" \
    --jq '.data.repository.issue.id'
}

attach_sub_issue() {
  local parent_id="$1"
  local child_id="$2"
  gh api graphql -f query='
    mutation($parentId: ID!, $childId: ID!) {
      addSubIssue(input: {issueId: $parentId, subIssueId: $childId}) {
        issue {
          number
          title
        }
        subIssue {
          number
          title
        }
      }
    }
  ' -F parentId="${parent_id}" -F childId="${child_id}" --jq '.data.addSubIssue'
}

detach_sub_issue() {
  local parent_id="$1"
  local child_id="$2"
  gh api graphql -f query='
    mutation($parentId: ID!, $childId: ID!) {
      removeSubIssue(input: {issueId: $parentId, subIssueId: $childId}) {
        issue {
          number
          title
        }
        subIssue {
          number
          title
        }
      }
    }
  ' -F parentId="${parent_id}" -F childId="${child_id}" --jq '.data.removeSubIssue'
}

list_sub_issues() {
  local parent_id="$1"
  gh api graphql -f query='
    query($parentId: ID!) {
      node(id: $parentId) {
        ... on Issue {
          number
          title
          subIssues(first: 100) {
            nodes {
              number
              title
              state
            }
            totalCount
          }
        }
      }
    }
  ' -F parentId="${parent_id}" --jq '.data.node'
}

# --- main ---

[[ $# -lt 2 ]] && usage

ACTION="$1"
PARENT_NUMBER="$2"
shift 2

echo "Resolving parent issue #${PARENT_NUMBER}..."
PARENT_ID=$(issue_node_id "${PARENT_NUMBER}")
echo "  Parent node ID: ${PARENT_ID}"

case "${ACTION}" in
  attach)
    [[ $# -lt 1 ]] && usage
    for CHILD_NUMBER in "$@"; do
      echo "Resolving child issue #${CHILD_NUMBER}..."
      CHILD_ID=$(issue_node_id "${CHILD_NUMBER}")
      echo "  Attaching #${CHILD_NUMBER} as sub-issue of #${PARENT_NUMBER}..."
      attach_sub_issue "${PARENT_ID}" "${CHILD_ID}"
      echo "  Done."
    done
    ;;
  detach)
    [[ $# -lt 1 ]] && usage
    for CHILD_NUMBER in "$@"; do
      echo "Resolving child issue #${CHILD_NUMBER}..."
      CHILD_ID=$(issue_node_id "${CHILD_NUMBER}")
      echo "  Detaching #${CHILD_NUMBER} from #${PARENT_NUMBER}..."
      detach_sub_issue "${PARENT_ID}" "${CHILD_ID}"
      echo "  Done."
    done
    ;;
  list)
    list_sub_issues "${PARENT_ID}"
    ;;
  *)
    usage
    ;;
esac

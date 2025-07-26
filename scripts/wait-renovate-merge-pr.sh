#!/bin/bash

for i in {1..60}; do
    PR_ID=$(gh pr list \
      --repo "https://github.com/$REPO" \
      --label "env::dev" \
      --json number,title \
      | jq -e --arg tag "$TARGET_VERSION" '.[] | select(.title | test("chore\\(deps\\): update.*" + $tag))| .number')

    if [ -n "$PR_ID" ]; then
      echo "No merged PR yet. Retrying in 5 seconds..."
      sleep 5
    else
      echo "Found merged Renovate PR"
      exit 0
    fi
done

echo "Timeout: PR was not merged after waiting."
exit 1

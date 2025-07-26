#!/bin/bash

WORKFLOW_NAME="Run Renovate on Dispatch"
ARTIFACT_NAME="version-${TARGET_VERSION}"

echo "Looking for a workflow run in $REPO with artifact: $ARTIFACT_NAME"

for i in {1..60}; do
  RUNS=$(gh run list -R "$REPO" --workflow "$WORKFLOW_NAME" --limit 10 --json databaseId,status,conclusion,name,headSha)

  while read -r RUN_ID; do
    ARTIFACTS=$(gh api "/repos/$REPO/actions/runs/$RUN_ID/artifacts" --jq '.artifacts[].name')

    for ARTIFACT in $ARTIFACTS; do
      if [[ "$ARTIFACT" == "$ARTIFACT_NAME" ]]; then
        # Get run status
        STATUS=$(gh api "/repos/$REPO/actions/runs/$RUN_ID" --jq '.status')
        CONCLUSION=$(gh api "/repos/$REPO/actions/runs/$RUN_ID" --jq '.conclusion')

        echo "Found run $RUN_ID with matching artifact. Status: $STATUS"

        if [[ "$STATUS" == "completed" && "$CONCLUSION" == "success" ]]; then
          echo "Workflow run $RUN_ID completed successfully."
          exit 0
        elif [[ "$STATUS" == "completed" ]]; then
          echo "Workflow $RUN_ID completed but failed."
          exit 1
        else
          echo "Run $RUN_ID still in progress..."
        fi
      fi
    done
  done <<< "$(echo "$RUNS" | jq -r '.[].databaseId')"

  echo "Waiting 10s for the workflow to finish..."
  sleep 10
done

echo "Timeout: Workflow with artifact $ARTIFACT_NAME did not finish in time"
exit 1

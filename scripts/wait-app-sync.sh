#!/bin/bash

echo "Waiting for vote-app-dev to be Synced and Healthy..."
    for i in {1..60}; do
      STATUS=$(argocd app get vote-app-dev \
        --server "$ARGOCD_SERVER" \
        --auth-token "$ARGOCD_TOKEN" \
        --grpc-web \
        -o json | jq -r '.status.sync.status + "-" + .status.health.status')

      echo "Current status: $STATUS"

      if [[ "$STATUS" == "Synced-Healthy" ]]; then
        echo "App is synced and healthy!"
        exit 0
      fi

      echo "App not yet synced/healthy. Retrying in 10 seconds..."
      sleep 10
done

echo "Timeout: App did not reach Synced/Healthy state in time."
exit 1

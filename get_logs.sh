kubectl logs -n spinnaker $(kubectl get pod -n spinnaker | grep spinnaker-audit | awk '{print $1}') | jq .

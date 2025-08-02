### Spinnaker AuditLogs Sink

```
make all
make dbuild
kind load docker-image spinaudit:latest --name spinnaker
kubectl apply -f deployment.yaml
```

# k8s-lab

A simple Go REST API and Kubernetes examples for learning.

## Run API

```bash
make run
```

The API starts on:

```
localhost:8080
```

## Endpoints

Health check:

```bash
curl localhost:8080/health
```

Response:

```json
{
  "status": "ok"
}
```

Message:

```bash
curl localhost:8080/api/v1/message
```

Response:

```json
{
  "message": "hello from go api"
}
```

## Development

Format code:

```bash
make fmt
```

Run tests:

```bash
make test
```

Run lint:

```bash
make lint
```

Build:

```bash
make build
```

## Kubernetes Examples

The `basics/` folder contains example Kubernetes manifests.

Apply:

```bash
kubectl apply -f basics/
```

Delete:

```bash
kubectl delete -f basics/
```
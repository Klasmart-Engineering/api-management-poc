# KrakenD Instructions

KrakenD is deployed as a standalone pod in kubernetes. It currently has everything from the host domain of
`krakend.kong-poc.kidsloop.live` routing directly to it.

## Authenticating to AWS ECR

```sh
aws ecr get-login-password --region eu-west-2 | docker login --username AWS --password-stdin <ECR URL>
```

## Deploy a new version/change config

1. Update `krakend.json` file with latest config
2. `docker build -t kidsloop/krakend:latest .`
3. `docker tag kidsloop/krakend 246958245973.dkr.ecr.eu-west-2.amazonaws.com/krakend:<image verison number>`
4. `docker push 246958245973.dkr.ecr.eu-west-2.amazonaws.com/krakend:<image version number>`
5. Edit `krakend/deployment.yaml` to point at the latest version tag
6. `kubectl apply -f deployment.yaml`

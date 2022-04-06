VERSION='0.0.11'
ECR_URL='246958245973.dkr.ecr.eu-west-2.amazonaws.com/krakend'
VERSION="$(echo $VERSION | awk -F. '{$NF = $NF + 1;} 1' | sed 's/ /./g')"
sed -i -e "1s/.*/VERSION='$VERSION'/" bump.sh
echo "Next Version: $VERSION"
rm bump.sh-e
docker build -t kidsloop/krakend:latest --platform=linux/amd64 .
docker tag kidsloop/krakend 246958245973.dkr.ecr.eu-west-2.amazonaws.com/krakend:$VERSION
docker push 246958245973.dkr.ecr.eu-west-2.amazonaws.com/krakend:$VERSION
sed -i -e "s,image.*$,\image: $ECR_URL:$VERSION," deployment.yaml
rm deployment.yaml-e
kubectl apply -f deployment.yaml

VERSION='0.0.25'
ECR_URL='246958245973.dkr.ecr.eu-west-2.amazonaws.com/krakend-enterprise'
VERSION="$(echo $VERSION | awk -F. '{$NF = $NF + 1;} 1' | sed 's/ /./g')"
sed -i -e "1s/.*/VERSION='$VERSION'/" bump.sh
echo "Next Version: $VERSION"
rm bump.sh-e
docker build -t kidsloop/krakend-e:latest --platform=linux/amd64 --build-arg ENV=dev .
docker tag kidsloop/krakend-e:latest $ECR_URL:$VERSION
docker push $ECR_URL:$VERSION
sed -i -e "s,image.*$,\image: $ECR_URL:$VERSION," deployment.yaml
rm deployment.yaml-e
kubectl apply -f deployment.yaml

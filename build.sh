echo "==========>> Start Fetch deps <<=========="
GO111MODULE=on go mod vendor 
echo "==========>> Start Build <<=========="
GO111MODULE=on go build -mod=vendor
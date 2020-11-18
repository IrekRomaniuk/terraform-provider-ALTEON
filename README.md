# terraform-provider-alteon


```
go mod terraform-provider-alteon
go clean --modcache (optional)
go mod vendor 
```

```
make install
```
or

```
go build -o terraform-provider-alteon
mkdir -p ~/.terraform.d/plugins/github.com/irekromaniuk/alteon/0.1/linux_amd64
mv terraform-provider-alteon ~/.terraform.d/plugins/github.com/irekromaniuk/alteon/0.1/linux_amd64
cd examples
terraform init && terraform apply --auto-approve
```


export ALTEON_USERNAME=
export ALTEON_PASSWORD=
export ALTEON_URI=

echo $ALTEON_USERNAME
echo $ALTEON_PASSWORD
echo $ALTEON_URI
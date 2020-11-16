# terraform-provider-alteon


```
go mod init github.com/irekromaniuk/terraform-provider-alteon
go clean --modcache (optional)
go mod vendor 
go build -o terraform-provider-alteon
mkdir -p ~/.terraform.d/plugins/github.com/irekromaniuk/alteon/0.1/linux_amd64
mv terraform-provider-alteon ~/.terraform.d/plugins/github.com/irekromaniuk/alteon/0.1/linux_amd64
cd examples
terraform init && terraform apply --auto-approve
```


export ALTEON_USERNAME=manager
export ALTEON_PASSWORD=
export ALTEON_URI="https://10.96.1.51:443/config"

echo $ALTEON_USERNAME
echo $ALTEON_PASSWORD
echo $ALTEON_URI
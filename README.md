# terraform-provider-alteon


```
go mod init terraform-provider-alteon
go mod vendor
go build -o terraform-provider-alteon
mkdir -p ~/.terraform.d/plugins/github.com/irekromaniuk/alteon/0.1/linux_amd64
mv terraform-provider-alteon ~/.terraform.d/plugins/github.com/irekromaniuk/alteon/0.1/linux_amd64
cd examples
terraform init && terraform apply --auto-approve
```


export ALTEON_USERNAME=
export ALTEON_USERNAME=
ALTEON_URI=10.96.1.51
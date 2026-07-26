---
title: Install from distribution package
weight: 20
description: Installation instructions for go-swagger, from a distribution package
---

`go-swagger` packages for linux distributions are shipped courtesy of the [CloudSmith](https://broadcasts.cloudsmith.com/go-swagger/go-swagger) service.

{{% notice style="warning" %}}
At this moment, we do not support pre-packaged binaries for Windows and Alpine Linux.
{{% /notice %}}

## Installing from distribution package

{{< tabs groupid="distroInstallType" >}}
{{% tab title="Linux deb (Ubuntu, Debian)" %}}

### Debian packages ![debian logo](images/debian.png)

[![Download](https://api-prd.cloudsmith.io/v1/badges/version/go-swagger/go-swagger/deb/go-swagger/latest/a=amd64;d=debian%252Fany-version;t=binary/?render=true&show_latest=true)](https://cloudsmith.io/~go-swagger/repos/go-swagger/packages/detail/deb/go-swagger/latest/a=amd64;d=debian%252Fany-version;t=binary/)

This package will work for any Debian. The only file it contains gets copied to `/usr/bin`

* Update and install prerequisite packages
```cmd
sudo apt update
sudo apt install -y apt-transport-https gnupg curl debian-keyring debian-archive-keyring
```

* Register our GPG signing key and the repository source
```cmd
curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/gpg.2F8CB673971B5C9E.key' | sudo gpg --dearmor -o /usr/share/keyrings/go-swagger-go-swagger-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/config.deb.txt?distro=debian&codename=any-version' | sudo tee /etc/apt/sources.list.d/go-swagger-go-swagger.list
```

* Install

```cmd
sudo apt update 
sudo apt install go-swagger
```

{{% notice style="warning" %}}
If you already had package `swagger` installed from a previous release, please remove it first (the package name has changed from `swagger` to `go-swagger`):

```cmd
sudo apt remove swagger
```
{{% /notice %}}

{{% /tab %}}

{{% tab title="Linux rpm (CentOS, Fedora, Suse)" %}}
### RPM packages ![fedora logo](images/fedora.png)
[![Download](https://api-prd.cloudsmith.io/v1/badges/version/go-swagger/go-swagger/rpm/go-swagger/latest/a=x86_64;d=fedora%252Fany-version;t=binary/?render=true&show_latest=true)](https://cloudsmith.io/~go-swagger/repos/go-swagger/packages/detail/rpm/go-swagger/latest/a=x86_64;d=fedora%252Fany-version;t=binary/)

This package should work on any distro that wants RPM packages. The only file it contains gets copied to `/usr/bin`

* Update and install prerequisite packages

```cmd
dnf install -y yum-utils
```

* Register our GPG signing key

```cmd
rpm --import 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/gpg.2F8CB673971B5C9E.key'
```

* Install

```cmd
curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/config.rpm.txt?distro=fedora&codename=any-version' > /tmp/go-swagger-go-swagger.repo
dnf config-manager --add-repo '/tmp/go-swagger-go-swagger.repo'
dnf -q makecache -y --disablerepo='*' --enablerepo='go-swagger-go-swagger' --enablerepo='go-swagger-go-swagger-source'
dnf install -y go-swagger
```

{{% notice style="warning" %}}
If you already had package `swagger` installed from a previous release, please remove it first (the package name has changed from `swagger` to `go-swagger`):

```cmd
dnf remove swagger
```
{{% /notice %}}
{{% /tab %}}
{{< /tabs >}}

<!-- TODO apk package for alpine -->
<!-- TODO msi package for windows -->

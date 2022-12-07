
##### 一、项目开发

![](https://www.showdoc.com.cn/server/api/attachment/visitFile?sign=b636173848377beec121930913d21b74&file=file.png)

- ` 1.git fork `

```c
打开 https://github.com/kubesphere/kubesphere
点击Fork
```
![](https://www.showdoc.com.cn/server/api/attachment/visitFile?sign=4fbcfdbc19e8b714b336b0b9f86b14bb&file=file.png)

- ` 2.编辑项目名称为kubesphere_dev`
- 点击Create Fork

![](https://www.showdoc.com.cn/server/api/attachment/visitFile?sign=6d97c45a55ec3298707fea498cf40fcb&file=file.png)

- `3.安装golang`

```c
yum install golang-1.18 -y
[root@master01 ~]# go env
GOMODCACHE="/root/go/pkg/mod"
GOPATH="/root/go"
GOROOT="/usr/lib/golang-1.18"
GOVERSION="go1.18.8"

go env -w GOPROXY=https://goproxy.io,direct
go env -w GO111MODULE=on
go env -w GOSUMDB=off

export https_proxy=http://10.130.0.16:7890

[root@master01 ~]# mkdir -p /root/go/src/kubesphere.io
```
- `4.克隆kubesphere_dev到本地` 

```c
[root@master01 ~]# cd /root/go/src/kubesphere.io
[root@master01 kubesphere.io]# git clone https://ghproxy.com/https://github.com/licloud/kubesphere_dev.git
[root@master01 kubesphere.io]# cd kubesphere_dev/
[root@master01 kubesphere_dev]# git remote add upstream https://ghproxy.com/https://github.com/kubesphere/kubesphere.git
[root@master01 kubesphere_dev]# git remote -v
origin	https://ghproxy.com/https://github.com/licloud/kubesphere_dev.git (fetch)
origin	https://ghproxy.com/https://github.com/licloud/kubesphere_dev.git (push)
upstream	https://ghproxy.com/https://github.com/kubesphere/kubesphere.git (fetch)
upstream	https://ghproxy.com/https://github.com/kubesphere/kubesphere.git (push)

[root@master01 kubesphere_dev]# git remote set-url --push upstream no_push
[root@master01 kubesphere_dev]# git remote -v
origin	https://ghproxy.com/https://github.com/licloud/kubesphere_dev.git (fetch)
origin	https://ghproxy.com/https://github.com/licloud/kubesphere_dev.git (push)
upstream	https://ghproxy.com/https://github.com/kubesphere/kubesphere.git (fetch)
upstream	no_push (push)

```
- `5.同步原始源码的master`

```c
[root@master01 kubesphere_dev]# git fetch upstream
来自 https://ghproxy.com/https://github.com/kubesphere/kubesphere
 * [新分支]              advanced-1.0                                                                -> upstream/advanced-1.0
 * [新分支]              advanced-2.0                                                                -> upstream/advanced-2.0
 * [新分支]              dependabot/go_modules/staging/src/kubesphere.io/utils/helm.sh/helm/v3-3.9.4 -> upstream/dependabot/go_modules/staging/src/kubesphere.io/utils/helm.sh/helm/v3-3.9.4
 * [新分支]              express                                                                     -> upstream/express
 * [新分支]              kse-3.3-fix                                                                 -> upstream/kse-3.3-fix
 * [新分支]              master                                                                      -> upstream/master
 * [新分支]              release-2.1                                                                 -> upstream/release-2.1
 * [新分支]              release-3.0                                                                 -> upstream/release-3.0
 * [新分支]              release-3.1                                                                 -> upstream/release-3.1
 * [新分支]              release-3.2                                                                 -> upstream/release-3.2
 * [新分支]              release-3.3                                                                 -> upstream/release-3.3

[root@master01 kubesphere_dev]# git checkout master
已经位于 'master'
您的分支与上游分支 'origin/master' 一致。

[root@master01 kubesphere_dev]# git rebase upstream/master
当前分支 master 是最新的。
```
- `6.创建新的分支并提交`

```c
[root@master01 kubesphere_dev]# git checkout -b loong64-v3.3.0 v3.3.0
[root@master01 kubesphere_dev]# ls
api  build  cmd  config  CONTRIBUTING.md  docs  go.mod  go.sum  hack  install  kube  LICENSE  Makefile  OWNERS  pkg  README.md  README_zh.md  SECURITY.md  staging  test  tools  vendor
```
- `创建token `

![](https://www.showdoc.com.cn/server/api/attachment/visitFile?sign=eae0a70c2b4cafdf643a351e41ff965b&file=file.png)
- 点击Developer settings
![](https://www.showdoc.com.cn/server/api/attachment/visitFile?sign=e3a08b573a32416b3a5edd6b1d0ec471&file=file.png)
![](https://www.showdoc.com.cn/server/api/attachment/visitFile?sign=6ae6e7da59dfe598e160cda198dcd647&file=file.png)

- Repository access选项选择All repositories并给权限

```c
[root@master01 kubesphere_dev]# git push --set-upstream origin loong64-v3.3.0 -vv
推送到 https://ghproxy.com/https://github.com/licloud/kubesphere_dev.git
Username for 'https://ghproxy.com': ligang_ll@163.com
Password for 'https://ligang_ll@163.com@ghproxy.com': 
总共 0（差异 0），复用 0（差异 0），包复用 0
POST git-receive-pack (193 bytes)
remote: 
remote: Create a pull request for 'loong64-v3.3.0' on GitHub by visiting:
remote:      https://github.com/licloud/kubesphere_dev/pull/new/loong64-v3.3.0
remote: 
To https://ghproxy.com/https://github.com/licloud/kubesphere_dev.git
 * [new branch]          loong64-v3.3.0 -> loong64-v3.3.0
分支 'loong64-v3.3.0' 设置为跟踪来自 'origin' 的远程分支 'loong64-v3.3.0'。
updating local tracking ref 'refs/remotes/origin/loong64-v3.3.0'

token : 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
```

- `7.编译 `

```c
编译报错：vendor/golang.org/x/sys/unix/affinity_linux.go:13:20: undefined: _CPU_SETSIZE

[root@master01 kubesphere_dev]# cd vendor/golang.org/x/
[root@master01 x]# ls
crypto  lint  net  oauth2  sync  sys  term  text  time  tools  xerrors
替换龙芯sys
[root@master01 x]# rm -rf sys
[root@master01 x]# ls
crypto  lint  net  oauth2  sync  term  text  time  tools  xerrors
[root@master01 x]# cp  -r /usr/lib/golang-1.18/src/cmd/vendor/golang.org/x/sys/ .
[root@master01 x]# ls
crypto  lint  net  oauth2  sync  sys  term  text  time  tools  xerrors

编译报错：
vendor/github.com/prometheus/procfs/cpuinfo.go:71:9: undefined: parseCPUInfo
增加龙芯代码支持：
[root@master01 kubesphere_dev]# cd vendor/github.com/prometheus/procfs/
[root@master01 procfs]# vim cpuinfo_loong64.go
// +build linux
  
package procfs

var parseCPUInfo = parseCPUInfoLoong64

[root@master01 procfs]# vim cpuinfo.go
	return cpuinfo, nil
}
+409行------------------------------------------------------
func parseCPUInfoLoong64(info []byte) ([]CPUInfo, error) {
	scanner := bufio.NewScanner(bytes.NewReader(info))

	// find the first "processor" line
	firstLine := firstNonEmptyLine(scanner)
	if !strings.HasPrefix(firstLine, "system type") || !strings.Contains(firstLine, ":") {
		return nil, errors.New("invalid cpuinfo file: " + firstLine)
	}
	field := strings.SplitN(firstLine, ": ", 2)
	cpuinfo := []CPUInfo{}
	systemType := field[1]

	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, ":") {
			continue
		}
		field := strings.SplitN(line, ": ", 2)
		switch strings.TrimSpace(field[0]) {
		case "processor":
			v, err := strconv.ParseUint(field[1], 0, 32)
			if err != nil {
				return nil, err
			}
			i = int(v)
			cpuinfo = append(cpuinfo, CPUInfo{}) // start of the next processor
			cpuinfo[i].Processor = uint(v)
			cpuinfo[i].VendorID = systemType
		case "cpu model":
			cpuinfo[i].ModelName = field[1]
		case "BogoMIPS":
			v, err := strconv.ParseFloat(field[1], 64)
			if err != nil {
				return nil, err
			}
			cpuinfo[i].BogoMips = v
		}
	}
	return cpuinfo, nil
}

// firstNonEmptyLine advances the scanner to the first non-empty line
// and returns the contents of that line
func firstNonEmptyLine(scanner *bufio.Scanner) string {

编译报错：gzip: stdin: not in gzip format
需要龙芯对应的kubebuilder，暂时跳过test编译。https://github.com/kubernetes-sigs/kubebuilder

编译报错：Unsupported host arch
[root@master01 kubesphere_dev]# make ks-apiserver
...Begin to build ks-apiserver binary.
hack/gobuild.sh cmd/ks-apiserver;
!!! [1207 14:07:37] Unsupported host arch. Must be x86_64, 386, arm, arm64, s390x or ppc64le.
Cmake: *** [Makefile:59: ks-apiserver] Interrupt

修改指定mod的依赖管理——本地vendor:
[root@master01 kubesphere_dev]# vim hack/gobuild.sh 
GOOS=${BUILD_GOOS} CGO_ENABLED=0 GOARCH=${BUILD_GOARCH} ${GOBINARY} build -mod=vendor\
增加龙芯架构支持：
[root@master01 kubesphere_dev]# vim hack/lib/util.sh +158
    loongarch64*)
      host_arch=loong64
      ;;
    *)
      kube::log::error "Unsupported host arch. Must be x86_64, 386, arm, arm64, s390x, loong64 or ppc64le."

```
```
[root@master01 kubesphere_dev]# make ks-apiserver
...Begin to build ks-apiserver binary.
hack/gobuild.sh cmd/ks-apiserver;

[root@master01 kubesphere_dev]# cd bin/cmd
[root@master01 cmd]# ./ks-apiserver -h

[root@master01 kubesphere_dev]# make ks-controller-manager
...Begin to build ks-controller-manager binary.
hack/gobuild.sh cmd/controller-manager

[root@master01 kubesphere_dev]# cd bin/cmd/
[root@master01 cmd]# ls
controller-manager  ks-apiserver

[root@master01 cmd]# ./controller-manager -h
```

- `8.`

```c

```

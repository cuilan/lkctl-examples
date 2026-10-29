# lkctl-examples

[license-key-verify](https://github.com/cuilan/license-key-verify) 的官方示例仓库，演示应用集成许可证校验的两种方式。

## 目录结构

```
lkctl-examples/
├── autocheck/          # 示例 A：import 即校验（autocheck）
│   ├── main.go
│   ├── licensekey/     # lkctl embed 生成的公钥源文件
│   └── license.lic
├── sdk/                # 示例 B：显式一行调用（SDK）
│   ├── main.go
│   ├── licensekey/
│   └── license.lic
└── keys/               # 演示用密钥对（仅本仓库测试，请勿用于生产）
```

## 两种方式对比

| | 示例 A（autocheck） | 示例 B（SDK） |
|---|---|---|
| 集成代码 | 2 行 import，`main` 零校验代码 | 1 行 import + 1 行 `license.MustCheck()` |
| 校验时机 | `init()` 自动执行 | 由你控制（`main` 开头或任意位置） |
| 校验结果 | 拿不到（失败直接退出） | 完整 `Result`（客户、features、有效期） |
| 适用场景 | fork 了 license-key-verify 仓库，追求最少代码 | 独立模块引用，需要按 feature 开关功能或展示授权信息 |

> **fork 场景提示**：如果你 fork 了 license-key-verify 并用 `lkctl embed` 覆盖其 `pkg/license/key.go`，则只需 `import _ ".../v2/autocheck"` 一行即可（key.go 在 autocheck 的依赖包里，init 顺序天然正确）。本仓库两个示例都是独立模块引用，因此 key.go 放在自己的 `licensekey/` 子包中，且**必须列在 autocheck 之前 import**。

## 快速开始

前提：[lkctl](https://github.com/cuilan/license-key-verify/releases/latest) 已下载、Go 1.23+ 已安装，且本仓库与 license-key-verify 仓库同级存放（示例的 `go.mod` 用 `replace` 指向本地 checkout，便于离线复现；你自己的工程中请删除 `replace`，改用 `go get github.com/cuilan/license-key-verify/v2@v2.0.0`）。

```bash
# 1. 生成密钥对（你本地，私钥绝不外发）
lkctl keygen --out keys/

# 2. 把公钥嵌入两个示例（包名自动取目录名）
lkctl embed --public-key keys/public.pem --out autocheck/licensekey/key.go
lkctl embed --public-key keys/public.pem --out sdk/licensekey/key.go

# 3. 编译两个示例
cd autocheck && go build -o autocheck.exe . && cd ..
cd sdk && go build -o sdk.exe . && cd ..

# 4. 获取本机指纹并分别签发
lkctl fingerprint
lkctl sign --private-key keys/private.pem --fingerprint <指纹> --out autocheck/license.lic
lkctl sign --private-key keys/private.pem --fingerprint <指纹> --out sdk/license.lic

# 5. 运行
./autocheck/autocheck.exe     # license OK, application starts... / 成功启动
./sdk/sdk.exe                 # license OK, application starts... / expires in 364 days
```

放到其他机器上运行会报 `machine fingerprint does not match`——这正是许可证绑定目标机器的预期行为。

## 注意事项

- `keys/`（私钥）与 `*.lic` 不应提交到你的真实项目仓库
- 示例中的密钥对仅用于演示流程；生产环境务必用自己生成的密钥重新 embed 后再编译
- 校验失败时两个示例都以退出码 1 结束，并打印具体原因（找不到文件 / 签名无效 / 已过期 / 指纹不匹配）

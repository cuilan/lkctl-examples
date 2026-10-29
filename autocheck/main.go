// 示例 A —— 方式一：import 即校验（autocheck）。
//
// 集成只需要两步：
//
//  1. 把公钥生成为 Go 源文件，放进应用的独立子包（本示例为 licensekey/）：
//     lkctl embed --public-key keys/public.pem --out licensekey/key.go
//
//  2. 在 main 包 import 该子包（注册公钥）和 autocheck（自动校验）：
//     import _ "autocheck/licensekey"
//     import _ "github.com/cuilan/license-key-verify/autocheck"
//
// main 里不写任何校验代码——autocheck 在 init 阶段自动定位并校验
// license.lic，失败时打印原因并以退出码 1 结束进程。
//
// 注意：key.go 所在子包必须列在 import 中，且在 autocheck 之前注册。
// 若你 fork 了 license-key-verify 仓库，key.go 直接覆盖 pkg/license/key.go，
// 则无需本示例的 licensekey 子包，只 import autocheck 一行即可。
//
// 复现签发（keys/ 与 license.lic 由 lkctl 生成，见仓库根 README）：
//
//	lkctl fingerprint
//	lkctl sign --private-key keys/private.pem --fingerprint <指纹> --out license.lic
package main

import (
	"fmt"

	_ "autocheck/licensekey" // 注册内嵌公钥（必须先于 autocheck）

	_ "github.com/cuilan/license-key-verify/v2/autocheck" // init 自动校验
)

func main() {
	// 运行到这里说明 autocheck 已校验通过
	fmt.Println("license OK, application starts...")
	fmt.Println("成功启动")
}

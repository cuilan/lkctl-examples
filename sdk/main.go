// 示例 B —— 方式二：显式一行调用（SDK）。
//
// 集成只需要两步：
//
//  1. 把公钥生成为 Go 源文件，放进应用的独立子包（本示例为 licensekey/）：
//     lkctl embed --public-key keys/public.pem --out licensekey/key.go
//
//  2. 在 main 里显式调用校验 API：
//     import _ "sdk/licensekey"                // 注册内嵌公钥
//     import "github.com/cuilan/license-key-verify/pkg/license"
//     license.MustCheck()                      // 失败打印原因并退出码 1
//
// 相比 autocheck（示例 A），SDK 方式由你控制校验时机，并能拿到完整的
// 校验结果（客户名、features、有效期等），适合按 feature 开关功能、
// 在 UI/日志中展示授权信息等场景。
//
// 复现签发（keys/ 与 license.lic 由 lkctl 生成，见仓库根 README）：
//
//	lkctl fingerprint
//	lkctl sign --private-key keys/private.pem --fingerprint <指纹> --out license.lic
package main

import (
	"fmt"

	_ "sdk/licensekey" // 注册内嵌公钥

	"github.com/cuilan/license-key-verify/v2/pkg/license"
)

func main() {
	// 一行集成：校验失败打印原因并以退出码 1 结束
	license.MustCheck()

	// 运行到这里说明许可证有效
	fmt.Println("license OK, application starts...")

	// 需要更细的控制时（例如按 feature 开关功能）：
	result := license.VerifyFile(license.DefaultLicensePath())

	if len(result.License.Features) > 0 {
		fmt.Printf("licensed features: %v\n", result.License.Features)
	}

	fmt.Printf("expires in %d days\n", result.ExpiresIn/(24*3600))
	fmt.Println("成功启动")
}

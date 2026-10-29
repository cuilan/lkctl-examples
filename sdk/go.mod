module sdk

go 1.26.1

require github.com/cuilan/license-key-verify/v2 v2.0.0

// 本仓库与 license-key-verify 同级存放，用 replace 指向本地 checkout 便于离线复现。
// 在你自己的工程中请删除 replace，改用：go get github.com/cuilan/license-key-verify/v2@v2.0.0
replace github.com/cuilan/license-key-verify/v2 => ../../license-key-verify

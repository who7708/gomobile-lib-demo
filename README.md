# Go Mobile Library Demo

In this demo, Go lib (mod) was compiled into xcframework and packaged into a Swift Package to be used as dependency in iOS project.

Please take a look at [使用 Go Mobile 开发跨平台 Library](https://blog.kevinzhow.com/2021/06/22/gomobile_library/)

## This demo works with these 2 repos shows below

[Go Mobile Lib With Swift Package Demo](https://github.com/kevinzhow/go_lib_swift_package_demo)

[Go Mobile With iOS](https://github.com/kevinzhow/go_lib_ios_demo)


# 使用方式
安装 gomobile ，安装好之后，执行
编译 jar 包
gomobile build -v -target android -androidapi 21 ./src/hello

编译 ios 包
gomobile build -v -target ios ./src/hello

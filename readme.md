# qf2 ReadMe
qf2是继Qf之后的新一代面向总线开发的辅助工具包
qf2主要包括以下几个基础模块

## define 定义
qf定义包主要包括qf的各种接口定义，hocon解析器，util包等

## bus 总线
 bus是qf模块之间连接的纽带，也是整个qf的基础，bus分为直通型总线、本地进程总线（IPC）和网络型总线（RPC）等，分别用于连接单体应用 插件型应用和和分体应用

## adapter 访问器
adapter是模块与总线沟通的途径。同样分为直通、插件和网络型三种

## middleware 中间件
middleware中间件可以侦听和拦截qf总线上的所有行为从而对其进行控制和管理。比较常见的有权限控制，数据加密等。中间件本身可以认为是一个特定的插件，并在总线中拥有绝对优先权。当一条总线上存在多个中间件时，需要按照其排序进行优先级排列（如果未设置排序，则按照实例化先后顺序（直通）或文件名排序（插件）自动加载）网络型总线如未配置则不加载任何中间件

## wrapper 打包器
wrapper打包器用于将qf总线及其模块打包成对外可发布的产品
打包器分为 web服务型、c库型、wasm型 aar型等

## qf-SDK sdk访问器
qf-SDK通过不同的编程语言实现对qf的通用访问，分为go型 ,c#型,python型,c++型, js/ts js/ts+wasm等

## qf-SDKBuilder SDK代码生成器
sdkBuilder 通过读取hocon格式的conf配置文件并基于qf-sdkCore生成qf访问代码（model+sdk），分为go型 ,c#型,python型,c++型, js/ts js/ts+wasm等

##　qf-PluginBuilder  插件生成器
pluginBuilder 插件生成器通过读取hocon格式的conf配置文件生成qf插件，分为go型，c#型，python型,c++型等


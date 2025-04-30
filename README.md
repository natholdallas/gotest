# gotest

关于各种 go 自带的库和三方库测试合集，持续更新并跟进文档，可以给初学者一个参考  
测试内容大多数与深度层面的问题相关

## run.sh

run.sh 执行文件使用说明

```shell
./run.sh copier # 运行项目目录下 cmd/copier/main.go 文件, cmd 目录下文件夹名是什么就输入什么
```

## 测试内容与执行参数

| 参数           | 说明                                   |
| -------------- | -------------------------------------- |
| copier         | copier 拷贝时的逻辑                    |
| defer          | 研究 defer 关键字的逻辑                |
| embedded-model | 研究嵌入字段与声明式字段的最终输出逻辑 |
| fiber-return   | fiber 接口中返回各种东西的结果         |

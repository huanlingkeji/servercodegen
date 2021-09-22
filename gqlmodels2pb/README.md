- 将项目pb文件中的json替换成对应的message结构
- 做法说明: 从graphql结构的几个生成文件models_gen.go(asset/bms/gate)中使用ast解析所有struct，然后生成他们对应的pb结构信息。
  从pb文件中查找那些所有未替换的json字段，在models_gen的生成结构信息中找出对应结构，在pb文件中替换成对应的message字段，并将依赖的结构添加到pb 文件中。 对于项目代码中的替换 使用copier生成对应的结构。

- 执行方式：go run . -cmd new-env
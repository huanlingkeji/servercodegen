##用于生成套路代码（会比复制粘贴快些）
- 运行原理：实现genItem的IGenerate接口，用来添加一项代码生成。  
- 该项代码生成可以往某些文件插入内容，也可以生成整个新的文件，它们都是通过go template语法来渲染的  
- 对于要插入内容的文件会先检查是否存在，对于新增的文件会创建完整路径
- 暂时只支持git 处于clean状态时运行生成，生成后的代码会自动添加到git中（即生成后不能再次调用生成）
- 插入内容是根据指定搜索内容位置来进行插入的，如果搜索不到指定内容会报错
- ...

###已有的套路代码
- [x] config配置
- [x] gate
- [x] grphql
- [x] proto
- [x] k8s配置 xx.yaml kustomization.yaml
- [x] wire
- [x] bundle wire(bundle)、rpc、server
- [x] cmd文件
- [x] convert(entity2pb)
- [x] convert(pb2gql)
- 
- [x] entity
- [x] usecase
- [x] service
- [x] repository index
- [x] main

### 将备份的文件删除（暂时不启用）

#### find . -name "*.back.txt"  | xargs rm -f

## TODOList

- [ ] 先做全量的前置检测，待全部检测成功再开始生成代码
- [ ] 补充tmlGenItem里面的实现
- [ ] 实现通用的CURD逻辑
- 
- [ ] 等待泛型的到来基于泛型进行实现...

- 短链创建
    - 对原始链接+用户uuid哈希后，生成对应的短链

        > （1）先查询，判断是否已经存在，如果已经存在 - 直接返回

        > （2）通过元素链接+用户uuid生产对应的短链，进行用户间的数据隔离

- 用户管理
    - 简单注册 - name & password
    - 查询短链 - all or 按原url信息模糊查询
    - 管理锻炼 - add / del

        >  (1) 不支持修改，会直接影响当前在使用短链的所有用户，存在误导、引流的风险。所有被删除的加一个提示

Go + Redis + Sqlite
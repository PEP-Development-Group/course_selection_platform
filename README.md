## TODO

- [x] 右上角悬浮头像，选中**更多信息**没反应。
- [x] 右上角全屏按钮去掉。
- [x] 页面最下方的powered和copyright修改（简单写个跟杭电OJ一样的那种页面？）



- [ ] 课程列表日期可读性优化（待完善）
- [x] 课程列表根据学分筛选，提供学分的下拉框，123456几个选项进行过滤
- [ ] 课程列表考虑根据选课开始和选课结束分一下类，不用每行记录都显示开始和结束



- [x] 左侧菜单栏整理
  - [x] 隐藏了关于我们
  - [x] 文件操作只留下 批量导入
  - [x] 删除工作流
- [x] 课程表增加两个字段，已选人数和总人数
  - [x] 需要前端在添加课程时增加 **总人数** 字段
- [x] 增加**一个**空白菜单，备用
- [ ] 限流：服务端令牌桶/漏桶，客户端多次请求验证码



- [ ] 服务端程序运行时初始化数据库
- [ ] 优化：调整所有字段的大小，建立索引
- [x] 优化：通过token只能查询本人信息，中间件中验证x-token



- [ ] 左侧菜单背景颜色，黑色改一下
- [x] 欢迎您，栾玉国老师；欢迎你，老段同学。分开显示
- [ ] 系统配置 增加个**开学第一天**，**最大取消次数**
- [x] 学生列表菜单查看 **已修学时/应修学时**
- [ ] 和学生成绩页面，增加平均成绩（所有实验计算出来的最终成绩）
- [x] 右上角标签给个颜色~~范围~~反馈，鼠标悬浮样式。
- [x] 学生选课页面，退选后，getlist了 前端页面并没有修改选课人数等信息 *[段]没有复现该效果*
- [x] 关于页面改一下样式，全英/前端归类到一个框里。*[段]目前为中文，不知道wls的昵称和座右铭咋翻译*
- [ ] 首页显示已修/应修，2/18
- [x] 登录页面
- [ ] 首页显示当天日期（xx周周x）
- [ ] 首页课程表区分已选未上课和已选已上课
- [x] 添加课程时存储转换后的时间
- [x] 选课界面【本周】【下周】提示
- [x] 选课界面选课成功后显示模态框回执（elementUI/sweetalert2）
- [ ] 选课界面高峰期要求验证码
- [ ] 选课界面提示退选次数限制
- [ ] 教师上成绩提供旷课选项
- [ ] 【后端/可选】取消选课时增加取消次数字段的值，以保证在学期中可以修改最大取消次数
- [ ] 【商议】学生管理中修改某学生的取消选课次数
- [ ] 选课查询为手机端优化

## 优化

- 学生权限接口缓存
- 准备几千条学生、课程数据。计算索引选择性。覆盖索引，优化查询速度。
- 单用户请求接口速率限制

待定...

## 安全问题

- xss转义？
- sql注入（基本解决）



可选课程列表

- 分光计 - 4学时（这是个下拉框）
  - 栾玉国 - 10周 - 9/18 - 9.30-11.30
  - 某某某 - 10周 - 9/18 - 9.30-11.30

- 重力xxx - 2学时
  - xxx - 11周 - 
  - xxx - 12周



未来待开放课程列表
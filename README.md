# golang-queue
像在python中使用queue一样在golang中使用queue. 线程安全, 支持读阻塞与超时.
# 编写背景
最近提供人像分割服务给业务方时, 发现需要有一个服务分发请求, 开始用python flask框架发现问题挺多, 用golang重写时发现网上实现的golang版队列, 都没有像python那样的Get超时功能. 于是结合网上代码和python queue代码自己写了一个, 看来golang还没完全忘记.
# future work
支持设定固定size, 队列满时Put等待和Put超时

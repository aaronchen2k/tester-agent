import antd from 'ant-design-vue/es/locale-provider/zh_CN'
import momentCN from 'moment/locale/zh-cn'

const components = {
  antLocale: antd,
  momentName: 'zh-cn',
  momentLocale: momentCN
}

const locale = {
  'message': '-',
  'menu.home': '主页',

  'menu.task': '测试任务',
  'menu.execution': '执行中',
  'menu.history': '历史',
  'menu.res': '测试资源',
  'menu.node': 'PVE节点',
  'menu.vm': '虚拟机',
  'menu.container': '容器',

  'menu.dashboard': '仪表盘',
  'menu.dashboard.analysis': '分析页',
  'menu.dashboard.monitor': '监控页',
  'menu.dashboard.workplace': '工作台'
}

export default {
  ...components,
  ...locale
}

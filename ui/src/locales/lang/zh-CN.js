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

  'menu.plan': '测试计划',
  'menu.res': '测试资源',

  'menu.task': '测试任务',
  'menu.execution': '正在执行',
  'menu.history': '历史执行',
  'menu.vm': '虚拟机',
  'menu.container': '容器',

  'menu.dashboard': '仪表盘',
  'menu.dashboard.analysis': '分析页',
  'menu.dashboard.monitor': '监控页',
  'menu.dashboard.workplace': '工作台',

  'action.create': '创建',
  'action.edit': '编辑',
  'action.back': '返回',
  'form.save': '保存',
  'form.reset': '重置',
  'form.cancel': '取消',

  'vm.name': '名称',
  'vm.ident': '标识',
  'vm.osPlatform': '系统平台',
  'vm.osType': '系统类型',
  'vm.osLevel': 'API级别',
  'vm.osLang': '系统语言',
  'vm.osBits': '系统位数',
  'vm.osVer': '系统版本',
  'vm.osBuild': '系统构建号'
}

export default {
  ...components,
  ...locale
}

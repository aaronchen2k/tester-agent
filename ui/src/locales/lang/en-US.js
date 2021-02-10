import antdEnUS from 'ant-design-vue/es/locale-provider/en_US'
import momentEU from 'moment/locale/eu'

const components = {
  antLocale: antdEnUS,
  momentName: 'eu',
  momentLocale: momentEU
}

const locale = {
  'message': '-',
  'menu.home': 'Home',

  'menu.plan': 'Test Plan',
  'menu.res': 'Test Environment',

  'menu.task': 'Test Task',
  'menu.execution': 'Test Execution',
  'menu.history': 'Test History',
  'menu.vm': 'VM',
  'menu.container': 'Container',

  'menu.dashboard': 'Dashboard',
  'menu.dashboard.analysis': 'Analysis',
  'menu.dashboard.monitor': 'Monitor',
  'menu.dashboard.workplace': 'Workplace',

  'layouts.usermenu.dialog.title': 'Message',
  'layouts.usermenu.dialog.content': 'Do you really log-out.',

  'app.setting.pagestyle': 'Page style setting',
  'app.setting.pagestyle.light': 'Light style',
  'app.setting.pagestyle.dark': 'Dark style',
  'app.setting.pagestyle.realdark': 'RealDark style',
  'app.setting.themecolor': 'Theme Color',
  'app.setting.navigationmode': 'Navigation Mode',
  'app.setting.content-width': 'Content Width',
  'app.setting.fixedheader': 'Fixed Header',
  'app.setting.fixedsidebar': 'Fixed Sidebar',
  'app.setting.sidemenu': 'Side Menu Layout',
  'app.setting.topmenu': 'Top Menu Layout',
  'app.setting.content-width.fixed': 'Fixed',
  'app.setting.content-width.fluid': 'Fluid',
  'app.setting.othersettings': 'Other Settings',
  'app.setting.weakmode': 'Weak Mode',
  'app.setting.copy': 'Copy Setting',
  'app.setting.loading': 'Loading theme',
  'app.setting.copyinfo': 'copy success，please replace defaultSettings in src/models/setting.js',
  'app.setting.production.hint': 'Setting panel shows in development environment only, please manually modify',

  'common.notify': 'Notification',

  'action.create': 'Create',
  'action.edit': 'Edit',
  'action.back': 'Back',
  'form.save': 'Save',
  'form.reset': 'Reset',
  'form.cancel': 'Cancel',

  'vm.name': 'Name',
  'vm.ident': 'Key',
  'vm.osPlatform': 'OS Platform',
  'vm.osType': 'OS Type',
  'vm.osLevel': 'API Level',
  'vm.osLang': 'OS Lang',
  'vm.osBits': 'OS Bits',
  'vm.osVer': 'OS Version',
  'vm.osBuild': 'OS Build',
  'vm.updateAll': 'Update all VM with same name',
  'vm.success.update.templ': 'Success to update VM template.'
}

export default {
  ...components,
  ...locale
}

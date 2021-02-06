<template>
  <div>
    <!-- BrowserType _const.BrowserType `gorm:"browserType" json:"browserType,omitempty"`
    BrowserVer  string             `gorm:"browserVer" json:"browserVer,omitempty"`
    BrowserLang _const.SysLang     `gorm:"browserLang" json:"browserLang,omitempty"`-->

    <a-form-model ref="editForm" :model="model" :rules="rules" :label-col="labelCol" :wrapper-col="wrapperCol">
      <a-form-model-item :label="$t('form.osPlatform')" prop="osPlatform">
        <a-select v-model="model.osPlatform">
          <a-select-option v-for="(item, index) in dirs" :value="item.name" :key="index">
            {{item.name}}</a-select-option>
        </a-select>
      </a-form-model-item>

      <a-form-model-item :label="$t('form.osName')" prop="osName">
        <a-select v-model="model.osName">
          <a-select-option v-for="(item, index) in dirs" :value="item.osName" :key="index">
            {{item.name}}</a-select-option>
        </a-select>
      </a-form-model-item>

      <a-form-model-item :label="$t('form.osLevel')" prop="osLevel">
        <a-input v-model="model.osLevel" />
      </a-form-model-item>

      <a-form-model-item :label="$t('form.osLang ')" prop="osLang">
        <a-select v-model="model.osLang">
          <a-select-option v-for="(item, index) in dirs" :value="item.osLang" :key="index">
            {{item.name}}</a-select-option>
        </a-select>
      </a-form-model-item>

      <a-form-model-item :label="$t('form.osBits')" prop="osBits">
        <a-select v-model="model.osBits">
          <a-select-option v-for="(item, index) in dirs" :value="item.osBits" :key="index">
            {{item.name}}</a-select-option>
        </a-select>
      </a-form-model-item>

      <a-form-model-item :label="$t('form.osVer')" prop="osVer">
        <a-input v-model="model.osVer" />
      </a-form-model-item>

      <a-form-model-item :label="$t('form.osBuild')" prop="osBuild">
        <a-input v-model="model.osBuild" />
      </a-form-model-item>

    </a-form-model>

  </div>
</template>

<script>

import { labelColLarge, wrapperColLarge } from '../../../../utils/const'
import Bus from '../../../../components/_util/bus'

export default {
  name: 'VmEdit',
  components: {
  },
  data () {
    return {
      labelCol: labelColLarge,
      wrapperCol: wrapperColLarge,
      rules: {
        title: [
          { required: true, message: this.$i18n.t('valid.required'), trigger: 'change' }
        ],
        fileName: [
          { required: true, message: this.$i18n.t('valid.required'), trigger: 'change' }
        ]
      },
      model: null
    }
  },
  componentDidMount () {
    Bus.$on('modelUpdate', model => {
      console.log('modelUpdate', model)
      this.model = model
    })
  },
  methods: {
  }
}
</script>

<style lang="less" scoped>

</style>

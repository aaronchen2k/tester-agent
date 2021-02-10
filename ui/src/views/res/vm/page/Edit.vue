<template>
  <div>
    <!-- BrowserType _const.BrowserType `gorm:"browserType" json:"browserType,omitempty"`
    BrowserVer  string             `gorm:"browserVer" json:"browserVer,omitempty"`
    BrowserLang _const.SysLang     `gorm:"browserLang" json:"browserLang,omitempty"`-->
    <div v-if="model" class="edit-panel">
      <div class="edit-head">
        <div class="title">
          <span v-if="model.id==0">{{ $t('action.create') }}</span>
          <span v-if="model.id!=0">{{ $t('action.edit') }}</span>
        </div>
        <div class="filter"></div>
        <div class="buttons"></div>
      </div>

      <a-form-model ref="editForm" :model="model" :rules="rules">
        <a-row :gutter="colsFull">
          <a-col :span="colsHalf" style="padding-left: 6px;">
            <a-form-model-item :label="$t('vm.ident')" prop="ident" :labelCol="labelColHalf" :wrapperCol="wrapperColHalf">
              {{ model.ident }}
            </a-form-model-item>
          </a-col>
          <a-col :span="colsHalf">
            <a-form-model-item :label="$t('vm.name')" prop="name" :labelCol="labelColHalf2" :wrapperCol="wrapperColHalf">
              {{ model.name }}
            </a-form-model-item>
          </a-col>
        </a-row>

        <a-row :gutter="colsFull">
          <a-form-model-item :label="$t('vm.osPlatform')" prop="osPlatform" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
            <a-select v-model="model.osPlatform" @change="onOsPlatformChanged">
              <a-select-option v-for="item in osPlatforms" :value="item.code" :key="item.code">
                {{ item.name }}</a-select-option>
            </a-select>
          </a-form-model-item>
        </a-row>

        <a-row :gutter="colsFull">
          <a-form-model-item :label="$t('vm.osType')" prop="osType" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
            <a-select ref="osTypeSelection" v-model="model.osType" @change="onOsTypeChanged">
              <a-select-option v-for="item in osTypes" :value="item.code" :key="item.code">
                {{ item.name }}</a-select-option>
            </a-select>
          </a-form-model-item>
        </a-row>

        <a-row :gutter="colsFull">
          <a-form-model-item :label="$t('vm.osVer')" prop="osVer" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
            <a-input v-model="model.osVer" />
          </a-form-model-item>
        </a-row>

        <a-row :gutter="colsFull">
          <a-form-model-item :label="$t('vm.osLang')" prop="osLang" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
            <a-select v-model="model.osLang">
              <a-select-option v-for="item in osLangs" :value="item.code" :key="item.code">
                {{ item.name }}</a-select-option>
            </a-select>
          </a-form-model-item>
        </a-row>

        <a-row :gutter="colsFull">
          <a-form-model-item :label="$t('vm.osBits')" prop="osBits" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
            <a-select v-model="model.osBits">
              <a-select-option v-for="item in osBits" :value="item" :key="item">
                {{ item }}</a-select-option>
            </a-select>
          </a-form-model-item>
        </a-row>

        <a-row :gutter="colsFull">
          <a-form-model-item prop="updateAll" :wrapperCol="noLabel">
            <a-switch default-checked v-model="model.updateAll" />
            {{ $t('vm.updateAll') }}
          </a-form-model-item>
        </a-row>

        <a-row :gutter="colsFull">
          <a-form-model-item  class="center">
            <a-button @click="save" type="primary">
              {{ $t('form.save') }}
            </a-button>
            <a-button @click="reset" style="margin-left: 10px;">
              {{ $t('form.reset') }}
            </a-button>
          </a-form-model-item>
        </a-row>

      </a-form-model>
    </div>
  </div>
</template>

<script>

import {
  colsFull,
  colsHalf,
  labelColFull,
  labelColHalf,
  labelColHalf2,
  wrapperColFull,
  wrapperColHalf,
  noLabel } from '../../../../utils/const'
import Bus from '../../../../components/_util/bus'
import { loadVmTempl, saveVmTempl, listEnv } from '@/api/manage'

export default {
  name: 'VmEdit',
  components: {
  },
  data () {
    return {
      colsFull: colsFull,
      colsHalf: colsHalf,
      labelColFull: labelColFull,
      wrapperColFull: wrapperColFull,
      labelColHalf: labelColHalf,
      labelColHalf2: labelColHalf2,
      wrapperColHalf: wrapperColHalf,
      noLabel: noLabel,

      rules: {
        osPlatform: [
          { required: true, message: this.$i18n.t('valid.required'), trigger: 'change' }
        ],
        osName: [
          { required: true, message: this.$i18n.t('valid.required'), trigger: 'change' }
        ],
        osLang: [
          { required: true, message: this.$i18n.t('valid.required'), trigger: 'change' }
        ]
      },

      osPlatforms: [],
      osTypesAll: [],
      osTypes: [],
      osLangs: [],
      osBits: [64, 32],
      model: null
    }
  },
  created () {
    listEnv().then(json => {
      console.log('listEnv', json)

      this.osPlatforms = json.data.osPlatforms
      this.osTypesAll = json.data.osTypes
      this.osLangs = json.data.osLangs
    })
  },
  mounted () {
    console.log('mounted')
    Bus.$on('onVmNodeSelected', node => {
      console.log('onVmNodeSelected', node)
      if (node || node.ident) {
        this.model = null
      }

      const data = { name: node.name, type: node.type, isTemplate: node.isTemplate, ident: node.ident, node: node.node, cluster: node.cluster }
      loadVmTempl(data).then(json => {
        console.log('loadVmTempl', json)
        this.model = json.data
        this.onOsPlatformChanged()
      })
    })
  },
  destroyed () {
    Bus.$off('onVmNodeSelected')
  },
  methods: {
    save () {
      this.$refs.editForm.validate(valid => {
        console.log(valid, this.model)
        if (!valid) {
          console.log('validation fail')
          return
        }
        saveVmTempl(this.model).then(json => {
          console.log('saveVmTempl', json)
          if (json.code === 200) {
            this.$notification['success']({
              message: this.$i18n.t('common.notify'),
              description: this.$i18n.t('vm.success.update.templ'),
              duration: 8
            })
          }
        })
      })
    },
    reset () {
      this.$refs.editForm.resetFields()
    },

    onOsPlatformChanged () {
      console.log('onOsPlatformChanged')

      const osTypes = []
      // this.model.osType = undefined

      this.osTypesAll.forEach((item, index) => {
        if (item.osPlatform === this.model.osPlatform) {
          osTypes.push(item)
        }
      })

      this.osTypes = osTypes
    },
    onOsTypeChanged (val) {
      // this.model.osType = val
      // console.log('onOsTypeChanged', val, this.model.osType)
    }
  }
}
</script>

<style lang="less" scoped>

</style>

<template>
  <div>
    <div class="main">
      <div class="left" :style="{height: colHeight}">
        <container-tree>
          ref="containerTreePage"
        </container-tree>
      </div>
      <div class="right" :style="{height: colHeight}">
        <container-edit
          ref="containerEditPage">
        </container-edit>
      </div>
    </div>

  </div>
</template>

<script>

import { ContainerTree, ContainerEdit } from './page'

export default {
  name: 'ContainerList',
  components: {
    ContainerTree, ContainerEdit
  },
  data () {
    return {

    }
  },
  computed: {
    colHeight () {
      return (document.documentElement.clientHeight - 140) + 'px'
    }
  },
  mounted () {
  },
  created () {

  },
  methods: {
  }
}
</script>

<style lang="less" scoped>
.main {
  display: flex;
  padding: 10px;
  background-color: #ffffff;

  .left {
    width: 360px;
    border-right: 1px solid #DDDDDD;
    overflow: auto;
  }
  .right {
    flex: 1;
    overflow: auto;
  }
}
</style>

<template>
  <div>
    <div class="main">
      <div class="left" :style="{height: colHeight}">
        <vm-tree>
          ref="vmTreePage"
        </vm-tree>
      </div>
      <div class="right" :style="{height: colHeight}">
        <vm-edit
          ref="vmEditPage">
        </vm-edit>
      </div>
    </div>

  </div>
</template>

<script>
import { VmTree, VmEdit } from './page'

export default {
  name: 'VmList',
  components: {
    VmTree, VmEdit
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
      width: 260px;
      border-right: 1px solid #DDDDDD;
      overflow: auto;
    }
    .right {
      flex: 1;
      overflow: auto;
    }
  }
</style>

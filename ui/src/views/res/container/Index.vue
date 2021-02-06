<template>
  <div>
    <div class="main">
      <div class="left" :style="{height: colHeight}">

      </div>
      <div class="right" :style="{height: colHeight}">

      </div>
    </div>

  </div>
</template>

<script>

export default {
  name: 'ContainerList',
  components: {
  },
  data () {
    return {

    }
  },
  computed: {
    isRoot () {
      console.log('isRoot', this.selectNode)
      return !this.selectNode.parentID || this.selectNode.parentID === 0 || this.selectNode.id === 0
    },
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

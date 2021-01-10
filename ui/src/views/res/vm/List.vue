<template>
  <div>
    <div class="main">
      <div class="left" :style="{height: colHeight}">
        <div>
          <a href="#" @click="expand">
            <span v-if="!expanded">展开全部</span>
            <span v-if="expanded">收缩全部</span>
          </a>
        </div>
        <div>
          <a-tree
            ref="machineTree"
            class="draggable-tree"
            :show-line="false"
            :show-icon="true"
            :expandedKeys.sync="openKeys"
            :selectedKeys.sync="selectedKeys"
            :tree-data="[treeData]"
            :replaceFields="fieldMap"
            @select="onSelect"
            @rightClick="onRightClick"
            :draggable="false"
          >
            <template slot="custom" slot-scope="{ type,isTemplate }">
              <a-icon v-if="type=='cluster'" type="cluster" />
              <a-icon v-else-if="type=='node'" type="cloud-server" />
              <a-icon v-else-if="type=='vm' && !isTemplate" type="desktop" />
              <a-icon v-else-if="type=='vm' && isTemplate" type="build" />
            </template>
          </a-tree>
        </div>

        <div v-if="rightClickNode" :style="rightMenuStyle" class="tree-context-menu">
          <a-menu @click="menuClick" mode="inline" class="menu">
            <a-menu-item key="addNeighbor" v-if="!isRoot">
              <a-icon type="plus" />{{ $t('msg.design.create.brother') }}
            </a-menu-item>
            <a-menu-item key="addChild" v-if="type=='def'|| ((type=='ranges' || type=='instances') && isRoot)">
              <a-icon type="plus" />{{ $t('msg.design.create.child') }}
            </a-menu-item>
            <a-menu-item key="remove" v-if="!isRoot">
              <a-icon type="delete" />{{ $t('msg.design.remove.node') }}
            </a-menu-item>
          </a-menu>
        </div>
      </div>
      <div class="right" :style="{height: colHeight}">

      </div>
    </div>

  </div>
</template>

<script>

import { listVm } from '@/api/manage'

export default {
  name: 'VmList',
  components: {
  },
  data () {
    return {
      models: [],
      treeData: {},
      openKeys: [],
      selectedKeys: [],

      nodeMap: {},
      fieldMap: { title: 'name', value: 'id' },
      selectNode: null,
      rightClickNode: null,
      rightMenuStyle: {},
      rightVisible: false,
      expanded: false
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
    console.log('created')
    listVm().then(json => {
      console.log('listVm', json)
      this.treeData = json.data
      this.loadTreeCallback(this.treeData, '')
    })
  },
  methods: {
    getModel (id) {
      console.log('getModel', id)

      if (this.type === 'host') {
        // getDefField(id).then(res => {
        //   console.log('getDefField', res)
        //   this.modelData = res.data
        //   this.time2 = Date.now() // trigger data refresh
        // })
      } else if (this.type === 'node') {
        // getResRangesItem(id).then(res => {
        //   console.log('getResRangesItem', res)
        //   this.modelData = res.data
        //   this.time2 = Date.now() // trigger data refresh
        // })
      } else if (this.type === 'vm') {
        // getResInstancesItem(id).then(res => {
        //   console.log('getResInstancesItem', res)
        //   this.modelData = res.data
        //   this.time2 = Date.now() // trigger data refresh
        // })
      }
    },
    loadTreeCallback (data, selectedKey) {
      this.getOpenKeys(data)

      if (selectedKey) {
        this.getModel(selectedKey)
        this.rightVisible = true
      } else {
        this.rightVisible = false
      }
    },

    onSelect (selectedKeys, e) { // selectedKeys, e:{selected: bool, selectedNodes, node, event}
      console.log('onSelect', selectedKeys, e.selectedNodes, e.node, e.node.eventKey)
      if (selectedKeys.length === 0) {
        selectedKeys[0] = e.node.eventKey // keep selected
      }

      const node = this.nodeMap[e.node.eventKey]
      console.log('node', node)
      if ((this.type === 'def' && node.parentID === 0) || (this.type === 'config' && node.id === 0) ||
          (node.fields && node.fields.length > 0)) {
        this.rightVisible = false
        this.modelData = {}
      } else {
        this.rightVisible = true
        this.tabKey = 'info'
      }

      // this.getModel(parseInt(selectedKeys[0]))
    },
    menuClick (e) {
      console.log('menuClick', e, this.rightClickNode)
      this.clearMenu()
    },
    onRightClick ({ event, node }) {
      event.preventDefault()
      console.log('onRightClick', node)

      const y = event.currentTarget.getBoundingClientRect().top
      const x = event.currentTarget.getBoundingClientRect().right

      this.rightClickNode = {
        pageX: x,
        pageY: y,
        id: node._props.eventKey,
        title: node._props.title,
        parentID: node._props.dataRef.parentID || null
      }

      this.rightMenuStyle = {
        position: 'fixed',
        maxHeight: 40,
        textAlign: 'center',
        left: `${x + 10 - 0}px`,
        top: `${y + 6 - 0}px`
        // display: 'flex',
        // flexDirection: 'row'
      }
    },
    clearMenu () {
      console.log('clearMenu')
      this.rightClickNode = null
    },
    expand () {
      this.expanded = !this.expanded
      this.openKeys = []
      this.getOpenKeys(this.treeData, this.expanded)
    },
    getOpenKeys (node, expanded) {
      if (!node) return

      node.scopedSlots = {
        icon: 'custom'
      }

      if (expanded || node.type === 'root' || node.type === 'cluster') {
        this.openKeys.push(node.key)
        this.nodeMap[node.key] = node

        if (node.children) {
          node.children.forEach((item) => {
            this.getOpenKeys(item, expanded)
          })
        }
      }
    }

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
    }
    .right {
      flex: 1;
    }
  }
</style>

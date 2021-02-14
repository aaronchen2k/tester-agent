<template>
  <div>
    <input id="input" type="text" v-model="inputModel" />
    <button id="sendBtn" @click="onClick">Send</button>
    <pre id="output">{{ outputModel }}</pre>
  </div>
</template>

<script>
import TaskForm from '../../list/modules/TaskForm'
import Info from '../../list/components/Info'
import { WsApi, listPlan } from '@/api/manage'
import * as neffos from 'neffos.js'

export default {
  name: 'PlanList',
  components: {
    TaskForm,
    Info
  },
  data () {
    return {
      data: [],
      status: 'all',
      wsConn: null,
      inputModel: 'abc',
      outputModel: ''
    }
  },
  computed: {
  },
  mounted () {
    console.log('mounted')
  },
  created () {
    console.log('created')
    listPlan().then(json => {
      console.log('listPlan', json)
      this.data = json.data
    })

    this.initConn()
  },

  methods: {
    async initConn () {
      const that = this
      try {
        const conn = await neffos.dial(WsApi, {
          default: {
            _OnNamespaceConnected: (nsConn, msg) => {
              if (nsConn.conn.wasReconnected()) {
                that.addMessage('re-connected after ' + nsConn.conn.reconnectTries.toString() + ' trie(s)')
              }

              that.addMessage('connected to namespace: ' + msg.Namespace)
              that.wsConn = nsConn
            },
            _OnNamespaceDisconnect: (nsConn, msg) => {
              that.addMessage('disconnected from namespace: ' + msg.Namespace)
            },
            OnChat: (nsConn, msg) => {
              console.log('OnChat')
              that.addMessage(msg.Body)
            },
            OnVisit: (nsConn, msg) => {
              console.log('OnVisit', msg)
            }
          }
        })
        await conn.connect('default')
      } catch (err) {
        console.log(err)
      }
    },

    addMessage (msg) {
      this.outputModel += msg + '\n'
    },

    handleError (reason) {
      console.log(reason)
      window.alert(reason)
    },

    onClick () {
      // console.log('onClick', this.inputModel, this.wsConn)
      this.wsConn.emit('OnChat', 'this.inputModel')
      this.outputModel += 'Me: ' + this.inputModel + '\n'
    }
  }
}
</script>

<style lang="less" scoped>
</style>

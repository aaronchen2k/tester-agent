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
import { WsApi } from '@/api/manage'
import * as neffos from 'neffos.js'

export default {
  name: 'GitLabList',
  components: {
    TaskForm,
    Info
  },
  data () {
    return {
      data: [],
      status: 'all',
      wsConn: null,
      room1: null,
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
              that.wsConn.joinRoom('room1')
            },
            _OnNamespaceDisconnect: (nsConn, msg) => {
              that.addMessage('disconnected from namespace: ' + msg.Namespace)
            },
            OnChat: (nsConn, msg) => {
              console.log('OnChat')
              that.addMessage(msg.Room + ': ' + msg.Body)
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

    onClick () {
      // console.log('onClick', this.inputModel, this.wsConn)
      this.wsConn.room('room1').emit('OnChat', 'this.inputModel')
      this.outputModel += 'Me: ' + this.inputModel + '\n'
    }
  }
}
</script>

<style lang="less" scoped>
</style>

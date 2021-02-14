<template>
  <div>
    <input id="input" type="text" value="ABC" />
    <button id="sendBtn" disabled>Send</button>
    <pre id="output"></pre>
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
      socket: null
    }
  },
  computed: {
  },
  mounted () {

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
      console.log('---initConn')
      const that = this
      try {
        const conn = await neffos.dial(WsApi, {
          default: {
            _OnNamespaceConnected: function (nsConn, msg) {
              if (nsConn.conn.wasReconnected()) {
                that.addMessage('1 re-connected after ' + nsConn.conn.reconnectTries.toString() + ' trie(s)')
              }
              that.addMessage('2 connected to namespace: ' + msg.Namespace)
              that.handleNamespaceConnectedConn(nsConn)
            },
            _OnNamespaceDisconnect: function (nsConn, msg) {
              that.addMessage('3 disconnected from namespace: ' + msg.Namespace)
            },
            OnChat: function (nsConn, msg) { // "chat" event.
              console.log('OnChat')
              that.addMessage('4 ' + msg.Body)
            }
          }
        })
        conn.connect('default')
      } catch (err) {
        console.log(err)
      }
    },

    addMessage (msg) {
      const outputTxt = document.getElementById('output')
      outputTxt.innerHTML += msg + '\n'
    },

    handleError (reason) {
      console.log(reason)
      window.alert(reason)
    },

    handleNamespaceConnectedConn (nsConn) {
      console.log('---handleNamespaceConnectedConn')
      const inputTxt = document.getElementById('input')
      const sendBtn = document.getElementById('sendBtn')

      sendBtn.disabled = false
      const that = this
      sendBtn.onclick = function () {
        const input = inputTxt.value
        nsConn.emit('OnChat', input)
        that.addMessage('Me: ' + input)
      }
    }
  }
}
</script>

<style lang="less" scoped>
</style>

<template>
  <div>
    <input id="input" type="text" />
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

    this.initConn()
    listPlan().then(json => {
      console.log('listPlan', json)
      this.data = json.data
    })
  },

  methods: {
    initConn () {
      const events = {}
      events._OnNamespaceConnected = function (nsConn, msg) {
        if (nsConn.conn.wasReconnected()) {
          this.addMessage('re-connected after ' + nsConn.conn.reconnectTries.toString() + ' trie(s)')
        }

        this.addMessage('connected to namespace: ' + msg.Namespace)
        this.handleNamespaceConnectedConn(nsConn)
      }

      events._OnNamespaceDisconnect = function (nsConn, msg) {
        this.addMessage('disconnected from namespace: ' + msg.Namespace)
      }

      events.chat = function (nsConn, msg) { // "chat" event.
        this.addMessage(msg.Body)
      }

      /* OR regiter those events as:
          neffos.dial(wsURL, {default: {
              chat: function (nsConn, msg) { [...] }
          }});
      */

      // If "await" and "async" are available, use them instead^, all modern browsers support those,
      // so all of the examples will be written using async/await method instead of promise then/catch callbacks.
      // A usage example of promise then/catch follows:
      const that = this
      neffos.dial(WsApi, {
        default: { // "default" namespace.
          _OnNamespaceConnected: function (ns, msg) {
            that.addMessage('connected to namespace: ' + msg.Namespace)
          },
          _OnNamespaceDisconnect: function (ns, msg) {
            that.addMessage('disconnected from namespace: ' + msg.Namespace)
          },
          OnChat: function (ns, msg) { // "chat" event.
            that.addMessage(msg.Body)
          }
        }
      }).then(function (conn) {
        conn.connect('default').then(that.handleNamespaceConnectedConn).catch(that.handleError)
      }).catch(that.handleError)
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
      const inputTxt = document.getElementById('input')
      const sendBtn = document.getElementById('sendBtn')

      sendBtn.disabled = false
      const that = this
      sendBtn.onclick = function () {
        const input = inputTxt.value
        inputTxt.value = ''
          nsConn.emit('OnChat', input)
        that.addMessage('Me: ' + input)
      }
    }
  }
}
</script>

<style lang="less" scoped>
</style>

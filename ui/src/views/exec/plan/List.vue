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
              chat: function (ns, msg) { // "chat" event.
                that.addMessage(msg.Body)
              }
          }
      }).then(function (conn) {
          conn.connect('default').then(that.handleNamespaceConnectedConn).catch(that.handleError)
      }).catch(that.handleError)

      // try {
      //   // You can omit the "default" namespace and simply define only Events,
      //   // the namespace will be an empty string"",
      //   // however if you decide to make any changes on
      //   // this example make sure the changes are reflecting inside the ../server.go file as well.
      //   //
      //   // At "wsURL" you can put the relative URL if the client and server
      //   // hosted in the same address, e.g. "/echo".
      //   const conn = await neffos.dial(WsApi, { default: events }, {
      //     // if > 0 then on network failures it tries to reconnect every 5 seconds, defaults to 0 (disabled).
      //     reconnect: 5000,
      //     // custom headers:
      //     // headers: {
      //     //    'X-Username': 'kataras',
      //     // }
      //   })
      //
      //   // You can either wait to conenct or just conn.connect("connect")
      //   // and put the `handleNamespaceConnectedConn` inside `_OnNamespaceConnected` callback instead.
      //   // const nsConn = await conn.connect("default");
      //   // handleNamespaceConnectedConn(nsConn);
      //   conn.connect("default")
      //
      // } catch (err) {
      //   this.handleError(err)
      // }

    //   const socket = new ReconnectingWebSocket(WsApi)
    //   this.socket = socket
    //   this.socket.debug = true
    //   this.socket.onmessage = this.OnMessage
    //   this.socket.onopen = this.OnOpen
    //
    //   const that = this
    //   setInterval(() => {
    //     console.log('===')
    //
    //     const mes = { msg: 'test' }
    //     mes.type = 'test'
    //     that.socket.send(JSON.stringify(mes))
    //   }, 1000)
    // },
    // OnOpen () {
    //   console.log('OnOpen')
    //   const mes = { msg: 'test' }
    //   mes.type = 'test'
    //   this.socket.send(JSON.stringify(mes))
    // },
    // OnMessage (e) {
    //   console.log('OnMessage')
    //   const resp = JSON.parse(e.data)
    //   console.log(resp)
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
        console.log(nsConn)
        nsConn.emit('chat', input)
        that.addMessage('Me: ' + input)
      }
    }
  }
}
</script>

<style lang="less" scoped>
</style>

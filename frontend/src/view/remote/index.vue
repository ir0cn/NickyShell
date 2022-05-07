<template>
  <div class="viewport" ref="terminal"/>
</template>

<script>

import {Terminal} from "xterm";
import {FitAddon} from "xterm-addon-fit";
import "xterm/css/xterm.css"

export default {
  name: "remote",
  data() {
    return {
      term: null,
      ws: null,
      fitAddon: null,
      wsUrl: "",
      token: '',
    }
  },
  methods: {
    initWsUrl() {
      let url = window.location.href
      let protocol = "wss://"
      if (url.substring(0, 7) === 'http://') {
        protocol = "ws://"
      }
      this.wsUrl = protocol + window.location.host + '/api/v1/assets/localhost/connect'
      // this.wsUrl = 'ws://127.0.0.1:8081/api/v1/assets/localhost/connect'
      // console.log('wsurl:' + this.wsUrl)
    },
    initTerminal() {
      this.term = new Terminal({
        cursorBlink: true, // 光标闪烁
        cursorStyle: "block",
        fontSize: 13,
        fontFamily: "Monaco, Menlo, Consolas, 'Courier New', monospace",
        theme: {
          background: '#181d28',
          cursor: '#FFD700',
        },
      })
      // let terminalContainer = document.getElementById("terminal");
      this.term.open(this.$refs['terminal'])
      this.fitAddon = new FitAddon();
      this.term.loadAddon(this.fitAddon);
      this.term.onData(data => this.ws.send(JSON.stringify({data})))
      this.term.focus()
      window.addEventListener('resize', this.resize)
    },
    initWebSocket() {
      this.ws = new WebSocket(this.wsUrl, [this.token])
      this.ws.onopen = () => {
        this.resize()
      }
      this.ws.onclose = () => {
        console.log("ws closed")
      }
      this.ws.onerror = () => {
        console.log("ws connect error" )
      }
      this.ws.onmessage = (msg) => {
        this._read_as_text(msg.data)
      }
    },
    _read_as_text(data) {
      const reader = new window.FileReader()
      reader.onload = () => this.term.write(reader.result)
      reader.readAsText(data, 'utf-8')
    },
    resize() {
      // console.log('resize')
      this.fitAddon.fit()
      let cols = this.term.cols
      let rows = this.term.rows
      // console.log(this.$refs['terminal'].offsetWidth)
      // console.log(cols, rows)
      this.ws.send(JSON.stringify({cols, rows}))
    },
  },
  mounted() {
    this. token = sessionStorage.getItem('token')
    if (this.token === null || this.token === ""){
      this.$router.push('/login')
    }
    this.initWsUrl()
    this.initTerminal()
    this.initWebSocket()
  }

}
</script>

<style scoped>
.viewport{
  width: 100%;
  height: 100%;
}
</style>

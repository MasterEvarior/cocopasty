<template>
  <div>
    <h1 class="title">Cocopasty</h1>
    <div class="container">
      <div class="subtitle">
        <p>
        Copy-and-paste your code, open this site on another device and then copy-and-paste again :)
        </p>
      </div>
      <PrismEditor
        class="my-editor"
        v-model="code" 
        :highlight="highlighter" 
        line-numbers
      />
      <button class="button save-button" @click="pushData">Save</button>
      <button class="button copy-button" @click="copyToClipboard">Copy</button>
    </div>
  </div>
</template>

<script>
import { PrismEditor } from 'vue-prism-editor';
import 'vue-prism-editor/dist/prismeditor.min.css';
import hljs from 'highlight.js';
import 'highlight.js/styles/default.css';
import getEnv from '@/utils/env';

export default {
  name: "App",
  components: {
    PrismEditor
  },
  mounted() {
    this.getData()
  },
  methods: {
    highlighter(code){
      return hljs.highlightAuto(code).value
    },
    getData(){
      let backendError = false;

      fetch(this.backendUrl)
        .then(response => {
          if (!response.ok) {
            console.error(response.status)
          }
          return response.json();
        })
        .then(json => {
          this.code = json.Code;
        })
        .catch(function() {
          backendError = true;
        })

        if(backendError){
          this.$toast.error(`Error while communicating with backend...`)
        }
    },
    pushData(){
      const options = {
        method: 'POST',
        body: JSON.stringify({"Code": this.code}),
        headers: {'Content-Type': 'application/json'}
      }

      let backendError = false;

      this.$toast.show(`Trying to save your code...`)

      fetch(this.backendUrl, options)
        .then(response => {
          if (!response.ok) {
            console.error(response.status)
          }
          return response.json();
        })
        .catch(function(){
          backendError = true;
        })

        if(backendError){
          this.$toast.error(`Error while communicating with backend...`)
          return
        }

        this.$toast.success(`Data successfully saved!`)
    },
    copyToClipboard(){
      navigator.clipboard.writeText(this.code);
      this.$toast.success(`Snippet copied into clipboard!`)
    }
  },
  data() {
    return {
      code: '',
      backendUrl: getEnv('VUE_APP_BACKEND_HOST') + ':' + getEnv('VUE_APP_BACKEND_PORT'),
    };
  }
};
</script>

<style lang="scss">

$body_width: 75%;

body {
  font-family: -apple-system, BlinkMacSystemFont, Segoe UI Variable, Segoe UI,
    system-ui, ui-sans-serif, Helvetica, Arial, sans-serif, Apple Color Emoji,
    Segoe UI Emoji;
  margin: 0;
  background: #181916;
}
h1,
p {
  color: white;
  a {
    font-family: Consolas, Monaco, monospace;
  }
}
h1 {
  font-family: 'Roboto Mono';
  margin: 5% 0;
  font-size: 46px;
  text-align: center;
}
.subtitle {
  text-align: center;
  & + div {
    margin-top: 60px;
  }
}
.container {
  margin: 0 auto;
  margin-bottom: 1%;
  max-width: $body_width;
}
.my-editor {
    background: #2d2d2d;
    color: #ccc;

    font-family: Fira code, Fira Mono, Consolas, Menlo, Courier, monospace;
    font-size: 14px;
    line-height: 1.5;
    padding: 5px;
}

.save-button {
	border-color:#18ab29;
	text-shadow:0#2f6627;
}

.copy-button {
  margin-right: 10px;
}

.button {
  margin-top: 2%;
  float: right;

  background-color:transparent;
	border-radius:5px;
	border:1px solid;
	display:inline-block;
	cursor:pointer;
	color:#ffffff;
	font-family:Arial;
	font-size:17px;
	padding: 0.5% 1%;
	text-decoration: none;
	text-shadow:0px 1px 0px;
}

.button:active {
	position:relative;
	top:1px;
}


.prism-editor__textarea:focus {
  outline: none;
}
</style>
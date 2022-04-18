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
        body: JSON.stringify({"Code": this.code, "Language": 'js'}),
        headers: {'Content-Type': 'application/json'}
      }

      let backendError = false;

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
        }
    }
  },
  data() {
    return {
      code: '',
      backendUrl: 'http://' + getEnv('VUE_APP_BACKEND_HOST') + ':' + getEnv('VUE_APP_BACKEND_PORT'),
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
    /* we dont use `language-` classes anymore so thats why we need to add background and text color manually */
    background: #2d2d2d;
    color: #ccc;

    /* you must provide font-family font-size line-height. Example: */
    font-family: Fira code, Fira Mono, Consolas, Menlo, Courier, monospace;
    font-size: 14px;
    line-height: 1.5;
    padding: 5px;
}

.prism-editor__textarea:focus {
  outline: none;
}
@media screen and (max-width: 560px) {
  .button_group {
    flex-wrap: wrap;
    margin-top: 0;
    button {
      width: calc(50% - 10px);
      overflow: hidden;
      white-space: nowrap;
      text-overflow: ellipsis;
      margin-top: 20px;
      padding: 12px 12px;
    }
  }
}
</style>
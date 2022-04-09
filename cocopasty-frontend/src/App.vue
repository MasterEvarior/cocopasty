<template>
  <div>
    <h1 class="title">Cocopasty</h1>
    <div class="container">
      <div class="subtitle">
        <p>
        Copy-and-paste your code, open this site on another device and then copy-and-paste again :)
        </p>
      </div>
      <CodeEditor
        :class="theme"
        :theme="theme"
        value=""
        width="100%"
        height="450px"
        :language_selector="true"
        v-model="code"
        :languages="languages"
      ></CodeEditor>
    </div>
  </div>
</template>

<script>
import CodeEditor from 'simple-code-editor';

export default {
  name: "App",
  components: {
    CodeEditor
  },
  mounted() {
    this.getData(),
    this.timer = setInterval(() => {
    this.shouldSave()
  }, 15000)
  },
  methods: {
    shouldSave(){
      console.log("hey :)")
      if(this.code !== this.oldCode){
        console.log("Code changed wooow!")
        this.oldCode = this.code;
        this.$toast.show('Trying to save your code snippet...')
        this.pushData()
      }
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
      oldCode: '',
      timer: null,
      backendUrl: 'http://' + process.env.VUE_APP_BACKEND_HOST + ':' + process.env.VUE_APP_BACKEND_PORT,
      languages: [
        ['javascript', 'JS'],
        ['python', 'Python'],
        ['xml', 'XML'],
        ['bash','Bash'],
        ['c','C'],
        ['csharp','C#'],
        ['css','CSS'],
        ['markdown','Markdown'],
        ['dart','Dart'],
        ['django','Django'],
        ['ruby','Ruby'],
        ['go','Go'],
        ['java','Java'],
        ['javascript','JavaScript'],
        ['json','JSON'],
        ['python','Python'],
        ['r','R'],
        ['rust','Rust'],
        ['shell','Shell'],
        ['sql','SQL'],
        ['swift','Swift'],
        ['yaml','YAML'],
        ['typescript','TypeScript']
      ]
    };
  },
  beforeUnmount() {
    clearInterval(this.timer)
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
.code_editor {
  & + .code_editor {
    margin-top: 32px;
  }
  & + p {
    margin-top: 32px;
  }
}
.container {
  margin: 0 auto;
  margin-bottom: 1%;
  max-width: $body_width;
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
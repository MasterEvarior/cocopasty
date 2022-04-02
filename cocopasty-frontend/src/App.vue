<template>
  <div>
    <h1 class="title">Cocopasty</h1>
    <div class="container">
      <div class="subtitle">
        <p>
        Copy-and-paste your code, open up this site on another device and then copy-and-paste again :)
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
    CodeEditor,
  },
  data() {
    return {
      code: '',
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
  margin: 50px 0;
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
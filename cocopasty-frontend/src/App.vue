<template>
  <div class="container">
    <h1>Cocopasty</h1>
    <CodeEditor
      :languages="languages"
      :language_selector="true"
      :autofocus="true"
      width="100%"
      v-model="code"/>
  </div>
</template>

<script>
import CodeEditor from 'simple-code-editor';

export default {
  name: 'App',
  components: {
    CodeEditor
  },
  data() {
    return {
      code: '',
      languages: [
        ['javascript', 'JS'],
        ['python', 'Python'],
        ['brainfuck','Brainfuck'],
        ['applescript','AppleScript'],
        ['xml', 'XML'],
        ['bash','Bash'],
        ['basic','Basic'],
        ['c','C'],
        ['clojure','Clojure'],
        ['coffeescript','CoffeeScript'],
        ['csharp','C#'],
        ['css','CSS'],
        ['markdown','Markdown'],
        ['dart','Dart'],
        ['django','Django'],
        ['dockerfile','Dockerfile'],
        ['ruby','Ruby'],
        ['fortran','Fortran'],
        ['go','Go'],
        ['haskell','Haskell'],
        ['java','Java'],
        ['javascript','JavaScript'],
        ['json','JSON'],
        ['kotlin','Kotlin'],
        ['lisp','Lisp'],
        ['lua','Lua'],
        ['perl','Perl'],
        ['powershell','PowerShell'],
        ['python','Python'],
        ['r','R'],
        ['rust','Rust'],
        ['scala','Scala'],
        ['scilab','Scilab'],
        ['scss','SCSS'],
        ['shell','Shell'],
        ['sql','SQL'],
        ['swift','Swift'],
        ['yaml','YAML'],
        ['typescript','TypeScript'],
        ['vbscript','VBScript'],
        ['vim','Vim']
      ]
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.container {
  margin: 0 auto;
  max-width: 580px; 
}
</style>

<template>
  <div class="home">
    <h1>{{name}}</h1>
    <ul>
   <li v-for="todo in todos" :key="todo.name">
    {{todo.name}} 
  </li>
    </ul>
  </div>
</template>

<script>

export default {
  name: 'Home',
  data(){
    return{
      name:"todos",
      todos:[{name:"tes"},{name:"we"}]
    }
  },
  methods:{
    getTodos(){
      fetch("http://localhost:3000/todos/")
      .then(res => (res.json()))
      .then(data =>(this.todos=data))
    }
  },
  mounted(){
    this.getTodos();
  },
  
}
</script>

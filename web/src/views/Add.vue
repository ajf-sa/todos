<template>
  <div class="add">
  <input v-model="todo" />
  <button @click="setTodo" >Add</button>
  </div>
</template>

<script>

export default {
  name: 'Add',
  data(){
    return{
      todo:"",
     
    }
  },
  methods:{
    setTodo(){
      if (undefined === this.todo || this.todo.length <=0){
        return
      }
      const requestOptions ={
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body: JSON.stringify({name:this.todo})
      };
      fetch("http://localhost:3000/todo/create/",requestOptions)
      .then(res => (res.json()))
      .then(data =>{
        this.todo ="";
        return this.todos=data})
    }
  },
  mounted(){
    
  },
  
}
</script>

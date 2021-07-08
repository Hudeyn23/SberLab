<template>
  <div>
    <input
        v-model="newTodoText"
        placeholder="New todo text"
        v-on:keydown.enter="addTodo"
    />
    <ul v-if="todos.length">
      <TodoListItem
          v-for="todo in todos"
          v-bind:key="todo.id"
          v-bind:todo="todo"
          v-on:remove="removeTodo"
          v-on:update_status="updateTodo"
      />
    </ul>
    <p v-else>
      Nothing left in the list. Add a new todo in the input above.
    </p>
    <p>Number of todos: {{ todosCount }}</p>
    <div id="button">
      <button v-on:click="counter += 1">Add 1</button>
      <p>The button above has been clicked {{ counter }} times.</p>
    </div>
  </div>
</template>

<script>
import TodoListItem from './TodoListItem.vue'
import axios from "axios"

const axios_instance = axios.create({
  baseURL: 'http://45.9.24.240:8080/products/',
  // baseURL: 'http://localhost:8080/products/',
});

export default {
  components: {
    TodoListItem
  },


  data() {
    return {
      newTodoText: '',
      todos: [],
      counter: 0,
    }
  },
  created() {
    axios_instance.get().then(result => {
      console.log(result.data)
      result.data.forEach(element => {
        this.todos.push(element)
      })
    }, error => {
      console.error(error);
    });
  },
  computed: {
    // a computed getter
    todosCount() {
      // `this` points to the vm instance
      return this.todos.length
    }
  },
  methods: {
    addTodo() {
      const trimmedText = this.newTodoText.trim()
      if (trimmedText) {
        axios_instance.post(
            "",
            {
              title: trimmedText,
              completed: false
            }
        ).then(result => {
          this.todos.push(result.data)
          this.newTodoText = ''
        }, error => {
          console.error(error);
        });
      }
    },
    removeTodo(todoToRemove) {
      axios_instance.delete(
          todoToRemove.id
      ).then(() => {
        this.todos = this.todos.filter(todo => {
          return todo.id !== todoToRemove.id
        })
      }, error => {
        console.error(error);
      });
    },
    updateTodo(todoToUpdate) {
      console.log(todoToUpdate)
      axios_instance.put(
          todoToUpdate.id,
          {
            title: todoToUpdate.title,
            completed: !todoToUpdate.completed
          }
      ).then(result => {
        todoToUpdate.completed = result.data.completed
      }, error => {
        console.error(error);
      });
    },
  }
}
</script>
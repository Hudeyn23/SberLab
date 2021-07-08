<template>
  <li>
    <input type="checkbox" v-bind:checked="todo.completed" v-on:change="$emit('update_status', todo)"/>
    {{ todo.title }}
    <button v-on:click="$emit('remove', todo)">X</button>
  </li>
</template>

<script>
export default {
  props: {
    todo: {
      type: Object,
      required: true
    }
  },
}
</script>
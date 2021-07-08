<template>
  <div id="button">
    <p>accessKey<input v-model="accessKey" placeholder="accessKey"/></p>
    <p>secretKey<input v-model="secretKey" placeholder="secretKey"/></p>
    <p>projID<input v-model="projID" placeholder="projID"/></p>
    <p>offset<input v-model="offset" placeholder="offset"/></p>
    <p>limit<input v-model="limit" placeholder="limit"/></p>
    <button v-on:click="showInfo">Query</button>
    {{ info }}
  </div>
</template>

<script>
//import TodoListItem from './TodoListItem.vue'
import axios from "axios"

const axios_instance = axios.create({
  baseURL: 'http://37.230.196.108:9999/v1/',
  //baseURL: process.env.BACKEND,
});

export default {
  data() {
    return {
      projID: '',
      offset: '',
      limit: '',
      accessKey: '',
      secretKey: '',
      info: '',
    }
  },
  methods: {
    showInfo() {
      console.log(process.env.BACKEND)
      axios_instance.get(this.projID + "/cloudservers/detail?offset=" + this.offset + "&limit=" + this.limit + "&aKey=" + this.accessKey + "&sKey=" + this.secretKey).then(result => {
        this.info = result.data
        console.log(result)
      }, error => {
        console.error(error);
      });
    },
  }
}
</script>
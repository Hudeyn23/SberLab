<template>
    <div id="button">
      <p><input v-model="accessKey" placeholder="accessKey"/></p>
      <p><input v-model="secretKey" placeholder="secretKey"/></p>
      <p><input v-model="projID" placeholder="projID"/></p>
      <p><input v-model="offset" placeholder="offset"/></p>
      <p><input v-model="limit" placeholder="limit"/></p>
      <button v-on:click="showInfo">Query</button>
      <h1 v-if="jsonInfo.error_msg">{{ jsonInfo.error_msg }}</h1>
      <h1 v-if="jsonInfo.count">Total ECS count: {{ jsonInfo.count }}</h1>
      <h1>{{ jsonInfo }}</h1>
    </div>
</template>

<script>
//import TodoListItem from './TodoListItem.vue'
import axios from "axios"

const axios_instance = axios.create({
  baseURL: process.env.VUE_APP_BACKEND_IP,
  //baseURL: process.env.BACKEND,
  // baseURL: 'http://localhost:9999/v1',
  // baseURL: 'http://37.230.196.108/v1/'
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
      jsonInfo: {},
    }
  },
  methods: {
    showInfo() {
      console.log(process.env.VUE_APP_BACKEND_IP)
      axios_instance.get(this.projID + "/cloudservers/detail?offset=" + this.offset + "&limit=" + this.limit + "&aKey=" + this.accessKey + "&sKey=" + this.secretKey, {
        //timeout: 500,
      }).
      then(result => {
        this.info = result.data
        this.jsonInfo = JSON.parse(result.data)
        console.log(result)
      }, error => {
        console.error(error);
      });
    },
  }
}
</script>

<style>

/**, *::before, *::after {
  box-sizing: border-box;
}*/

@import url('https://fonts.googleapis.com/css?family=Proza+Libre|Fira+Mono');

html, body {
  height: 100%;
  margin: 0;
  padding: 0;
}

#button {
  display: flex;
  height: 100%;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}

h1 {
  font-family: 'Proza Libre', sans-serif;
  color: #ffffff;
  font-weight: 30;
}


</style>
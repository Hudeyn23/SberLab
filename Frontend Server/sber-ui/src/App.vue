<template>
  <div id="app" v-bind:style="{ backgroundColor: color}">
    <div class="color">
      <h1>Pick a color<input type="color" v-model="color"/></h1>
    </div>
    <h1>Get ECS list</h1>
    <SberButton/>
  </div>
</template>

<script>
import SberButton from './components/SberButton.vue'

export default {
  components: {
    SberButton
  },
  data() {
    return {color: '#666666'}
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

#app {
  display: flex;
  height: 100%;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}

input {
  padding: 1rem 1rem;
  font-size: 1rem;
  border: 0;
  border-radius: 90px;
  background: #fff;
}


h1 {
  font-family: 'Proza Libre', sans-serif;
  color: #ffffff;
  font-weight: 300;
}

input[type="color"] {
  width: 2rem;
  height: 2rem;
  border: none;
  outline: none;
  border-radius: 100%;
  margin: 50px;
  padding: 0px !important;
}


</style>
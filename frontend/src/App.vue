<template>
  <p>{{ text }}</p>
  <button @click="getTime">点击</button>
  <button @click="save">save</button>
</template>

<script setup>
import {Gopher, Save} from "../wailsjs/go/main/App.js";
import {ref} from "vue";

const text = ref("")

function getTime() {
  Gopher("我是js，呼叫gopher").then((res) => {
    text.value = res
  })
}

function save() {
  Save().then((res) => {
    text.value = res
  })
}

</script>

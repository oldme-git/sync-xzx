<template>
  <p>{{ text }}</p>
  <button @click="getTime">点击</button>
</template>

<script setup>
import { Gopher } from "../wailsjs/go/main/App.js";
import {ref} from "vue";
const text = ref("")

function getTime() {
  Gopher("我是js，呼叫gopher").then((res) => {
    text.value = res
  })
}

</script>

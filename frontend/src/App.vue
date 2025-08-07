<script lang="ts" setup>
import { onMounted } from 'vue'
import BottomBar from './components/BottomBar.vue'
import MemeBox from './components/meme/MemeBox.vue'
import { initializeStoreFromCache } from './store'
// import MeSetting from './components/Setting.vue';
// const closeSetting = () => {
//     console.log("call")
// }
onMounted(() => {
  initializeStoreFromCache()
})
</script>

<template>
  <MemeBox/>
  <BottomBar/>
</template>

<style>

</style>

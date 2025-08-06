<script lang="ts" setup>
import { onMounted } from 'vue';
import BottomBar from './components/BottomBar.vue';
import MemeBox from './components/MemeBox.vue';
import { initializeStoreFromCache } from './store';
// import MeSetting from './components/Setting.vue';
// const closeSetting = () => {
//     console.log("call")
// }
onMounted(() => {
  initializeStoreFromCache()
})
</script>

<template>
  <div>

    <MemeBox style="height: 90vh;"/>
    <BottomBar style="height: 10vh;" />
  </div>
</template>

<style>

</style>

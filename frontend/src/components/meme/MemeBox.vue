<script lang="ts" setup>
import { store } from '../../store'
import MemePane from './MemePane.vue'

</script>

<template>
  <main class="meme-box">
    <template v-for="memes in store.allMemesPath" :key="memes.Code">
      <MemePane :memeInfo="memes"/>
    </template>
  </main>
</template>

<style scoped>
.meme-box {
  width: 100%;
  height: 100%;
}
</style>

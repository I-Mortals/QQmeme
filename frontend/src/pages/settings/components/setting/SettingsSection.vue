<template>
  <div class="content-section">
    <h3 class="section-title">{{ title }}</h3>
    <div class="settings-container">
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  title: string
}

const props = defineProps<Props>()
</script>

<style scoped lang="less">
@import '@/styles/variables.less';

.content-section {
  margin-bottom: 1.5rem;

  &:last-child {
    margin-bottom: 0;
  }
}

.section-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: @rgb-bc;
  margin: 0 0 1rem 0;
  padding-bottom: 0.5rem;
  border-bottom: 1px solid @rgb-b3;
}

.settings-container {
  background: @rgb-b1;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}
</style>



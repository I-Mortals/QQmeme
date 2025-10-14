<template>
  <div class="setting-group">
    <div class="group-header">
      <div class="group-title">
        <slot name="title" />
      </div>
      <div v-if="$slots.desc" class="group-desc">
        <slot name="desc" />
      </div>
    </div>
    <div class="group-content">
      <slot />
    </div>
  </div>

</template>

<script setup lang="ts">
</script>

<style scoped lang="less">
@import '@/styles/variables.less';

.setting-group {
  background: @rgb-b1;
}

.group-header {
  display: flex;
  flex-direction: column;
  gap: 0.15rem;
  padding: 0 0.25rem;
  margin-bottom: .25rem;
}

.group-title {
  color: @rgb-bc;
}

.group-desc {
  font-size: 0.75rem;
  color: @rgb-bc;
  opacity: 0.7;
}

.group-content {
  border: 1px solid @rgb-b3;
  border-radius: 0.5rem;
  overflow: hidden;
}
</style>

<template>
  <div class="setting-item">
    <div class="setting-label">
      <span class="label-text">
        <slot name="text" />
      </span>
      <span v-if="$slots.desc" class="label-desc">
        <slot name="desc" />
      </span>
    </div>
    <div class="setting-control">
      <slot name="actions" />
    </div>
  </div>
  
</template>

<style scoped lang="less">
@import '@/styles/variables.less';

.setting-item {
  display: flex;
  align-items: center;
  padding: 0.5rem;
  border-bottom: 1px solid @rgb-b3;

  &:last-child {
    border-bottom: none;
  }
}

.setting-label {
  flex: 1;
  min-width: 0;
}

.label-text {
  display: block;
  font-size: 0.875rem;
  font-weight: 500;
  color: @rgb-bc;
}

.label-desc {
  display: block;
  font-size: 0.75rem;
  color: @rgb-bc;
  opacity: 0.7;
  line-height: 1.4;
}

.setting-control {
  flex-shrink: 0;
  margin-left: 1rem;
  display: flex;
  align-items: center;
}
</style>



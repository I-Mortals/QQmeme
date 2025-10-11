<script lang="ts" setup>
import ThemeSwitch from '@/components/common/ThemeSwitch.vue';

</script>

<template>
  <div id="appearance-settings" class="content-section">
    <h3 class="section-title">外观设置</h3>
    <div class="setting-group">
      <div class="setting-item">
        <div class="setting-label">
          <span class="label-text">主题方案</span>
          <span class="label-desc">选择应用的主题配色方案</span>
        </div>
        <div class="setting-control theme-control">
          <ThemeSwitch />
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>
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

.setting-group {
  background: @rgb-b1;
  border-radius: 0.5rem;
  border: 1px solid @rgb-b3;
  overflow: hidden;
}

.setting-item {
  display: flex;
  align-items: center;
  padding: 0.875rem 1.25rem;
  border-bottom: 1px solid @rgb-b3;
  position: relative;

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
  margin-bottom: 0.25rem;
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
  gap: .5rem;
}

.color-trigger {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.color-preview-wrapper {
  position: relative;
  width: 1.5rem;
  height: 1.5rem;
  border-radius: 0.25rem;
  border: 1px solid rgba(0, 0, 0, 0.2);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  background-image: url("data:image/svg+xml,%3csvg width='100' height='100' xmlns='http://www.w3.org/2000/svg'%3e%3cdefs%3e%3cpattern id='a' patternUnits='userSpaceOnUse' width='10' height='10'%3e%3crect width='10' height='10' fill='%23fff'/%3e%3crect width='5' height='5' fill='%23ccc'/%3e%3crect x='5' y='5' width='5' height='5' fill='%23ccc'/%3e%3c/pattern%3e%3c/defs%3e%3crect width='100' height='100' fill='url(%23a)'/%3e%3c/svg%3e");
  overflow: hidden;
  cursor: pointer;
  transition: all 0.2s ease;

  &:hover {
    transform: scale(1.05);
  }
}

.color-preview {
  width: 100%;
  height: 100%;
  border-radius: 0.25rem;
  transition: transform 0.1s ease;
}

.color-hex {
  font-family: 'SF Mono', 'Monaco', 'Inconsolata', 'Roboto Mono', monospace;
  font-size: 0.875rem;
  font-weight: 500;
  color: @rgb-bc;
  background: @rgb-b1;
  padding: 0.375rem 0.625rem;
  border-radius: 0.375rem;
  border: 1px solid @rgb-b3;
  min-width: 4.5rem;
  text-align: center;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}
</style>

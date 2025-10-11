<script setup lang="ts">
import { ref, watch } from 'vue'
import Modal from '@/components/Modal.vue'

interface DeleteMemeModalProps {
  visible: boolean
  fileName?: string
  loading?: boolean
}

const props = withDefaults(defineProps<DeleteMemeModalProps>(), {
  loading: false
})

const emit = defineEmits<{
  close: []
  'update:visible': [value: boolean]
  confirm: []
}>()

const isDeleting = ref(false)

watch(() => props.visible, (visible) => {
  if (visible) {
    isDeleting.value = false
  }
})

const hideModal = () => {
  if (!isDeleting.value && !props.loading) {
    emit('close')
    emit('update:visible', false)
  }
}

const handleConfirm = () => {
  isDeleting.value = true
  emit('confirm')
}
</script>

<template>
  <Modal
    :visible="props.visible"
    :closable="!isDeleting && !props.loading"
    :mask-closable="!isDeleting && !props.loading"
    :cancelBtn="{
      text: '取消',
      disabled: isDeleting || props.loading
    }"
    :confirmBtn="{
      text: '确定',
      loading: isDeleting || props.loading,
      disabled: isDeleting || props.loading
    }"
    @close="hideModal"
    @cancel="hideModal"
    @confirm="handleConfirm"
  >
    <div class="delete-confirm-content">
      <div class="message">
        <p>确定要删除 <strong>{{ props.fileName || '这个表情包' }}</strong> 吗？</p>
      </div>
    </div>

  </Modal>
</template>

<style lang="less" scoped>
@import '@/styles/variables.less';

.delete-confirm-content {
  padding: 1rem 0;
  text-align: center;
}

.message {
  p {
    font-size: 1rem;
    color: @rgb-bc;
    margin: 0;

    strong {
      color: @rgb-e;
      font-weight: 600;
    }
  }
}

</style>

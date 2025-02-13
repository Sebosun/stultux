<script setup lang="ts">
defineProps<{
  label: string
  name: string
}>()
</script>

<template>
  <div
    class="base-label"
  >
    <label
      class="base-label__label"
      :for="name"
    >
      {{ label }}
    </label>
    <slot />
  </div>
</template>

<style scoped lang="scss">
.base-label {
  &__label {
    font-size: 1rem;
    font-weight: 500;
    margin-bottom: 0.5rem;
  }
}
</style>

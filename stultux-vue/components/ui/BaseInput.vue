<script setup lang="ts">
const input = defineModel<string>('input')
defineProps<{
  placeholder: string
  name: string
  label: string
}>()
</script>

<template>
  <div class="flex flex-col gap-2">
    <label :for="name">
      {{ label }}
    </label>

    <input
      v-model="input"
      :name="name"
      class="rounded-xs border border-black bg-inherit px-4 py-2 outline-none transition-all focus-within:rounded-sm focus-within:ring focus-within:ring-[#f85552] dark:border-white"
      type="text"
      :placeholder="placeholder"
    >
  </div>
</template>

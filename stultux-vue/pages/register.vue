<script setup lang="ts">
import type { SelectValue } from '~/components/UI/BaseSelect.vue'

definePageMeta({
  layout: 'centered-layout',
})

interface User {

}

const arr = ref<Array<SelectValue<string>>>([
  {
    label: 'Naruto',
    value: 'Naruto',
  },
  {
    label: 'Saske',
    value: 'Saske',
  },
])

const BASE_URL = 'http://localhost:1323/api/v1'
onMounted(async () => {
  const result = await $fetch(`${BASE_URL}/registration`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  })
  console.log(result)
})
</script>

<template>
  <div class="bg-[#f2ebdc] dark:bg-zinc-950 dark:text-white">
    <UIBaseCard
      color="red"
      class="py-40"
    >
      <div class="mx-auto flex max-w-[300px] flex-col justify-center gap-4">
        <UIBaseSelect :options="arr" />
        dupa
      </div>
    </UIBaseCard>
  </div>
</template>

<script setup lang="ts">
import type { SelectValue } from '~/components/UI/BaseSelect.vue'

definePageMeta({
  layout: 'centered-layout',
})

export interface Result {
  firstNames: string[]
  lastNames: string[]
  passwords: string[]
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

const vals = ref<Result>()

const passwrdOptions = computed(() => {
  if (!vals.value) return []
  return vals.value?.passwords.map((item) => {
    return { label: item, value: item }
  })
})

const nameOptions = computed(() => {
  if (!vals.value) return []
  return vals.value?.firstNames.map((item) => {
    return { label: item, value: item }
  })
})

const lastNameOptions = computed(() => {
  if (!vals.value) return []
  return vals.value?.lastNames.map((item) => {
    return { label: item, value: item }
  })
})

const BASE_URL = 'http://localhost:1323/api/v1'
onMounted(async () => {
  const result = await $fetch<Result>(`${BASE_URL}/registration`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  })
  vals.value = result
})

const onSubmit = () => {
  const payload = {
    first_name: vals.value?.firstNames[0],
    last_name: vals.value?.lastNames[0],
    password: vals.value?.passwords[0],
  }

  const result = $fetch(`${BASE_URL}/user`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(payload),
  })
  console.log('submit')
}
</script>

<template>
  <div class="bg-[#f2ebdc] dark:bg-zinc-950 dark:text-white">
    <div class="py-40">
      <form
        class="mx-auto flex max-w-[300px] flex-col justify-center gap-4"
        @submit.prevent="onSubmit"
      >
        <UIBaseSelect
          name="First name"
          label="First name"
          :options="nameOptions"
        />
        <UIBaseSelect
          name="Last Name"
          label="First name"
          :options="lastNameOptions"
        />
        <UIBaseSelect
          name="Password"
          label="First name"
          :options="passwrdOptions"
        />
        <UIBaseButton class="w-full">
          Register
        </UIBaseButton>
      </form>
    </div>
  </div>
</template>

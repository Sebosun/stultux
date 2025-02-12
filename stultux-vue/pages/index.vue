<script setup lang="ts">
const login = ref('dupa')
const password = ref('dupa')
const router = useRouter()

definePageMeta({
  layout: 'centered-layout',
})

const onSubmit = (submitCase: 'login' | 'register') => {
  if (submitCase === 'login') {
    return
  }

  if (submitCase === 'register') {
    router.push('/register')
  }
}
</script>

<template>
  <div class="bg-[#f2ebdc] dark:bg-zinc-950 dark:text-white">
    <UIBaseCard
      color="red"
      class="py-40"
    >
      <div class="mx-auto flex max-w-[300px] flex-col justify-center gap-4">
        <div>
          <UIBaseInput
            v-model="login"
            label="Email"
            name="email"
            placeholder="Email"
          />
        </div>

        <div>
          <UIBaseInput
            v-model="password"
            label="Password"
            name="password"
            placeholder="Password"
          />
        </div>

        <div>
          <UIBaseButton
            class="mb-4 w-full"
            @click="onSubmit('login')"
          >
            Login
          </UIBaseButton>

          <UIBaseButton
            class="w-full"
            @click="onSubmit('register')"
          >
            Create Account
          </UIBaseButton>
        </div>
      </div>
    </UIBaseCard>
  </div>
</template>

<script setup lang="ts">
import BaseCard from '@/components/ui/BaseCard.vue'
import BaseButton from '@/components/ui/BaseButton.vue'
import BaseInput from '@/components/ui/BaseInput.vue'

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
    <BaseCard
      color="red"
      class="py-40"
    >
      <div class="mx-auto flex max-w-[300px] flex-col justify-center gap-4">
        <div>
          <BaseInput
            v-model="login"
            label="Email"
            name="email"
            placeholder="Email"
          />
        </div>

        <div>
          <BaseInput
            v-model="password"
            label="Password"
            name="password"
            placeholder="Password"
          />
        </div>

        <div>
          <BaseButton
            class="mb-4 w-full"
            @click="onSubmit('login')"
          >
            Login
          </BaseButton>

          <BaseButton
            class="w-full"
            @click="onSubmit('register')"
          >
            Create Account
          </BaseButton>
        </div>
      </div>
    </BaseCard>
  </div>
</template>

<script setup lang="ts">
const route = useRoute()

useHead({
  meta: [{ property: 'og:title', content: `${route.meta.title}` }],
})
</script>

<template>
  <div
    class="bg-transition min-h-screen bg-[#f2ebdc] dark:bg-zinc-950 dark:text-white relative"
  >
    <h1 class="absolute text-4xl font-bold left-4 top-4">
      Stultux
    </h1>
    <header class="absolute right-0">
      <div class="pt-4 flex justify-end gap-4">
        <UIBaseButton @click="$colorMode.preference = 'light'">
          Light
        </UIBaseButton>

        <UIBaseButton @click="$colorMode.preference = 'dark'">
          Dark
        </UIBaseButton>

        <button />
      </div>
    </header>
    <main
      class="bg-p-surface-200 flex min-h-screen max-w-full items-center justify-center"
    >
      <div class="m-auto min-w-[500px] max-w-xl">
        <slot />
      </div>
    </main>
  </div>
</template>

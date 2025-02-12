<script setup lang="ts">
type Props = {
  type?: 'reactive' | 'normal'
}

withDefaults(defineProps<Props>(), {
  type: 'normal',
})
</script>

<template>
  <button
    class="gumroadish-button--normal box-shadow-hover rounded-xs flex items-center justify-center border border-solid border-black px-4 py-3 text-center tracking-wide dark:border-white"
  >
    <slot />
  </button>
</template>

<style scoped langs="scss">
.gumroadish-button {
    &--normal {
        &:hover {
            transition: all;
            transition-duration: 0.2s;
            transform: translate(-0.25rem, -0.25rem);
            box-shadow: 0.25rem 0.25rem 0rem #f85552;
        }
    }
}
</style>

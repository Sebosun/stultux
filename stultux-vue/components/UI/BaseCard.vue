<script setup lang="ts">
const { color } = defineProps<{
  color?: 'black' | 'brand' | 'red'
}>()
</script>

<template>
  <div
    class="gumroadish-card rounded-xs border border-black p-2 dark:border-white"
    :class="{
      'gumroadish-card--black': color === 'black',
      'gumroadish-card--brand': color === 'brand',
      'gumroadish-card--red': color === 'red',
    }"
  >
    <slot />
  </div>
</template>

<style scoped lang="scss">
.gumroadish-card {
    &:hover {
        transition: all;
        transition-duration: 0.2s;
        /* transform: translate(-0.25rem, -0.25rem); */
        box-shadow: 0.25rem 0.25rem 0rem #f85552;
    }

    .light-mode &--black:hover {
        box-shadow: 0.25rem 0.25rem 0rem black;
    }

    .dark-mode &--black:hover {
        box-shadow: 0.25rem 0.25rem 0rem white;
    }

    &--brand:hover {
        box-shadow: 0.25rem 0.25rem 0rem #f85552;
    }

    &--red:hover {
        box-shadow: 0.25rem 0.25rem 0rem #ce141e;
    }
}
</style>

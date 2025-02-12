<script setup lang="ts">
export type SelectValue<T> = {
  value: T
  label: string
}

const props = defineProps<{
  icon?: string
  options: SelectValue<any>[] // eslint-ignore-line
}>()

const value = defineModel<SelectValue<any>>()

onBeforeMount(() => {
  value.value = props.options[0]
})

const showDropdown = ref<boolean>(false)

const onOptClick = (val: SelectValue<any>) => {
  value.value = val
  showDropdown.value = false
}
</script>

<template>
  <div class="base_select">
    <button
      class="w-full box-shadow-hover rounded-xs flex items-center justify-center border border-solid border-black   tracking-wide dark:border-white"
      role="listbox"
      @click="showDropdown = !showDropdown"
    >
      {{ value?.value }}
    </button>
    <ul
      v-show="showDropdown"
      class="base_select__list"
    >
      <li
        v-for="opt in options"
        :key="opt.label"
        role="option"
        class="base_select__item"
        @click="onOptClick(opt)"
      >
        {{ opt.label }}
      </li>
    </ul>
  </div>
</template>

<style scoped lang="scss">
.base_select {
    position: relative;
    width: 100%;

    &__list {
        width: 100%;
        position: absolute;
        background: var(--secondary-bg);
    }

    &__item {
        padding: 2px 10px;

    }
}
</style>

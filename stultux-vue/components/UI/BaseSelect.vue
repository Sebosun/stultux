<script setup lang="ts">
import { onClickOutside } from '@vueuse/core'

export type SelectValue<T> = {
  value: T
  label: string
}

const props = defineProps<{
  icon?: string
  name: string
  label: string
  options: SelectValue<any>[] // eslint-ignore-line
}>()

const value = defineModel<SelectValue<any>>()

onBeforeMount(() => {
  value.value = props.options[0]
})

const showDropdown = ref<boolean>(false)

const hideDropdown = () => {
  console.log('hideDropdown')
  showDropdown.value = false
}

const onOptClick = (val: SelectValue<any>) => {
  value.value = val
  showDropdown.value = false
}

const labelRef = ref<HTMLElement | null>(null)

onClickOutside(labelRef, hideDropdown)
</script>

<template>
  <div class="base-select">
    <UIBaseLabel
      ref="labelRef"
      :label="label"
      :name="name"
    >
      <button
        class="base-select__button"
        role="listbox"
        @click.stop.prevent="showDropdown = !showDropdown"
      >
        <span>
          {{ value?.value ?? "  " }}
        </span>
        <Transition
          name="arrow-up"
          mode="out-in"
        >
          <Icon
            v-if="showDropdown"
            name="material-symbols:expand-less"
            class="text-white text-xl"
          />
          <Icon
            v-else
            name="material-symbols:expand-more"
            class="text-white text-xl"
          />
        </Transition>
      </button>
    </UIBaseLabel>
    <Transition name="slide-up">
      <ul
        v-show="showDropdown"
        class="base-select__list"
      >
        <li
          v-for="opt in options"
          :key="opt.label"
          role="option"
          class="base-select__item"
          :class="{ 'base-select__item--chosen': opt.value === value?.value }"
          @click="onOptClick(opt)"
        >
          {{ opt.label }}
        </li>
      </ul>
    </Transition>
  </div>
</template>

<style scoped lang="scss">
.base-select {
    position: relative;
    width: 100%;

    &__button {
        width: 100%;
        padding: 9px 8px;
        min-height: 48px;

        display: flex;
        justify-content: space-between;
        align-items: center;

        border: 1px solid var(--border);
        border-radius: 2px;

    }

    &__list {
        z-index: 1;
        max-height: 200px;
        overflow-y: auto;
        width: 100%;
        position: absolute;
        margin-top: 10px;
        border: 1px solid var(--border);
        box-shadow: 0.25rem 0.25rem 0rem #f85552;
    }

    &__item {
        cursor: pointer;
        padding: 9px 8px;

        background: var(--primary-bg);
        border-radius: 2px;

        &--chosen {
            background: var(--brand);
        }

        &:hover {
            background: var(--brand);
        }
    }
}

.slide-up-enter-active,
.slide-up-leave-active {
    transition: all 0.50s ease-out;
}

.slide-up-enter-from {
    width: 20%;
    opacity: 0;
    transform: translateY(-30px);
    box-shadow: 0.25rem 0.25rem 0.55rem #f85552;
    filter: blur(5px);
}

.slide-up-leave-to {
    opacity: 0;
    width: 20%;
    transform: translateY(30px);
    box-shadow: 0.25rem 0.25rem 1rem #f85552;
    filter: blur(5px);
}

.arrow-up-enter-active,
.arrow-up-leave-active {
    transition: all 0.20s ease-out;
}

.arrow-up-enter-from,
.arrow-up-leave-to {
    transform: translateY(5px);
    opacity: 0;
}
</style>

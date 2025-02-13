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
      class="base_select__button"
      role="listbox"
      @click="showDropdown = !showDropdown"
    >
      <span>
        {{ value?.value }}
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
    <Transition name="slide-up">
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
    </Transition>
  </div>
</template>

<style scoped lang="scss">
.base_select {
    position: relative;
    width: 100%;

    &__button {
        width: 100%;
        padding: 9px 8px;

        display: flex;
        justify-content: space-between;

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

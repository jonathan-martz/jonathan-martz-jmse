<template>
  <section class="hidden md:flex lg:flex xl:hidden mx-auto max-w-6xl mb-3 md:justify-between lg:justify-between">
    <section class="logo">
      <img src="https://cdn.jmse.cloud/jmse-logo.svg" alt="" class="h-24 md:h-32 md:px-3 md:py-3 bg-white rounded-xl">
    </section>
    <section class="navigation mt-3">
      <section v-if="!open" class="bars">
        <button class="" @click="open = !open">
          <FontAwesomeIcon :icon="faBars" class="mt-3 mr-6" size="5x" />
        </button>
      </section>
      <section v-else class="times">
        <button class="" @click="open = !open">
          <FontAwesomeIcon :icon="faTimes" class="mt-3 mr-6" size="5x" />
        </button>
      </section>
    </section>
  </section>
</template>

<script lang="ts" setup>
import { faBars, faTimes } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { useLocalStorage } from '@vueuse/core';

const open = useLocalStorage('open', false, {});
</script>

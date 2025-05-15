<template>
  <div v-if="open" class="mx-auto max-w-6xl mb-3 bg-white px-3 py-3 block md:hidden">
   <section class="grid grid-cols-12 gap-3">
    <div class="col-span-6 md:col-span-3">
      <a href="/" class="btn btn-primary btn-block">Startseite</a>
    </div>
    <div class="col-span-6 md:col-span-3">
      <a href="/" class="btn btn-primary btn-block">Services</a>
    </div>
    <div class="col-span-6 md:col-span-3">
      <a href="/" class="btn btn-primary btn-block">Portfolio</a>
    </div>
    <div class="col-span-6 md:col-span-3">
      <a href="/" class="btn btn-primary btn-block">Kontakt</a>
    </div>
   </section>
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';

const open = useLocalStorage('open', false, {});
</script>
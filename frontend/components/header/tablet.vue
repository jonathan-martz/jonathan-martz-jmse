<template>
  <div class="flex justify-between block xl:hidden">
    <div class="drawer block xl:hidden mb-3">
      <input id="my-drawer" type="checkbox" v-model="open" class="drawer-toggle" />
      <div class="drawer-content flex justify-between">
        <!-- Page content here -->
        <a href="/">
          <img src="https://cdn.jmse.cloud/jmse-logo.svg" class="h-26 rounded-xl bg-white px-3 py-3" alt="">
        </a>
        <label for="my-drawer" class="btn btn-primary btn-lg drawer-button mt-6 h-20">
          <FontAwesomeIcon :icon="faBars" color="white" class="" size="2x" />
        </label>
      </div>
      <div class="drawer-side">
        <label for="my-drawer" aria-label="close sidebar" class="drawer-overlay"></label>
        <ul class="menu bg-gray-500 text-white min-h-full w-80 p-4 space-y-3">
          <li class="">
            <a href="/">
              <img src="https://cdn.jmse.cloud/jmse-logo.svg" class="h-26 rounded-xl bg-white px-3 py-3" alt="">
            </a>
          </li>
          <!-- Sidebar content here -->
          <li><a href="/#services" @click="open = false" class="btn btn-primary">Services</a></li>
          <li><a href="/projekte" @click="open = false" class="btn btn-primary">Projekte</a></li>
          <li><a href="/kunden" @click="open = false" class="btn btn-primary">Kunden</a></li>
          <li><a href="/team" @click="open = false" class="btn btn-primary">Unser Team</a></li>
          <li><a href="/hardware" @click="open = false" class="btn btn-primary">Hardware</a></li>
          <li><a href="/impressum" @click="open = false" class="btn btn-primary">Impressum</a></li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { faBars } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { useLocalStorage } from '@vueuse/core';
import { onUnmounted } from 'vue';

const open = useLocalStorage('open', false, {});

onUnmounted(()=>{
  open.value = false
});
</script>
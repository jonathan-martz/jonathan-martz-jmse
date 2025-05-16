<template>
  <section class="page bg-black py-3 px-3">
    <HeaderDesktop />
    <HeaderTablet />
    <section>
      <FontAwesomeIcon :icon="faBars" />
    </section>
    <main class="max-w-6xl mx-auto px-3 py-3 bg-white min-h-screen text-black rounded-t-xl">
      <slot />
    </main>
    <footer class="max-w-6xl mx-auto md:px-3 md:py-3 bg-blue-400">
      Footer
    </footer>
  </section>
</template>

<style scoped>

.page {
  --s: 200px; /* control the size*/
  --c1: #1d1d1d;
  --c2: #4e4f51;
  --c3: #3c3c3c;
  
  background:
    repeating-conic-gradient(from 30deg,#0000 0 120deg,var(--c3) 0 50%) 
     calc(var(--s)/2) calc(var(--s)*tan(30deg)/2),
    repeating-conic-gradient(from 30deg,var(--c1) 0 60deg,var(--c2) 0 120deg,var(--c3) 0 50%);
  background-size: var(--s) calc(var(--s)*tan(30deg));
}</style>

<script lang="ts" setup>

</script>

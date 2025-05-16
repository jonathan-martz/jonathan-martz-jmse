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
    <footer class="footer sm:footer-horizontal max-w-6xl mx-auto px-6 py-6 bg-blue-400">
      <nav>
        <h6 class="footer-title text-lg">Seiten</h6>
        <a href="/" class="link link-hover">Startseite</a>
        <a href="/" class="link link-hover">Projekte</a>
        <a href="/" class="link link-hover">Kunden</a>
      </nav>
      <nav>
        <h6 class="footer-title text-lg">Services</h6>
        <a href="/" class="link link-hover">Webseiten Entwicklung</a>
        <a href="/" class="link link-hover">Online Shop Entwicklung</a>
        <a href="/" class="link link-hover">Web-Based Software</a>
      </nav>
      <nav>
        <h6 class="footer-title text-lg">Keywords</h6>
        <button type="button" class="link link-hover">Bergneustadt Webentwickler</button>
        <button type="button" class="link link-hover">Bergneustadt Software Entwickler</button>
        <button type="button" class="link link-hover">Softwareentwicklung Bergneustadt</button>
      </nav>
      <nav>
        <h6 class="footer-title text-lg">Development</h6>
        <h3 class="font-bold">Made with <FontAwesomeIcon :icon="faHeart" color="red" class="w-5 mx-1 mt-1"/></h3>
        <h3 class="font-bold">Copyright by Jonathan</h3>
        <h3 class="font-bold">Design by Jonathan</h3>
      </nav>
    </footer>
  </section>
</template>

<style scoped>
.page {
  --s: 200px;
  /* control the size*/
  --c1: #1d1d1d;
  --c2: #4e4f51;
  --c3: #3c3c3c;

  background:
    repeating-conic-gradient(from 30deg, #0000 0 120deg, var(--c3) 0 50%) calc(var(--s)/2) calc(var(--s)*tan(30deg)/2),
    repeating-conic-gradient(from 30deg, var(--c1) 0 60deg, var(--c2) 0 120deg, var(--c3) 0 50%);
  background-size: var(--s) calc(var(--s)*tan(30deg));
}
</style>

<script lang="ts" setup>
import { faHeart } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';


</script>

<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>


<script setup lang="ts">
onMounted(() => {
  useHead({
    link: [
      { rel: 'manifest', href: '/manifest.webmanifest' },
      {
        rel: 'icon',
        type: 'image/png',
        href: 'https://cdn.jmse.cloud/avatar-2024-modified.png'
      }
    ]
  });
});
</script>
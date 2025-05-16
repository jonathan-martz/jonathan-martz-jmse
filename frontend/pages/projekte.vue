<template>
  <div class="grid grid-cols-6 gap-3">
    <div class="col-span-6">
      <h2 class="font-bold text-lg block text-center">Meine Projekte</h2>
      <p class="block text-center text-sm max-w-xl mx-auto">
        Ich bin private wie geschäftlich immer an kleinere Projekten dran dies ist eine kleine Auswahl davon. Damit
        verfolge ich nicht immer kommerzielle Interessen, den ich bin Webentwickler mit Leidenschaft.
      </p>
    </div>
    <div class="col-span-6 md:col-span-2">
      <ProjektePocketbasephpsdk />
    </div>
    <div class="col-span-6 md:col-span-2">
      <ProjektePocketstore />
    </div>
    <div class="col-span-6 md:col-span-2">
      <ProjekteHetznerapp />
    </div>
    <div class="col-span-6 md:col-span-2">
      <ProjekteElektroJonny />
    </div>
    <div class="col-span-6 md:col-span-2">
      <ProjekteGummersbachBikerShop />
    </div>
    <div class="col-span-6 md:col-span-2">
      <Projekte3dWorldPro />
    </div>

  </div>
</template>

<script lang="ts" setup>
onMounted(() => {
  useHead({
    title: 'JMSE - Projekte',
    meta: [
      {
        name: 'desc',
        content: 'Ich bin private wie geschäftlich immer an kleinere Projekten dran dies ist eine kleine Auswahl davon.' +
          'Damit verfolge ich nicht immer kommerzielle Interessen, den ich bin Webentwickler mit Leidenschaft.'
      }
    ]
  });
});
</script>

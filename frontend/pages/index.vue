<template>
    <div class="hero py-12">
        <div class="hero-content text-center">
            <div class="max-w-lg">
                <img src="https://cdn.jmse.cloud/avatar-2024-modified.png"
                    class="w-32 mx-auto border-3 border-black rounded-full mb-6" alt="">
                <h1 class="text-3xl font-bold">Jonathan Martz</h1>
                <p class="py-6 text-sm">
                    Ich bin ein Webentwickler aus Leidenschaft seit 2012 und entwickele moderne Web-Apps für
                    Unternehmen aus dem Oberbergischen seit 2023. Vorher habe ich Erfahrung im E-Commerce gemacht.
                </p>
                <section class="actions space-x-5">
                    <a href="tel:+491745141698" class="btn btn-primary">Jetzt anrufen</a>
                    <a href="mailto:support@jonathan-martz.de" class="btn btn-secondary mt-2 btn-sm">Mail schreiben</a>
                </section>
            </div>
        </div>
    </div>
    <section class="bg-black pb-12">
        <section class="h-[24rem] w-full bg-cover bg-center"
            style="background-image: url('/unsplash/cherrydeck-rMILC1PIwM0-unsplash.jpg');">
        </section>
        <section class="mx-6 text-black -mt-24 min-h-24">
            <IndexServicesOverview />
        </section>
    </section>
</template>

<script setup lang="ts">
import { useHead } from '@vueuse/head';
import { onMounted } from 'vue';

onMounted(() => {
    useHead({
        title: 'JMSE - Jonathan Martz Software Entwicklung',
        meta: [{
            name: 'description',
            content: 'Wir entwickeln Webseite und Online-Shops für Unternehmen und Händler aus dem Oberbergischen Kreis seit 2023.'
        }]
    });
});
</script>

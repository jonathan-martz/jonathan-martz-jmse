<script setup lang="ts">
import { faGithub } from '@fortawesome/free-brands-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

</script>

<template>
    <div class="card bg-gray-400 shadow-sm">
        <figure>
            <img src="/We-are-Hetzner_thmb_small-v2.webp" alt="Shoes" />
        </figure>
        <div class="card-body">
            <h2 class="card-title">Hetzner.app</h2>
            <section class="alert alert-warning block py-2 text-center">Unabhägig von Hetzner Online GmbH</section>
            <p>
                Eine inoffizielle App zur Benutzung der Hetzner Cloud Api. Als Kunde von Hetzner mit einer genialen
                API. Habe ich schon mit Vue und auch schonmal mit Svelte entwickelt.
            </p>
            <div class="card-actions justify-end">
                <a href="https://www.Hetzner.app?ref=jmse" target="_blank" class="btn btn-primary">Live</a>
                <a href="https://github.com/HetznerApp/nuxt-hetzner-app" target="_blank" class="btn btn-secondary">
                    <FontAwesomeIcon :icon="faGithub" />
                    <span>Github</span>
                </a>
            </div>
        </div>
    </div>
</template>
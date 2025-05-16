<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faGithub } from '@fortawesome/free-brands-svg-icons'
import { faLink } from '@fortawesome/free-solid-svg-icons';
</script>

<template>
    <div class="card bg-gray-400 shadow-sm">
        <figure>
            <img src="https://www.pocketstore.io/pocketstore-logo.svg" class="bg-white py-24" alt="Shoes" />
        </figure>
        <div class="card-body">
            <h2 class="card-title">PocketStore.io</h2>
            <p>Pocketstore ist ein Projekte was eine
                <a href="https://pocketbase.io" class="text-blue-600" target="_blank" rel="noopener noreferrer">
                    Pocketbase
                </a> Instanz einsetzt um einen kleinen Online Store bereit zu
                stellen mit einem Docker Compose Container.
            </p>
            <div class="card-actions justify-end">
                <a href="https://pocketstore.io" target="_blank" class="btn btn-primary">
                    <FontAwesomeIcon :icon="faGithub" />
                    <span>Github</span>
                </a>
                <a href="https://demo.pocketstore.io?ref=jmse" target="_blank" class="btn btn-secondary">
                    <FontAwesomeIcon :icon="faLink" />
                    <span>Demo</span>
                </a>
            </div>
        </div>
    </div>
</template>
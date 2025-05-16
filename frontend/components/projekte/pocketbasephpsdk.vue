<script setup lang="ts">
import { faGithub } from '@fortawesome/free-brands-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

</script>

<template>
    <div class="card bg-gray-400 shadow-sm">
        <figure>
            <img src="/unsplash/ben-griffiths-Bj6ENZDMSDY-unsplash.jpg" alt="Shoes" />
        </figure>
        <div class="card-body">
            <h2 class="card-title">Pocketbase - PHP SDK</h2>
            <p>
                Ich bin leidenschaftlicher PHP Entwickler und <a href="https://pocketbase.io" class="text-blue-600"
                    target="_blank" rel="noopener noreferrer">
                    Pocketbase
                </a> setze ich seit 2023 ein für alle möglichen
                Projekte unter anderem für PocketStore.io
            </p>
            <div class="card-actions justify-end">
                <a href="https://github.com/jonathan-martz/pocketbase-php-sdk" target="_blank"
                    class="btn btn-secondary">
                    <FontAwesomeIcon :icon="faGithub" />
                    <span>Github</span>
                </a>
            </div>
        </div>
    </div>
</template>
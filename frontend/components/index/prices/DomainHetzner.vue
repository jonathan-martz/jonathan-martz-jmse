<template>
  <div class="card bg-gray-400 shadow-sm">
    <div class="card-body">
      <span class="badge badge-sm badge-warning">Hetzner Online GmbH</span>
      <div class="flex justify-between">
        <h2 class="text-xl font-bold">Domain - Hetzner</h2>
        <span class="text-xl">?? €</span>
      </div>
      <p class="text-sm">
        Du willst die Wahl zwischen hunderten TLD's haben ? Dann schaue dir jetzt die Domains bei United Domains an.
      </p>
      <ul class="mt-6 flex flex-col gap-2 text-xs">
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>.de, .com, .net, .org, .at, .be, .eu, .name oder .nl</span>
        </li>
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>kostenloses SSL Cert by LetsEncrypt</span>
        </li>
      </ul>
      <div class="mt-6">
        <a href="https://www.hetzner.com/de/cloud?ref=OJpbQQDpJnE8" target="_blank"
          class="btn btn-primary btn-block">Jetzt buchen</a>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { faCheck } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

</script>
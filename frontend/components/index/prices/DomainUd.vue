<template>
  <div class="card bg-gray-400 shadow-sm">
    <div class="card-body">
      <span class="badge badge-sm badge-warning">united-domains AG</span>
      <div class="flex justify-between">
        <h2 class="text-xl font-bold">Domain - United Domains</h2>
        <span class="text-xl">?? €</span>
      </div>
      <p class="text-sm">
        Du willst die Wahl zwischen hunderten TLD's haben ? Dann schaue dir jetzt die Domains bei United Domains an.
      </p>
      <ul class="mt-6 flex flex-col gap-2 text-xs">
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>Jede Domain die das Herz begehrt</span>
        </li>
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>Mail Server(extra + 2€)</span>
        </li>
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>kostenloses SSL Cert by LetsEncrypt</span>
        </li>
      </ul>
      <div class="mt-6">
        <a href="https://www.united-domains.de/domain-kaufen/" target="_blank"
          class="btn btn-primary btn-block">Jetzt buchen</a>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { faCheck } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

</script>
<template>
  <div class="card bg-gray-400 shadow-sm">
    <div class="card-body">
      <span class="badge badge-sm badge-warning">Hetzner Online GmbH</span>
      <div class="flex justify-between">
        <h2 class="text-xl font-bold">Hetzner Cloud - CX22</h2>
        <span class="text-xl">5 €</span>
      </div>
      <p class="text-sm">Du willst einen Online Shop erstellen, dann reicht schon ein Cloud Server für den Anfang, den
        man beliebig erweitern kann mit steigenden Kunde zahlen.</p>
      <ul class="mt-6 flex flex-col gap-2 text-xs">
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>Eigene Adresse (extra)</span>
        </li>
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>Speicher: 40GB / RAM: 4GB / CPU's: 2</span>
        </li>
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>Daily Backup</span>
        </li>
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>Kostenlose SSL Cert by LetEncrypt</span>
        </li>
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>Eigene Adresse (extra)</span>
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
<template>
  <div class="card bg-gray-400 shadow-sm">
    <div class="card-body">
      <span class="badge badge-sm badge-warning">Ionos - eine 1&1 Firma</span>
      <div class="flex justify-between">
        <h2 class="text-xl font-bold">Ionos - vServer</h2>
        <span class="text-xl">1 €</span>
      </div>
      <p class="text-sm">
        Du willst eine Webseite mit einem Admin Interface erweitern ?
        
      </p>
      <ul class="mt-6 flex flex-col gap-2 text-xs">
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>1 CPU / 1GB RAM / 10GB NVMe</span>
        </li>
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>12 Monat Laufzeit</span>
        </li>
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>Pocketbase.io - Self Hosted einrichtbar</span>
        </li>
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>Eigene Adresse (extra)</span>
        </li>
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>Kostenlose SSL Cert by LetEncrypt</span>
        </li>
      </ul>
      <div class="mt-6">
        <a href="https://aklam.io/27RiRV" target="_blank" class="btn btn-primary btn-block">Jetzt buchen</a>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { faCheck, faEdit } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

</script>
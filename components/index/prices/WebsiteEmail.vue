<template>
  <div class="card bg-gray-400 shadow-sm">
    <div class="card-body">
      <span class="badge badge-sm badge-warning">Hetzner Online GmbH</span>
      <div class="flex justify-between">
        <h2 class="text-xl font-bold">Hetzner Webhosting</h2>
        <span class="text-xl">2 €</span>
      </div>
      <p class="text-sm">Du willst eine Webseite mit professionelle Adresse, sowie Email Adresse ? </p>
      <ul class="mt-6 flex flex-col gap-2 text-xs">
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>Eigene Adresse</span>
        </li>
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>Speicher: 10GB</span>
        </li>
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>Mailboxen: 100</span>
        </li>
        <li>
          <FontAwesomeIcon :icon="faCheck" color="green" class="mr-3" />
          <span>Kostenlose SSL Cert by LetEncrypt</span>
        </li>
      </ul>
      <div class="mt-6">
        <a href="https://www.hetzner.com/de/webhosting?ref=OJpbQQDpJnE8" target="_blank" class="btn btn-primary btn-block">Jetzt buchen</a>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { faCheck, faEdit } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

</script>
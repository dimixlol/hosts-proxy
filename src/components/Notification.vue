<template>
  <transition name="notification">
    <div v-if="store.notificationVisible">
      <div class="rounded-4 w-50 bg-danger overflow-y-scroll text-center p-2 pt-5">
        {{store.notificationMessage}}
      </div>
    </div>
  </transition>
</template>

<script setup lang="ts">
import {useMainStore} from "../store";
const store = useMainStore()
</script>

<style>
.notification-enter-active, .notification-leave-active { transition: all 1s cubic-bezier(.28,.68,0,1.52); }
.notification-enter-from, .notification-leave-to { transform: translateY(-50%); opacity: 0; }
</style>
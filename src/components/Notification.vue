<template>
  <transition name="notification">
    <div v-if="notificationVisible">
      <div class="rounded-4 w-50 bg-danger overflow-y-scroll text-center p-2 pt-5">
        {{notificationMessage}}
      </div>
    </div>
  </transition>
</template>

<style>
.notification-enter-active, .notification-leave-active { transition: all 1s cubic-bezier(.28,.68,0,1.52); }
.notification-enter-from, .notification-leave-to { transform: translateY(-50%); opacity: 0; }
</style>

<script setup lang="ts">

import {useStore} from "vuex";
import {computed} from "vue";

const store = useStore()

const notificationVisible = computed(() => store.getters.notificationVisible);
const notificationMessage = computed(() => store.getters.notificationMessage);
</script>
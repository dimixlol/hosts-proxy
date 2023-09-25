<template>
<div>
  <Form :formData="formData" @submit.prevent="createSite" class="form-group d-grid align-self-center w-50 app-form"/>
  <Egg v-model:eggVisible="eggToggle"/>
</div>
</template>

<script lang="ts" setup>
import Form from "../components/Form.vue";
import Egg from "../components/Egg.vue";

import VueCookies from "vue-cookies";
import {inject, ref, reactive} from "vue";
import {IpValidator, HostValidator} from "../validators";
import {useRouter} from "vue-router";
import {useMainStore} from "../store";

const eggToggle = ref(false);
const store = useMainStore()
const $cookies = inject<VueCookies>('$cookies');
const formData = reactive({
  host: {
    name: "host",
    placeholder: "domain-name.tld",
    required: true,
    validator: HostValidator
  },
  ip: {
    name: "ip",
    placeholder: "255.255.255.255",
    required: true,
    validator: IpValidator
  }
})
const router = useRouter()
const toggleEgg = () => {
  if (Math.floor(Math.random() * 100) == 50 || $cookies.get("testCookie") === store.testCookie) {
    eggToggle.value = !eggToggle.value;
    setTimeout(() => eggToggle.value = !eggToggle.value, 2000);
  }
}
const createSite = (e:any) => {
  toggleEgg();
  store.client.createSite(e.target.host.value, e.target.ip.value)
    .then((resp: any) => {
      store.setSlugData(resp.data.data)
      router.push({name: "siteView"})
    })
    .catch((err: any) => {
      store.toggleNotification("Something went wrong!")
  })
}
</script>

<style lang="scss">
.app-form {
  padding: 10em 1em 1em 1em;
  grid-template: 1fr / 1fr 1fr;
  grid-gap:  1em;
  :last-child {
    grid-column: 1/3;
    width: 75%;
    justify-self: center;
  }
}


</style>
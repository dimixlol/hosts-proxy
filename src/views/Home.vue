<template>
<div>
  <Form :formData="formData" @submit.prevent="createSite" class="form-group d-grid align-self-center w-50 app-form"/>
</div>
</template>

<script lang="ts" setup>
import Form from "../components/Form.vue";
import {reactive} from "vue";
import {IpValidator, HostValidator} from "../validators";
import {useRouter} from "vue-router";
import {useMainStore} from "../store";

const store = useMainStore()
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

const createSite = (e:any) => {
  store.client.createSite(e.target.host.value, e.target.ip.value)
    .then((resp: any) => {
      store.setSlugData(resp.data)
      router.push({name: "siteView"})
    })
    .catch((err: any) => {
      store.toggleNotification("Something went wrong!")
  })
}
</script>

<style lang="scss">
@import "bootstrap/scss/functions";
@import "bootstrap/scss/variables";
@import "bootstrap/scss/mixins";

.app-form {
  grid-template: 1fr / 1fr;
  grid-gap: 1em;
  padding: 10em 1em 1em 1em;

  @include media-breakpoint-up(md) {
    grid-template: 1fr / 1fr 1fr;
      :last-child {
      grid-column: 1/3;
      width: 75%;
      justify-self: center;
    }
  }
}
</style>
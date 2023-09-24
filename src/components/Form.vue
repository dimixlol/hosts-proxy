<template>
  <form @submit.prevent="submitForm">
    <input v-for="(props, item) in formData"
      :type="props.type||'text'"
      :name="item"
      :placeholder="props.placeholder||''"
      :required="props.required||false"
      class="form-control form-control-lg submit"
      @input="resetInvalid"
    />

    <input
        class="btn btn-outline-success opacity-75"
        type="submit"
        value="Create"
    />
  </form>
</template>

<script lang="ts" setup>
const emit = defineEmits(["submit"])
const props = defineProps({
  formData: {
    type: Object,
    required: true
  }
})
const resetInvalid = (e: any) => { e.target.classList.remove("is-invalid"); }
const setInvalid = (h: HTMLElement) => { h.classList.add("is-invalid"); }
const submitForm = (e:any) => {
  Object.keys(props.formData).every((item: string) => {
    const validator = props.formData[item].validator
    if (validator && !new validator(e.target[item].value).validate()) {
      console.log(`Validation failed for ${item}`)
      setInvalid(e.target[item])
      return false
    }
    console.log(`Validation passed for ${item}`)
    return true
  }) && emit("submit", e)
}
</script>

<style>
.form-control::placeholder { color: rgba(0,0,0,0.25)!important; }
.is-invalid { border: inset 1px red!important; }
</style>

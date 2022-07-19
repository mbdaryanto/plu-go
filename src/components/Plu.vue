<script setup lang="ts">
import { ref, reactive } from 'vue'
import { getItem, PluResponseType } from './item'
import axios from 'axios'
import LabelValue from './LabelValue.vue';

const code = ref("")
const plu = ref<PluResponseType>({
  item: {
    IDItem: 0,
    Barcode: '',
    Kode: '',
    Nama: '',
    HargaJual: 0.0,
    HargaNormal: 0.0,
    // JumlahDos: 0,
    KodePabrik: '',
    // Satuan: '',
    // Singkatan: '',
  },
  hargaPromo: [],
  hargaGrosir: [],
})

const nf = new Intl.NumberFormat("id-ID")
// defineComponent({
//   methods: {
//     handleSubmit(ev) {
//       console.log({target: ev.target})
//       const formData = new FormData(ev.target)
//     }
//   }
// })
function handleSubmit(ev: Event) {
  // console.log({target: ev.target, currentTarget: ev.currentTarget})
  const formData = new FormData(ev.target as HTMLFormElement)
  // console.log({code: formData.get('code')})
  const code = formData.get('code') as string
  if (!code) {
    return
  }
  getItem({ axios, code }).then(response => {
    console.log(response)
    plu.value = response
  })
}

</script>

<template>
  <div class="p-8">
    <h1 class="text-2xl font-bold mb-8 mt-4">Cek Harga</h1>
    <div class="text-left">
      <div class="">
        <form class="" @submit.prevent="handleSubmit">
          <div class="">
            <label
              for="barcode"
              class="block text-sm font-medium text-gray-500"
            >
              Barcode
            </label>
            <input
              type="text"
              id="barcode"
              name="code"
              v-model="code"
              autocomplete="barcode"
              class="mt-1 text-2xl focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm border-gray-300 rounded-md border-2"
            />
          </div>
        </form>
      </div>
      <div class="grid mt">
        <LabelValue name="Kode">{{plu.item.Kode}}</LabelValue>
        <LabelValue name="Barcode">{{plu.item.Barcode}}</LabelValue>
        <LabelValue name="Nama">{{plu.item.Nama}}</LabelValue>
        <LabelValue name="Harga">
          <span v-if="plu.item.HargaJual !== 0">
            <del v-if="plu.item.HargaNormal !== 0">{{nf.format(plu.item.HargaNormal)}}</del>
            {{nf.format(plu.item.HargaJual)}}
          </span>
        </LabelValue>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>

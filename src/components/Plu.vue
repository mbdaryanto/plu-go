<script setup lang="ts">
import { ref, reactive } from 'vue'
import { getItem, PluResponseType } from './item'
import axios from 'axios'

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
  getItem({ axios, code: formData.get('code') as string }).then(response => {
    console.log(response)
    plu.value = response
  })
}

</script>

<template>
  <h1>Cek Harga</h1>
  <div class="container">
    <div>
      <form @submit.prevent="handleSubmit">
        <input type="text" name="code" v-model="code"/>
      </form>
    </div>
    <div class="grid">
      <div>Nama</div><div>{{plu.item.Nama}}</div>
      <div>Kode</div><div>{{plu.item.Kode}}</div>
      <div>Barcode</div><div>{{plu.item.Barcode}}</div>
      <div>Harga</div><div><del>{{plu.item.HargaNormal}}</del> {{plu.item.HargaJual}}</div>
    </div>
  </div>
</template>

<style scoped>
.container {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  width: 100%;
}
.grid {
  display: grid;
  grid-template-columns: 200px 1fr;
  max-width: 600px;
  text-align: left;
}
</style>

import { createApp } from 'vue'
import Profile from './vue/profile.vue'
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

// Buat app yang merupakan file HTML yang dibuat di .vue
// Bisa dipakai di .html dengan memasukkan nama elemen yang ditulis di #
createApp(Profile).mount('#profileJs')

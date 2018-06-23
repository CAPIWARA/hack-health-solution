import Vue from 'vue';
import App from './App';
import Harsh from '@/components/Harsh'
import router from './router';

Vue.use(Harsh);

new Vue({
  el: '#root',
  router,
  render: (λ) => λ(App)
});

import Vue from 'vue';
import App from '@/App';
import Harsh from '@/components/Harsh'
import store from '@/store';
import router from '@/router';


Vue.use(Harsh);

new Vue({
  el: '#root',
  store,
  router,
  render: (λ) => λ(App)
});

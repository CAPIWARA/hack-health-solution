import Vue from 'vue';
import App from '@/App';
import Harsh from '@/components/Harsh'
import store from '@/store';
import router from '@/router';
import format from 'tiny-date-format';

Vue.filter('toDate', (date, formatter = 'DD/MM/YYYY') => format(date, formatter));
Vue.filter('toAnswer', (value) => value ? 'Sim' : 'Não');

Vue.use(Harsh);

new Vue({
  el: '#root',
  store,
  router,
  render: (λ) => λ(App)
});

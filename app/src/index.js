import Vue from 'vue';
import App from './App';
import router from './router';

new Vue({
  el: '#root',
  router,
  render: (λ) => λ(App)
});

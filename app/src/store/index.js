import Vue from 'vue';
import Vuex, { Store } from 'vuex';
import sarrada from '@/store/modules/sarrada';

Vue.use(Vuex);

const store = new Store({
  strict: true,
  modules: {
    sarrada
  }
});

export default store;

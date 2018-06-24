import * as types from '@/store/types';
import * as services from '@/services/sarrada';

const state = {
  sarrada: null
};

const getters = {
  [types.SARRADA]: (_) => _.sarrada
};

const mutations = {
  [types.SARRADA]: (_, sarrada) => {
    _.sarrada = sarrada
  }
};

const actions = {
  [types.SARRADA_CREATE]: async ({ dispatch }, params) => {
    const sarrada = await services.create(params);
    return sarrada;
  },
  [types.SARRADA]: async ({ commit }, id) => {
    const sarrada = await services.fetch(id);
    commit(types.SARRADA, sarrada);
  }
};

export default { state, getters, mutations, actions }

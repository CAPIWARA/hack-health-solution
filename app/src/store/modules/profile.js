import * as types from '@/store/types';
import * as services from '@/services/profile';

const state = {
  profile: null
};

const getters = {
  [types.PROFILE]: (_) => _.profile
};

const mutations = {
  [types.PROFILE]: (_, profile) => {
    _.profile = profile;
  }
};

const actions = {
  [types.PROFILE]: async ({ commit }) => {
    const profile = await services.fetch();
    commit(types.PROFILE, profile);
  }
};

export default { state, getters, mutations, actions };

import Vue from 'vue'
import Vuex, { StoreOptions } from 'vuex'

import { RootState } from '@/store/model/RootState'

Vue.use(Vuex)

const store: StoreOptions<RootState> = {
  actions: {},
  modules: {},
}

export default new Vuex.Store<RootState>(store)

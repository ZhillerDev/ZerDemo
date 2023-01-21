import { defineStore } from "pinia";
import NAMES from "./names";

const store = defineStore(NAMES.STORE, {
  state: ()=>{
    return{
      demo:1
    }
  },
  actions: {},
  getters: {},
});

export default store;
